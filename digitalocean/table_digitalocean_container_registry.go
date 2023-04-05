package digitalocean

import (
	"context"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDigitalOceanContainerRegistry(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_container_registry",
		Description: "DigitalOcean Container Registry",
		List: &plugin.ListConfig{
			Hydrate: getContainerRegistry,
		},
		Columns: []*plugin.Column{
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "A globally unique name for the container registry.",
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "A time value given in ISO8601 combined date and time format that represents when the registry was created.",
			},
			{
				Name:        "storage_usage_bytes",
				Type:        proto.ColumnType_INT,
				Description: "The amount of storage used in the registry in bytes.",
				Transform:   transform.FromField("StorageUsageBytes"),
			},
			{
				Name:        "storage_usage_bytes_updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The time at which the storage usage was updated. Storage usage is calculated asynchronously, and may not immediately reflect pushes to the registry.",
				Transform:   transform.FromField("StorageUsageBytesUpdatedAt"),
			},
			{
				Name:        "urn",
				Type:        proto.ColumnType_STRING,
				Description: "The uniform resource name (URN) for the container registry.",
				Transform:   transform.FromValue().Transform(registryToURN),
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
				Transform:   transform.FromValue().Transform(registryToURN).Transform(ensureStringArray),
			},
		},
	}
}

//// LIST FUNCTION

func getContainerRegistry(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_container_registry.listContainerRegistries", "connection_error", err)
		return nil, err
	}

	registry, _, err := conn.Registry.Get(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_container_registry.listContainerRegistries", "api_error", err)
		return nil, err
	}
	d.StreamListItem(ctx, registry)

	return nil, nil
}

//// TRANSFORM FUNCTION

func registryToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	registry := d.HydrateItem.(*godo.Registry)
	return "do:registry:" + registry.Name, nil
}
