package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDigitalOceanAction() *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_action",
		Description: "Actions are records of events that have occurred on the resources in your account. These can be things like rebooting a Droplet, or transferring an image to a new region.",
		List: &plugin.ListConfig{
			Hydrate: listAction,
		},
		Get: &plugin.GetConfig{
			KeyColumns:  plugin.SingleColumn("id"),
			ItemFromKey: actionFromKey,
			Hydrate:     getAction,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "A unique numeric ID that can be used to identify and reference an action."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "This is the type of action that the object represents.  For example, this could be \"transfer\" to represent the state of an image transfer action."},
			// Other columns
			{Name: "completed_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("CompletedAt").Transform(timestampToDateTime), Description: "A time value given in ISO8601 combined date and time format that represents when the action was completed."},
			{Name: "region_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Region.Name")},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Description: "The region where the action occurred."},
			{Name: "resource_id", Type: proto.ColumnType_INT, Description: "A unique identifier for the resource that the action is associated with."},
			{Name: "resource_type", Type: proto.ColumnType_STRING, Description: "The type of resource that the action is associated with."},
			{Name: "started_at", Type: proto.ColumnType_STRING, Transform: transform.FromField("StartedAt").Transform(timestampToDateTime), Description: "A time value given in ISO8601 combined date and time format that represents when the action was initiated."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "The current status of the action.  This can be \"in-progress\", \"completed\", or \"errored\"."},
		},
	}
}

func actionFromKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	quals := d.KeyColumnQuals
	id := quals["id"].GetInt64Value()
	item := &godo.Action{
		ID: int(id),
	}
	return item, nil
}

func listAction(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx)
	if err != nil {
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		actions, resp, err := conn.Actions.List(ctx, opts)
		if err != nil {
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
			return nil, err
		}
		// set the page we want for the next request
		opts.Page = page + 1
	}
	return nil, nil
}

func getAction(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	action := h.Item.(*godo.Action)
	conn, err := connect(ctx)
	if err != nil {
		return nil, err
	}
	result, _, err := conn.Actions.Get(ctx, action.ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}
