package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanRegion(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_region",
		Description: "A region in DigitalOcean represents a datacenter where Droplets can be deployed and images can be transferred. Each region represents a specific datacenter in a geographic location. Some geographical locations may have multiple \"regions\" available. This means that there are multiple datacenters available within that area.",
		List: &plugin.ListConfig{
			Hydrate: listRegion,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "A human-readable string that is used as a unique identifier for each region."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The display name of the region. This will be a full name that is used in the control panel and other interfaces."},
			// Other columns
			{Name: "available", Type: proto.ColumnType_BOOL, Description: "This is a boolean value that represents whether new Droplets can be created in this region."},
			{Name: "features", Type: proto.ColumnType_JSON, Description: "This attribute is set to an array which contains features available in this region."},
			{Name: "sizes", Type: proto.ColumnType_JSON, Description: "This attribute is set to an array which contains the identifying slugs for the sizes available in this region."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(regionToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		}),
	}
}

func listRegion(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_region.listRegion", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		regions, resp, err := conn.Regions.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_region.listRegion", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, t := range regions {
			d.StreamListItem(ctx, t)
		}
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_region.listRegion", "paging_error", err, "opts", opts, "page", page)
			return nil, err
		}
		// set the page we want for the next request
		opts.Page = page + 1
	}
	return nil, nil
}

func regionToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(godo.Region)
	return "do:region:" + i.Slug, nil
}
