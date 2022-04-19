package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableDigitalOceanSize(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_size",
		Description: "The sizes objects represent different packages of hardware resources that can be used for Droplets. This includes the amount of RAM, the number of virtual CPUs, disk space, and transfer. The size object also includes the pricing details and the regions that the size is available in.",
		List: &plugin.ListConfig{
			Hydrate: listSize,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "This is a boolean value that represents whether new Droplets can be created with this size."},
			// Other columns
			{Name: "available", Type: proto.ColumnType_BOOL, Description: "An array containing the region slugs where this size is available for Droplet creates."},
			{Name: "disk", Type: proto.ColumnType_INT, Description: "The amount of disk space set aside for Droplets of this size. The value is represented in gigabytes."},
			{Name: "memory", Type: proto.ColumnType_INT, Description: "The amount of RAM allocated to Droplets created of this size. The value is represented in megabytes."},
			{Name: "price_hourly", Type: proto.ColumnType_DOUBLE, Description: "This describes the price of the Droplet size as measured hourly. The value is measured in US dollars."},
			{Name: "price_monthly", Type: proto.ColumnType_DOUBLE, Description: "This attribute describes the monthly cost of this Droplet size if the Droplet is kept for an entire month. The value is measured in US dollars."},
			{Name: "regions", Type: proto.ColumnType_JSON, Description: "An array containing the region slugs where this size is available for Droplet creates."},
			{Name: "transfer", Type: proto.ColumnType_DOUBLE, Description: "The amount of transfer bandwidth that is available for Droplets created in this size. This only counts traffic on the public interface. The value is given in terabytes."},
			{Name: "vcpus", Type: proto.ColumnType_INT, Description: "The integer of number CPUs allocated to Droplets of this size."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(sizeToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Slug"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listSize(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_size.listSize", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		sizes, resp, err := conn.Sizes.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_size.listSize", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, t := range sizes {
			d.StreamListItem(ctx, t)
		}
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			return nil, err
		}
		// set the page we want for the next request
		opts.Page = page + 1
	}
	return nil, nil
}

func sizeToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(godo.Size)
	return "do:size:" + i.Slug, nil
}
