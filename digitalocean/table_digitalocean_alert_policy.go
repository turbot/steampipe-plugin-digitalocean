package digitalocean

import (
	"context"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDigitalOceanAlertPolicy(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_alert_policy",
		Description: "DigitalOcean Alert Policy",
		List: &plugin.ListConfig{
			Hydrate: listAlertPolicies,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("uuid"),
			Hydrate:    getAlertPolicy,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "uuid",
				Type:        proto.ColumnType_STRING,
				Description: "UUID of the alert policy.",
				Transform:   transform.FromField("UUID"),
			},
			{
				Name:        "compare",
				Type:        proto.ColumnType_STRING,
				Description: "The compare parameter for the metric in alert policy.",
			},
			{
				Name:        "enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "Alert Policy enabled or not.",
			},
			{
				Name:        "type",
				Type:        proto.ColumnType_STRING,
				Description: "Alert Policy type.",
			},
			{
				Name:        "description",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the alert policy.",
			},
			{
				Name:        "value",
				Type:        proto.ColumnType_INT,
				Description: "The value of the metric threshold in alert policy.",
			},
			{
				Name:        "interval",
				Type:        proto.ColumnType_STRING,
				Description: "The interval time of the metric in alert policy.",
				Transform:   transform.FromField("Window"),
			},
			{
				Name:        "urn",
				Type:        proto.ColumnType_STRING,
				Description: "The uniform resource name (URN) for the alert policy.",
				Transform:   transform.FromValue().Transform(alertPolicyToURN),
			},
			{
				Name:        "alerts",
				Type:        proto.ColumnType_JSON,
				Description: "The notification details where alert details will be send.",
			},
			{
				Name:        "entities",
				Type:        proto.ColumnType_JSON,
				Description: "An array of entities of the alert policy.",
			},
			{
				Name:        "tags_src",
				Type:        proto.ColumnType_JSON,
				Description: "An array of tags for the resource.",
				Transform:   transform.FromField("Tags"),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: resourceInterfaceDescription("title"),
				Transform:   transform.FromField("UUID"),
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("tags"),
				Transform:   transform.FromField("Tags").Transform(labelsToTagsMap),
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("akas"),
				Transform:   transform.FromValue().Transform(alertPolicyToURN).Transform(ensureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

func listAlertPolicies(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listAlertPolicies", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		policies, resp, err := conn.Monitoring.ListAlertPolicies(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("listAlertPolicies", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, policy := range policies {
			d.StreamListItem(ctx, policy)
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

func getAlertPolicy(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getAlertPolicy", "connection_error", err)
		return nil, err
	}

	uuid := d.EqualsQuals["uuid"].GetStringValue()

	// Handle empty uuid
	if uuid == "" {
		return nil, nil
	}

	result, resp, err := conn.Monitoring.GetAlertPolicy(ctx, uuid)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("getAlertPolicy", "not_found_error", err, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("getAlertPolicy", "query_error", err, "resp", resp)
		return nil, err
	}
	return result, nil
}

//// TRANSFORM FUNCTION

func alertPolicyToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	var policy godo.AlertPolicy
	switch d.Value.(type) {
	case *godo.AlertPolicy:
		policy = *d.Value.(*godo.AlertPolicy)
	case godo.AlertPolicy:
		policy = d.Value.(godo.AlertPolicy)
	}
	return "do:alertPolicy:" + policy.UUID, nil
}
