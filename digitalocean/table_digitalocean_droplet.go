package digitalocean

import (
	"context"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDigitalOceanDroplet(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_droplet",
		Description: "A Droplet is a DigitalOcean virtual machine.",
		List: &plugin.ListConfig{
			Hydrate: listDroplet,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getDroplet,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "A unique identifier for each Droplet instance."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The human-readable name set for the Droplet instance."},
			// Other columns
			{Name: "backup_ids", Type: proto.ColumnType_JSON, Description: "An array of backup IDs of any backups that have been taken of the Droplet instance."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the Droplet was created."},
			{Name: "disk", Type: proto.ColumnType_INT, Description: "The size of the Droplet's disk in gigabytes."},
			{Name: "features", Type: proto.ColumnType_JSON, Description: "An array of features enabled on this Droplet."},
			{Name: "image", Type: proto.ColumnType_JSON, Transform: transform.FromField("Image"), Description: "Information about the base image used to create the Droplet instance."},
			{Name: "kernel", Type: proto.ColumnType_JSON, Transform: transform.FromField("Kernel"), Description: "The current kernel. This will initially be set to the kernel of the base image when the Droplet is created."},
			{Name: "locked", Type: proto.ColumnType_BOOL, Description: "A boolean value indicating whether the Droplet has been locked, preventing actions by users."},
			{Name: "memory", Type: proto.ColumnType_INT, Description: "Memory of the Droplet in megabytes."},
			{Name: "networks", Type: proto.ColumnType_JSON, Description: "The details of the network that are configured for the Droplet instance. This is an object that contains keys for IPv4 and IPv6. The value of each of these is an array that contains objects describing an individual IP resource allocated to the Droplet. These will define attributes like the IP address, netmask, and gateway of the specific network depending on the type of network it is."},
			{Name: "next_backup_window_start", Type: proto.ColumnType_STRING, Transform: transform.FromField("NextBackupWindow.Start").Transform(timestampToIsoTimestamp), Description: "Start time of the window during which the backup will start."},
			{Name: "next_backup_window_end", Type: proto.ColumnType_STRING, Transform: transform.FromField("NextBackupWindow.End").Transform(timestampToIsoTimestamp), Description: "End time of the window during which the backup will start."},
			{Name: "private_ipv4", Type: proto.ColumnType_IPADDR, Transform: transform.FromMethod("PrivateIPv4"), Description: "Private IPv4 address of the Droplet."},
			{Name: "public_ipv4", Type: proto.ColumnType_IPADDR, Transform: transform.FromMethod("PublicIPv4"), Description: "Public IPv4 address of the Droplet."},
			{Name: "public_ipv6", Type: proto.ColumnType_IPADDR, Transform: transform.FromMethod("PublicIPv6"), Description: "Public IPv6 address of the Droplet."},
			{Name: "region", Type: proto.ColumnType_JSON, Transform: transform.FromField("Region"), Description: "Information about region that the Droplet instance is deployed in."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Region.Slug"), Description: "The unique slug identifier for the region the Droplet is deployed in."},
			{Name: "size", Type: proto.ColumnType_JSON, Transform: transform.FromField("Size"), Description: "Information about the size of the Droplet. Note: Due to resize operations, the disk column is more accurate than the disk field in this size data."},
			{Name: "size_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Size.Slug"), Description: "The unique slug identifier for the size of this Droplet."},
			{Name: "snapshot_ids", Type: proto.ColumnType_JSON, Description: "An array of snapshot IDs of any snapshots created from the Droplet instance."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "A status string indicating the state of the Droplet instance.  This may be \"new\", \"active\", \"off\", or \"archive\"."},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags"), Description: "An array of tags the Droplet has been tagged with."},
			{Name: "urn", Type: proto.ColumnType_STRING, Transform: transform.FromMethod("URN"), Description: "The uniform resource name (URN) for the Droplet."},
			{Name: "vcpus", Type: proto.ColumnType_INT, Description: "The number of virtual CPUs."},
			{Name: "volume_ids", Type: proto.ColumnType_JSON, Description: "A flat array including the unique identifier for each Block Storage volume attached to the Droplet."},
			{Name: "vpc_uuid", Type: proto.ColumnType_STRING, Description: "A string specifying the UUID of the VPC to which the Droplet is assigned."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromMethod("URN").Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags").Transform(labelsToTagsMap), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listDroplet(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_droplet.listDroplet", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		droplets, resp, err := conn.Droplets.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_droplet.listDroplet", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, t := range droplets {
			d.StreamListItem(ctx, t)
		}
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_droplet.listDroplet", "paging_error", err, "opts", opts, "page", page)
			return nil, err
		}
		// set the page we want for the next request
		opts.Page = page + 1
	}
	return nil, nil
}

func getDroplet(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_droplet.getDroplet", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := int(quals["id"].GetInt64Value())
	result, resp, err := conn.Droplets.Get(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("digitalocean_droplet.getDroplet", "not_found_error", err, "quals", quals, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_droplet.getDroplet", "query_error", err, "quals", quals, "resp", resp)
		return nil, err
	}
	return result, nil
}
