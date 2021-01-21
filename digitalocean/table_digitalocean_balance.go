package digitalocean

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableDigitalOceanBalance() *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_balance",
		Description: "Balance information for the current account.",
		List: &plugin.ListConfig{
			Hydrate: listBalance,
		},
		Columns: []*plugin.Column{
			{Name: "account_balance", Type: proto.ColumnType_DOUBLE, Description: "Current balance of the customer's most recent billing activity. Does not reflect month_to_date_usage."},
			{Name: "generated_at", Type: proto.ColumnType_DATETIME, Description: "The time at which balances were most recently generated."},
			{Name: "month_to_date_balance", Type: proto.ColumnType_DOUBLE, Description: "Balance as of the generated_at time. This value includes the account_balance and month_to_date_usage."},
			{Name: "month_to_date_usage", Type: proto.ColumnType_DOUBLE, Description: "Amount used in the current billing period as of the generated_at time."},
		},
	}
}

func listBalance(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx)
	if err != nil {
		return nil, err
	}
	balance, _, err := conn.Balance.Get(ctx)
	if err != nil {
		return nil, err
	}
	d.StreamListItem(ctx, balance)
	return nil, nil
}
