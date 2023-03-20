package digitalocean

import (
	"context"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableDigitalOceanVPC(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_vpc",
		Description: "VPCs (virtual private clouds) are virtual networks containing resources that can communicate with each other in full isolation using private IP addresses.",
		List: &plugin.ListConfig{
			Hydrate: listVPC,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getVPC,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "A unique ID that can be used to identify and reference the VPC."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the VPC. Must be unique and may only contain alphanumeric characters, dashes, and periods."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "A time value given in ISO8601 combined date and time format."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "A free-form text field for describing the VPC's purpose. It may be a maximum of 255 characters."},
			{Name: "ip_range", Type: proto.ColumnType_CIDR, Description: "The range of IP addresses in the VPC in CIDR notation."},
			// Rename to avoid conflict with default keyword in postgres
			{Name: "is_default", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Default"), Description: "A boolean value indicating whether or not the VPC is the default network for the region. All applicable resources are placed into the default VPC network unless otherwise specified during their creation. The `default` field cannot be unset from `true`. If you want to set a new default VPC network, update the `default` field of another VPC network in the same region. The previous network's `default` field will be set to `false` when a new default VPC has been defined."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("RegionSlug"), Description: "The slug identifier for the region where the VPC will be created."},
			{Name: "urn", Type: proto.ColumnType_STRING, Transform: transform.FromField("URN"), Description: "The uniform resource name (URN) for the VPC."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromField("URN").Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listVPC(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_vpc.listVPC", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		vpcs, resp, err := conn.VPCs.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_vpc.listVPC", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, t := range vpcs {
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

func getVPC(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_vpc.getVPC", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	result, resp, err := conn.VPCs.Get(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("digitalocean_vpc.getVPC", "not_found_error", err, "quals", quals, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_vpc.getVPC", "query_error", err, "quals", quals, "resp", resp)
		return nil, err
	}
	return result, nil
}
