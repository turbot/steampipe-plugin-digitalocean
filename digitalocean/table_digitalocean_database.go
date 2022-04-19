package digitalocean

import (
	"context"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableDigitalOceanDatabase(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_database",
		Description: "DigitalOcean's managed database service simplifies the creation and management of highly available database clusters. Currently, it offers support for PostgreSQL, Redis, and MySQL.",
		List: &plugin.ListConfig{
			Hydrate: listDatabase,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getDatabase,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "A unique ID that can be used to identify and reference a database cluster."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "A unique, human-readable name referring to a database cluster."},
			{Name: "engine", Type: proto.ColumnType_STRING, Description: "A slug representing the database engine used for the cluster. The possible values are: \"pg\" for PostgreSQL, \"mysql\" for MySQL, and \"redis\" for Redis."},
			{Name: "version", Type: proto.ColumnType_STRING, Description: "A string representing the version of the database engine in use for the cluster."},
			// Other columns
			{Name: "connection_database", Type: proto.ColumnType_STRING, Transform: transform.FromField("Connection.Database"), Description: "The name of the default database."},
			{Name: "connection_host", Type: proto.ColumnType_STRING, Transform: transform.FromField("Connection.Host"), Description: "A public FQDN pointing to the database cluster's current primary node."},
			{Name: "connection_password", Type: proto.ColumnType_STRING, Transform: transform.FromField("Connection.Password"), Description: "The randomly generated password for the default user."},
			{Name: "connection_port", Type: proto.ColumnType_INT, Transform: transform.FromField("Connection.Port"), Description: "The port on which the database cluster is listening."},
			{Name: "connection_ssl", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Connection.SSL"), Description: "A boolean value indicating if the connection should be made over SSL."},
			{Name: "connection_uri", Type: proto.ColumnType_STRING, Transform: transform.FromField("Connection.URI"), Description: "A connection string in the format accepted by the psql command. This is provided as a convenience and should be able to be constructed by the other attributes."},
			{Name: "connection_user", Type: proto.ColumnType_STRING, Transform: transform.FromField("Connection.User"), Description: "The default user for the database."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Description: "Time when the database was created."},
			{Name: "db_names", Type: proto.ColumnType_JSON, Description: "An array of strings containing the names of databases created in the database cluster.", Hydrate: getDatabaseNames, Transform: transform.FromValue()},
			{Name: "firewall_rules", Type: proto.ColumnType_JSON, Description: "A list of rules describing the inbound source to a database.", Hydrate: getDatabaseFirewallRules, Transform: transform.FromValue()},
			{Name: "maintenance_window_day", Type: proto.ColumnType_STRING, Transform: transform.FromField("MaintenanceWindow.Day"), Description: "The day of the week on which to apply maintenance updates (e.g. \"tuesday\")."},
			{Name: "maintenance_window_description", Type: proto.ColumnType_JSON, Transform: transform.FromField("MaintenanceWindow.Description"), Description: "A list of strings, each containing information about a pending maintenance update."},
			{Name: "maintenance_window_hour", Type: proto.ColumnType_STRING, Transform: transform.FromField("MaintenanceWindow.Hour"), Description: "The hour in UTC at which maintenance updates will be applied in 24 hour format (e.g. \"16:00:00\")."},
			{Name: "maintenance_window_pending", Type: proto.ColumnType_BOOL, Transform: transform.FromField("MaintenanceWindow.Pending"), Description: "A boolean value indicating whether any maintenance is scheduled to be performed in the next window."},
			{Name: "num_nodes", Type: proto.ColumnType_INT, Description: "The number of nodes in the database cluster."},
			{Name: "private_connection_database", Type: proto.ColumnType_STRING, Transform: transform.FromField("PrivateConnection.Database"), Description: "The name of the default database."},
			{Name: "private_connection_host", Type: proto.ColumnType_STRING, Transform: transform.FromField("PrivateConnection.Host"), Description: "The private FQDN pointing to the database cluster's current primary node."},
			{Name: "private_connection_password", Type: proto.ColumnType_STRING, Transform: transform.FromField("PrivateConnection.Password"), Description: "The randomly generated password for the default user."},
			{Name: "private_connection_port", Type: proto.ColumnType_INT, Transform: transform.FromField("PrivateConnection.Port"), Description: "The port on which the database cluster is listening."},
			{Name: "private_connection_ssl", Type: proto.ColumnType_BOOL, Transform: transform.FromField("PrivateConnection.SSL"), Description: "A boolean value indicating if the connection should be made over SSL."},
			{Name: "private_connection_uri", Type: proto.ColumnType_STRING, Transform: transform.FromField("PrivateConnection.URI"), Description: "A connection string in the format accepted by the psql command. This is provided as a convenience and should be able to be constructed by the other attributes."},
			{Name: "private_connection_user", Type: proto.ColumnType_STRING, Transform: transform.FromField("PrivateConnection.User"), Description: "The default user for the database."},
			{Name: "private_network_uuid", Type: proto.ColumnType_STRING, Description: "A string specifying the UUID of the VPC to which the database cluster is assigned."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("RegionSlug"), Description: "The unique slug identifier for the region the database is deployed in."},
			{Name: "size_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("SizeSlug"), Description: "The slug identifier representing the size of the nodes in the database cluster."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "A string representing the current status of the database cluster. Possible values include creating, online, resizing, and migrating."},
			{Name: "tags_src", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags"), Description: "An array of tags that have been applied to the database cluster."},
			{Name: "urn", Type: proto.ColumnType_STRING, Transform: transform.FromValue().Transform(databaseToURN), Description: "The uniform resource name (URN) for the database."},
			{Name: "users", Type: proto.ColumnType_JSON, Description: "An array containing objects describing the database's users.", Hydrate: getDatabaseUsers, Transform: transform.FromValue()},
			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(databaseToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags").Transform(labelsToTagsMap), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listDatabase(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_database.listDatabase", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		databases, resp, err := conn.Databases.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_database.listDatabase", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, t := range databases {
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

func getDatabase(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_database.getDatabase", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	result, resp, err := conn.Databases.Get(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("digitalocean_database.getDatabase", "not_found_error", err, "quals", quals, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_database.getDatabase", "query_error", err, "quals", quals, "resp", resp)
		return nil, err
	}
	return *result, nil
}

func getDatabaseUsers(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_database.getDatabaseUsers", "connection_error", err)
		return nil, err
	}
	id := h.Item.(godo.Database).ID
	users, resp, err := conn.Databases.ListUsers(ctx, id, nil)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("digitalocean_database.getDatabaseUsers", "not_found_error", err, "id", id, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_database.getDatabaseUsers", "query_error", err, "id", id, "resp", resp)
		return nil, err
	}
	return users, nil
}

func getDatabaseNames(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_database.getDatabaseNames", "connection_error", err)
		return nil, err
	}
	id := h.Item.(godo.Database).ID
	dbs, resp, err := conn.Databases.ListDBs(ctx, id, nil)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("digitalocean_database.getDatabaseNames", "not_found_error", err, "id", id, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_database.getDatabaseNames", "query_error", err, "id", id, "resp", resp)
		return nil, err
	}
	return dbs, nil
}

func getDatabaseFirewallRules(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_database.getDatabaseFirewallRules", "connection_error", err)
		return nil, err
	}
	id := h.Item.(godo.Database).ID
	firewall, resp, err := conn.Databases.GetFirewallRules(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("digitalocean_database.getDatabaseFirewallRules", "not_found_error", err, "id", id, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_database.getDatabaseFirewallRules", "query_error", err, "id", id, "resp", resp)
		return nil, err
	}
	return firewall, nil
}

func databaseToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(godo.Database)
	return "do:database:" + i.ID, nil
}
