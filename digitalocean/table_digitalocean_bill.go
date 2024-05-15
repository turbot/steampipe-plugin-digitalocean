package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableDigitalOceanBill(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_bill",
		Description: "Billing history is a record of billing events for your account. For example, entries may include events like payments made, invoices issued, or credits granted.",
		List: &plugin.ListConfig{
			Hydrate: listBill,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "date", Type: proto.ColumnType_TIMESTAMP, Description: "Time the billing history entry occured."},
			// Other columns
			{Name: "amount", Type: proto.ColumnType_STRING, Description: "Amount of the billing history entry."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the billing history entry."},
			{Name: "invoice_id", Type: proto.ColumnType_STRING, Description: "ID of the invoice associated with the billing history entry, if applicable."},
			{Name: "invoice_uuid", Type: proto.ColumnType_STRING, Description: "UUID of the invoice associated with the billing history entry, if applicable."},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of billing history entry."},
		}),
	}
}

func listBill(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_bill.listBill", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		billingHistory, resp, err := conn.BillingHistory.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_bill.listBill", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, i := range billingHistory.BillingHistory {
			d.StreamListItem(ctx, i)
		}
		// if we are at the last page, break out the for loop
		if resp.Links == nil || resp.Links.IsLastPage() {
			break
		}
		page, err := resp.Links.CurrentPage()
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_bill.listBill", "paging_error", err, "opts", opts, "page", page)
			return nil, err
		}
		// set the page we want for the next request
		opts.Page = page + 1
	}
	return nil, nil
}
