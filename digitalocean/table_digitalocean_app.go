package digitalocean

import (
	"context"
	"fmt"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDigitalOceanApp(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_app",
		Description: "DigitalOcean App",
		List: &plugin.ListConfig{
			Hydrate: listApps,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getApp,
		},
		Columns: []*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The id of the app.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the app",
				Transform:   transform.FromField("Spec.Name"),
			},
			{
				Name:        "owner_uuid",
				Type:        proto.ColumnType_STRING,
				Description: "OwnerUUID of the app.",
				Transform:   transform.FromField("OwnerUUID"),
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time when the app was created.",
			},
			{
				Name:        "last_deployment_created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time when the app last deployed.",
			},
			{
				Name:        "live_url",
				Type:        proto.ColumnType_STRING,
				Description: "The live URL of the app.",
				Transform:   transform.FromField("LiveURL"),
			},
			{
				Name:        "live_url_base",
				Type:        proto.ColumnType_STRING,
				Description: "The live URL base of the app.",
				Transform:   transform.FromField("LiveURLBase"),
			},
			{
				Name:        "live_domain",
				Type:        proto.ColumnType_STRING,
				Description: "The live domain of the app.",
			},
			{
				Name:        "tier_slug",
				Type:        proto.ColumnType_STRING,
				Description: "Tier slug of the app",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "Time when the app was updated.",
			},
			{
				Name:        "urn",
				Type:        proto.ColumnType_STRING,
				Description: "The uniform resource name (URN) for the app.",
				Transform:   transform.FromValue().Transform(appToURN),
			},
			{
				Name:        "active_deployment",
				Type:        proto.ColumnType_JSON,
				Description: "The app's currently active deployment.",
			},
			{
				Name:        "in_progress_deployment",
				Type:        proto.ColumnType_JSON,
				Description: "In progress deployment of the app.",
			},
			{
				Name:        "region",
				Type:        proto.ColumnType_JSON,
				Description: "The DigitalOcean data center region hosting the app.",
			},
			{
				Name:        "spec",
				Type:        proto.ColumnType_JSON,
				Description: "A DigitalOcean App spec describing the app.",
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: resourceInterfaceDescription("title"),
				Transform:   transform.FromField("Spec.Name"),
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("akas"),
				Transform:   transform.FromValue().Transform(appToURN).Transform(ensureStringArray),
			},
		},
	}
}

//// LIST FUNCTION

func listApps(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listApps", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		apps, resp, err := conn.Apps.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("listApps", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, app := range apps {
			d.StreamListItem(ctx, app)
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

//// HYDRATE FUNCTIONS

func getApp(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getApp", "connection_error", err)
		return nil, err
	}

	id := d.EqualsQuals["id"].GetStringValue()

	// Handle empty id
	if id == "" {
		return nil, nil
	}

	result, resp, err := conn.Apps.Get(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("getApp", "not_found_error", err, "resp", resp)
			return nil, nil
		}
		if strings.Contains(err.Error(), ": 400") {
			plugin.Logger(ctx).Warn("getApp", "invalid_id", err, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("getApp", "query_error", err, "resp", resp)
		return nil, err
	}
	return result, nil
}

//// TRANSFORM FUNCTION

func appToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(*godo.App)
	return fmt.Sprintf("do:app:%s", i.ID), nil
}
