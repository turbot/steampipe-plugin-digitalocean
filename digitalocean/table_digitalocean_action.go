package digitalocean

import (
	"context"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanAction(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_action",
		Description: "Actions are records of events that have occurred on the resources in your account. These can be things like rebooting a Action, or transferring an image to a new region.",
		List: &plugin.ListConfig{
			Hydrate: listAction,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getAction,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "A unique numeric ID that can be used to identify and reference an action."},
			{Name: "resource_id", Type: proto.ColumnType_INT, Description: "A unique identifier for the resource that the action is associated with."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "This is the type of action that the object represents.  For example, this could be \"transfer\" to represent the state of an image transfer action."},
			{Name: "started_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("StartedAt").Transform(timestampToIsoTimestamp), Description: "Time when when the action was initiated."},
			// Other columns
			{Name: "completed_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CompletedAt").Transform(timestampToIsoTimestamp), Description: "Time when the action was completed."},
			// Skip this object dump for now, they can join on the slug if really necessary
			//{Name: "region", Type: proto.ColumnType_JSON, Transform: transform.FromField("Region"), Description: "A full region object containing information about the region where the action occurred."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Description: "The region where the action occurred."},
			{Name: "resource_type", Type: proto.ColumnType_STRING, Description: "The type of resource that the action is associated with."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The current status of the action.  This can be \"in-progress\", \"completed\", or \"errored\"."},
		},
	}
}

func listAction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_action.listAction", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		actions, resp, err := conn.Actions.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_action.listAction", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, t := range actions {
			d.StreamListItem(ctx, t)
		}
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_action.listAction", "paging_error", err, "opts", opts, "page", page)
			return nil, err
		}
		// set the page we want for the next request
		opts.Page = page + 1
	}
	return nil, nil
}

func getAction(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_action.getAction", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := int(quals["id"].GetInt64Value())
	result, resp, err := conn.Actions.Get(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("digitalocean_action.getAction", "not_found_error", err, "quals", quals, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_action.getAction", "query_error", err, "quals", quals, "resp", resp)
		return nil, err
	}
	return result, nil
}
