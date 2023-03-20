package digitalocean

import (
	"context"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableDigitalOceanFloatingIP(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_floating_ip",
		Description: "DigitalOcean Floating IPs are publicly-accessible static IP addresses that can be mapped to one of your Droplets. They can be used to create highly available setups or other configurations requiring movable addresses.",
		List: &plugin.ListConfig{
			Hydrate: listFloatingIP,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("ip"),
			Hydrate:    getFloatingIP,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "ip", Type: proto.ColumnType_IPADDR, Description: "The public IP address of the floating IP. It also serves as its identifier."},
			// Other columns
			{Name: "droplet", Type: proto.ColumnType_JSON, Transform: transform.FromField("Droplet"), Description: "The Droplet that the floating IP has been assigned to."},
			{Name: "droplet_id", Type: proto.ColumnType_INT, Transform: transform.FromField("Droplet.ID"), Description: "ID of the Droplet that the floating IP has been assigned to."},
			{Name: "region", Type: proto.ColumnType_JSON, Transform: transform.FromField("Region"), Description: "The region that the floating IP is reserved to."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Region.Slug"), Description: "The slug of the region that the floating IP is reserved to."},
			{Name: "urn", Type: proto.ColumnType_STRING, Transform: transform.FromMethod("URN"), Description: "The uniform resource name (URN) for the database."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromMethod("URN").Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("IP"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listFloatingIP(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_floating_ip.listFloatingIP", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		floatingIPs, resp, err := conn.FloatingIPs.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_floating_ip.listFloatingIP", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, t := range floatingIPs {
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

func getFloatingIP(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_floating_ip.getFloatingIP", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	ip := quals["ip"].GetInetValue().GetAddr()
	result, resp, err := conn.FloatingIPs.Get(ctx, ip)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("digitalocean_floating_ip.getFloatingIP", "not_found_error", err, "quals", quals, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_floating_ip.getFloatingIP", "query_error", err, "quals", quals, "resp", resp)
		return nil, err
	}
	return result, nil
}
