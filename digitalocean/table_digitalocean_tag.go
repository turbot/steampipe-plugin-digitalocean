package digitalocean

import (
	"context"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanTag(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_tag",
		Description: "A tag is a label that can be applied to a resource (currently Droplets, Images, Volumes, Volume Snapshots, and Database clusters) in order to better organize or facilitate the lookups and actions on it.",
		List: &plugin.ListConfig{
			Hydrate: listTag,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getTag,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the tag. Tags may contain letters, numbers, colons, dashes, and underscores. There is a limit of 255 characters per tag."},
			{Name: "resource_count", Type: proto.ColumnType_INT, Transform: transform.FromField("Resources.Count"), Description: "The number of resources with this tag."},
			{Name: "resources", Type: proto.ColumnType_JSON, Transform: transform.FromField("Resources"), Description: "An embedded object containing key value pairs of resource type and resource statistics. It also includes a count of the total number of resources tagged with the current tag as well as a last_tagged_uri attribute set to the last resource tagged with the current tag."},
		}),
	}
}

func listTag(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_tag.listTag", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		tags, resp, err := conn.Tags.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_tag.listTag", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, t := range tags {
			d.StreamListItem(ctx, t)
		}
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_tag.listTag", "paging_error", err, "opts", opts, "page", page)
			return nil, err
		}
		// set the page we want for the next request
		opts.Page = page + 1
	}
	return nil, nil
}

func getTag(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_tag.getTag", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	name := quals["name"].GetStringValue()
	result, resp, err := conn.Tags.Get(ctx, name)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("digitalocean_tag.getTag", "not_found_error", err, "quals", quals, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_tag.getTag", "query_error", err, "quals", quals, "resp", resp)
		return nil, err
	}
	return result, nil
}
