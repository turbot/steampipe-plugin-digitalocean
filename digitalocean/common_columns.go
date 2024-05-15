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
			Name:        "profile_id",
			Description: "The unique universal identifier for the current user.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getProfileId,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getProfileIdMemoized = plugin.HydrateFunc(getProfileIdUncached).Memoize(memoize.WithCacheKeyFunction(getProfileIdCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getProfileId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getProfileIdMemoized(ctx, d, h)
}

// Build a cache key for the call to getProfileIdCacheKey.
func getProfileIdCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getProfileId"
	return key, nil
}

func getProfileIdUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
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
