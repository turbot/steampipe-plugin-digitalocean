package digitalocean

import (
	"context"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableDigitalOceanSnapshot(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_snapshot",
		Description: "Snapshots are saved instances of a Droplet or a block storage volume.",
		List: &plugin.ListConfig{
			Hydrate: listSnapshot,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getSnapshot,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "The unique identifier for the snapshot."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "A human-readable name for the snapshot."},
			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the block storage volume was created."},
			{Name: "min_disk_size", Type: proto.ColumnType_INT, Description: "The minimum size in GB required for a volume or Droplet to use this snapshot."},
			{Name: "regions", Type: proto.ColumnType_JSON, Description: "An array of regions the snapshot is available in. The region slug is used."},
			{Name: "resource_id", Type: proto.ColumnType_STRING, Description: "A unique identifier for the resource that the action is associated with."},
			{Name: "resource_type", Type: proto.ColumnType_STRING, Description: "The type of resource that the action is associated with."},
			{Name: "size_gigabytes", Type: proto.ColumnType_INT, Description: "The billable size of the snapshot in gigabytes."},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags"), Description: "An array of Tags the snapshot has been tagged with."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(snapshotToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags").Transform(labelsToTagsMap), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		}),
	}
}

func listSnapshot(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_snapshot.listSnapshot", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		snapshots, resp, err := conn.Snapshots.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_snapshot.listSnapshot", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, t := range snapshots {
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

func getSnapshot(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_snapshot.getSnapshot", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := quals["id"].GetStringValue()
	result, resp, err := conn.Snapshots.Get(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("digitalocean_snapshot.getSnapshot", "not_found_error", err, "quals", quals, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_snapshot.getSnapshot", "query_error", err, "quals", quals, "resp", resp)
		return nil, err
	}
	return *result, nil
}

func snapshotToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(godo.Snapshot)
	return "do:snapshot:" + i.ID, nil
}
