package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDigitalOceanAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_account",
		Description: "Retrieve information about your current account.",
		List: &plugin.ListConfig{
			Hydrate: listAccount,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "email", Type: proto.ColumnType_STRING, Description: "The email address used by the current user to register for DigitalOcean."},
			// Other columns
			{Name: "droplet_limit", Type: proto.ColumnType_INT, Description: "The total number of Droplets the current user or team may have at one time."},
			{Name: "email_verified", Type: proto.ColumnType_BOOL, Description: "If true, the user has verified their account via email. False otherwise."},
			{Name: "floating_ip_limit", Type: proto.ColumnType_INT, Description: "The total number of floating IPs the current user or team may have."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "This value is one of \"active\", \"warning\" or \"locked\"."},
			{Name: "status_message", Type: proto.ColumnType_STRING, Description: "A human-readable message giving more details about the status of the account."},
			{Name: "uuid", Type: proto.ColumnType_STRING, Description: "The unique universal identifier for the current user."},
			{Name: "volume_limit", Type: proto.ColumnType_INT, Description: "The total number of volumes the current user or team may have."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(accountAkas), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Email"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listAccount(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	account, _, err := conn.Account.Get(ctx)
	if err != nil {
		return nil, err
	}
	d.StreamListItem(ctx, account)
	return nil, nil
}

func accountAkas(_ context.Context, d *transform.TransformData) (interface{}, error) {
	a := d.Value.(*godo.Account)
	return []string{"do:account:" + a.UUID}, nil
}
