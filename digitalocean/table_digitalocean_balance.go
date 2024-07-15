package digitalocean

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDigitalOceanBalance(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_balance",
		Description: "Balance information for the current account.",
		List: &plugin.ListConfig{
			Hydrate: listBalance,
		},
		Columns: commonColumns([]*plugin.Column{
			{Name: "account_balance", Type: proto.ColumnType_DOUBLE, Description: "Current balance of the customer's most recent billing activity. Does not reflect month_to_date_usage."},
			{Name: "generated_at", Type: proto.ColumnType_TIMESTAMP, Description: "The time at which balances were most recently generated."},
			{Name: "month_to_date_balance", Type: proto.ColumnType_DOUBLE, Description: "Balance as of the generated_at time. This value includes the account_balance and month_to_date_usage."},
			{Name: "month_to_date_usage", Type: proto.ColumnType_DOUBLE, Description: "Amount used in the current billing period as of the generated_at time."},
		}),
	}
}

func listBalance(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_balance.listBalance", "connection_error", err)
		return nil, err
	}
	balance, resp, err := conn.Balance.Get(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_balance.listBalance", "query_error", err, "resp", resp)
		return nil, err
	}
	d.StreamListItem(ctx, balance)
	return nil, nil
}
