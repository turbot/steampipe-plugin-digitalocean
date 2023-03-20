package digitalocean

import (
	"context"
	"fmt"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableDigitalOceanKey(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_key",
		Description: "DigitalOcean allows you to add SSH public keys to the interface so that you can embed your public key into a Droplet at the time of creation. Only the public key is required to take advantage of this functionality.",
		List: &plugin.ListConfig{
			Hydrate: listKey,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AnyColumn([]string{"id", "fingerprint"}),
			Hydrate:    getKey,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_INT, Description: "This is a unique identification number for the key."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "The human-readable display name for the given SSH key."},
			// Other columns
			{Name: "fingerprint", Type: proto.ColumnType_STRING, Description: "The fingerprint value that is generated from the public key."},
			{Name: "public_key", Type: proto.ColumnType_STRING, Description: "The entire public key string that was uploaded.  This is what is embedded into the root user's authorized_keys file if you choose to include this SSH key during Droplet creation."},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(keyToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromConstant(map[string]bool{}), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listKey(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_key.listKey", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		keys, resp, err := conn.Keys.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_key.listKey", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, t := range keys {
			d.StreamListItem(ctx, t)
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

func getKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_key.getKey", "connection_error", err)
		return nil, err
	}
	quals := d.EqualsQuals
	id := int(quals["id"].GetInt64Value())
	fingerprint := quals["fingerprint"].GetStringValue()
	var result *godo.Key
	var resp *godo.Response
	if id != 0 {
		result, resp, err = conn.Keys.GetByID(ctx, id)
	} else if len(fingerprint) > 0 {
		result, resp, err = conn.Keys.GetByFingerprint(ctx, fingerprint)
	} else {
		plugin.Logger(ctx).Warn("digitalocean_key.getKey", "invalid_quals", "id and fingerprint both empty", "quals", quals)
		return nil, nil
	}
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("digitalocean_key.getKey", "not_found_error", err, "quals", quals, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_key.getKey", "query_error", err, "quals", quals, "resp", resp)
		return nil, err
	}
	return result, nil
}

func keyToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	var i godo.Key
	switch d.Value.(type) {
	case *godo.Key:
		i = *d.Value.(*godo.Key)
	case godo.Key:
		i = d.Value.(godo.Key)
	}
	return fmt.Sprintf("do:key:%d", i.ID), nil
}
