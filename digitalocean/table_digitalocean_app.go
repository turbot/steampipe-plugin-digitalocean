package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableDigitalOceanApp() *plugin.Table {
	return &plugin.Table{
		Name: "digitalocean_app",
		List: &plugin.ListConfig{
			Hydrate: listApp,
		},
		Get: &plugin.GetConfig{
			KeyColumns:  plugin.SingleColumn("id"),
			ItemFromKey: appFromKey,
			Hydrate:     getApp,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING},
			// Other columns
			{Name: "owner_uuid", Type: proto.ColumnType_STRING},
			{Name: "spec", Type: proto.ColumnType_JSON},
			{Name: "default_ingress", Type: proto.ColumnType_STRING},
			{Name: "created_at", Type: proto.ColumnType_DATETIME},
			{Name: "updated_at", Type: proto.ColumnType_DATETIME},
			{Name: "active_deployment", Type: proto.ColumnType_JSON},
			{Name: "in_progress_deployment", Type: proto.ColumnType_JSON},
			{Name: "last_deployment_created_at", Type: proto.ColumnType_DATETIME},
			{Name: "live_url", Type: proto.ColumnType_STRING},
			{Name: "region", Type: proto.ColumnType_JSON},
			{Name: "tier_slug", Type: proto.ColumnType_STRING},
			{Name: "live_url_base", Type: proto.ColumnType_STRING},
			{Name: "live_domain", Type: proto.ColumnType_STRING},
		},
	}
}

func appFromKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	item := &godo.App{
		ID: id,
	}
	return item, nil
}

func listApp(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		apps, resp, err := conn.Apps.List(ctx, opts)
		if err != nil {
			return nil, err
		}
		for _, t := range apps {
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

func getApp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	app := h.Item.(*godo.App)
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	result, _, err := conn.Apps.Get(ctx, app.ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
