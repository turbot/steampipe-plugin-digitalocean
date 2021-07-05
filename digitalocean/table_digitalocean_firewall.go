package digitalocean

import (
	"context"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDigitalOceanFirewall(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_firewall",
		Description: "DigitalOcean Cloud Firewalls are a network-based, stateful firewall service for Droplets provided at no additional cost. Cloud firewalls block all traffic that isn’t expressly permitted by a rule.",
		List: &plugin.ListConfig{
			Hydrate: listFirewalls,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getFirewall,
		},
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique universal identifier of this firewall.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The name of the Firewall.",
			},

			// Other columns
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "A time value given in ISO8601 combined date and time format that represents when the Firewall was created.",
			},
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "A status string indicating the current state of the Firewall. ",
			},
			{
				Name:        "droplet_ids",
				Type:        proto.ColumnType_JSON,
				Description: "The list of the IDs of the Droplets assigned to the Firewall.",
				Transform:   transform.FromField("DropletIDs"),
			},
			{
				Name:        "inbound_rules",
				Type:        proto.ColumnType_JSON,
				Description: "The inbound access rule block for the Firewall.",
			},
			{
				Name:        "outbound_rules",
				Type:        proto.ColumnType_JSON,
				Description: "The outbound access rule block for the Firewall.",
			},
			{
				Name:        "pending_changes",
				Type:        proto.ColumnType_JSON,
				Description: "An list of object containing the fields, `droplet_id`, `removing`, and `status`. It is provided to detail exactly which Droplets are having their security policies updated. When empty, all changes have been successfully applied.",
			},

			// Resource interface
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("akas"),
				Transform:   transform.FromValue().Transform(firewallToURN).Transform(ensureStringArray),
			},
			{
				Name:        "tags",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("tags"),
				Transform:   transform.FromField("Tags").Transform(labelsToTagsMap),
			},
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: resourceInterfaceDescription("title"),
				Transform:   transform.FromField("Name"),
			},
		},
	}
}

func listFirewalls(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listFirewalls", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		firewalls, resp, err := conn.Firewalls.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("listFirewalls", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, firewall := range firewalls {
			d.StreamListItem(ctx, firewall)
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

func getFirewall(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getFirewall", "connection_error", err)
		return nil, err
	}

	id := d.KeyColumnQuals["id"].GetStringValue()

	firewall, resp, err := conn.Firewalls.Get(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("getFirewall", "not_found_error", err, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("getFirewall", "query_error", err, "resp", resp)
		return nil, err
	}
	return firewall, nil
}

func firewallToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	var firewall godo.Firewall
	switch d.Value.(type) {
	case *godo.Firewall:
		firewall = *d.Value.(*godo.Firewall)
	case godo.Firewall:
		firewall = d.Value.(godo.Firewall)
	}
	return "do:firewall:" + firewall.ID, nil
}
