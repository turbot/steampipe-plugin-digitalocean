package digitalocean

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "uuid",
			Description: "The unique universal identifier for the current user.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getCurrentUserUuid,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getCurrentUserUuidMemoized = plugin.HydrateFunc(getCurrentUserUuidUncached).Memoize(memoize.WithCacheKeyFunction(getCurrentUserUuidCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getCurrentUserUuid(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getCurrentUserUuidMemoized(ctx, d, h)
}

// Build a cache key for the call to getCurrentUserUuidCacheKey.
func getCurrentUserUuidCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getCurrentUserUuid"
	return key, nil
}

func getCurrentUserUuidUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}
	account, _, err := conn.Account.Get(ctx)
	if err != nil {
		return nil, err
	}
	return account.UUID, nil
}
