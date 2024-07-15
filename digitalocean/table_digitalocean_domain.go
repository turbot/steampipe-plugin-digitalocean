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

func tableDigitalOceanDomain(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_domain",
		Description: "DigitalOcean Domain",
		List: &plugin.ListConfig{
			Hydrate: listDomains,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("name"),
			Hydrate:    getDomain,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The globally unique human-readable name for the domain.",
			},
			{
				Name:        "ttl",
				Type:        proto.ColumnType_INT,
				Description: "TTL value of domain.",
			},
			{
				Name:        "zone_file",
				Type:        proto.ColumnType_STRING,
				Description: "It contains the DNS record details.",
			},
			{
				Name:        "urn",
				Type:        proto.ColumnType_STRING,
				Description: "The uniform resource name (URN) for the domain.",
				Transform:   transform.FromValue().Transform(domainToURN),
			},

			// Steampipe standard columns
			{
				Name:        "title",
				Type:        proto.ColumnType_STRING,
				Description: resourceInterfaceDescription("title"),
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("akas"),
				Transform:   transform.FromValue().Transform(domainToURN).Transform(ensureStringArray),
			},
		}),
	}
}

//// LIST FUNCTION

func listDomains(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listDomains", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		domains, resp, err := conn.Domains.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("listDomains", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, domain := range domains {
			d.StreamListItem(ctx, domain)
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

func getDomain(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getDomain", "connection_error", err)
		return nil, err
	}

	name := d.EqualsQuals["name"].GetStringValue()

	// Handle empty name
	if name == "" {
		return nil, nil
	}

	result, resp, err := conn.Domains.Get(ctx, name)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("getDomain", "not_found_error", err, "resp", resp)
			return nil, nil
		}
		if strings.Contains(err.Error(), ": 400") {
			plugin.Logger(ctx).Warn("getDomain", "invalid_name", err, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("getDomain", "query_error", err, "resp", resp)
		return nil, err
	}
	return result, nil
}

//// TRANSFORM FUNCTION

func domainToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	var i godo.Domain
	switch d.Value.(type) {
	case *godo.Domain:
		i = *d.Value.(*godo.Domain)
	case godo.Domain:
		i = d.Value.(godo.Domain)
	}
	return "do:domain:" + i.Name, nil
}
