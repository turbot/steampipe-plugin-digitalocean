package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

//// TABLE DEFINITION

func tableDigitalOceanSpace(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_space",
		Description: "DigitalOcean Space",
		List: &plugin.ListConfig{
			Hydrate: listSpaces,
		},
		Columns: []*plugin.Column{
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
		},
	}
}

//// LIST FUNCTION

func listSpaces(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connectSpaces(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listSpaces", "connection_error", err)
		return nil, err
	}

	spaces, err := conn.GetBucketPolicy()()
	if err != nil {
		plugin.Logger(ctx).Error("listSpaces", "query_error", err)
		return nil, err
	}
	for _, space := range spaces {
		d.StreamListItem(ctx, space)
	}

	return nil, nil
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
