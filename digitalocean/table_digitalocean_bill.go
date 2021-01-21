package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
)

func tableDigitalOceanBill() *plugin.Table {
	return &plugin.Table{
		Name: "digitalocean_bill",
		List: &plugin.ListConfig{
			Hydrate: listBill,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "date", Type: proto.ColumnType_DATETIME},
			// Other columns
			{Name: "amount", Type: proto.ColumnType_DOUBLE},
			{Name: "description", Type: proto.ColumnType_STRING},
			{Name: "invoice_id", Type: proto.ColumnType_STRING},
			{Name: "invoice_uuid", Type: proto.ColumnType_STRING},
			{Name: "type", Type: proto.ColumnType_STRING},
		},
	}
}

func listBill(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx)
	if err != nil {
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		billingHistory, resp, err := conn.BillingHistory.List(ctx, opts)
		if err != nil {
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
			return nil, err
		}
		// set the page we want for the next request
		opts.Page = page + 1
	}
	return nil, nil
}
