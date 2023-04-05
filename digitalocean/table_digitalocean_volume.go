package digitalocean

import (
	"context"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanVolume(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_volume",
		Description: "DigitalOcean Block Storage Volumes provide expanded storage capacity for your Droplets and can be moved between Droplets within a specific region. Volumes function as raw block devices, meaning they appear to the operating system as locally attached storage which can be formatted using any file system supported by the OS.",
		List: &plugin.ListConfig{
			Hydrate: listVolume,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getVolume,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the block storage volume."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "A human-readable name for the block storage volume. Must be lowercase and be composed only of numbers, letters and \"-\", up to a limit of 64 characters. The name must begin with a letter."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the block storage volume was created."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "An optional free-form text field to describe a block storage volume."},
			{Name: "droplet_ids", Type: proto.ColumnType_JSON, Description: "An array containing the IDs of the Droplets the volume is attached to. Note that at this time, a volume can only be attached to a single Droplet."},
			{Name: "filesystem_label", Type: proto.ColumnType_STRING, Description: "The label currently applied to the filesystem."},
			{Name: "filesystem_type", Type: proto.ColumnType_STRING, Description: "The type of filesystem currently in-use on the volume."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Region.Slug"), Description: "The unique slug identifier for the region the volume is deployed in."},
			{Name: "region_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Region.Name"), Description: "The name of the region the volume is deployed in."},
			{Name: "size_gigabytes", Type: proto.ColumnType_INT, Description: "The size of the block storage volume in GiB (1024^3)."},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags"), Description: "An array of tags the volume has been tagged with"},
			{Name: "urn", Type: proto.ColumnType_STRING, Transform: transform.FromMethod("URN"), Description: "The uniform resource name (URN) for the volume."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromMethod("URN").Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags").Transform(labelsToTagsMap), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listVolume(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_volume.listVolume", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListVolumeParams{
		ListOptions: &godo.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	}
	for {
		volumes, resp, err := conn.Storage.ListVolumes(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_volume.listVolume", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, t := range volumes {
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
		opts.ListOptions.Page = page + 1
	}
	return nil, nil
}

func getVolume(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_volume.getVolume", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	result, resp, err := conn.Storage.GetVolume(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("digitalocean_volume.getVolume", "not_found_error", err, "quals", quals, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_volume.getVolume", "query_error", err, "quals", quals, "resp", resp)
		return nil, err
	}
	return result, nil
}
