package digitalocean

import (
	"context"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDigitalOceanLoadBalancer(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_load_balancer",
		Description: "DigitalOcean Load Balancers provide a way to distribute traffic across multiple Droplets.",
		List: &plugin.ListConfig{
			Hydrate: listLoadBalancer,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getLoadBalancer,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "A unique ID that can be used to identify and reference a load balancer."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "A human-readable name for a load balancer instance."},
			// Other columns
			{Name: "algorithm", Type: proto.ColumnType_STRING, Description: "The load balancing algorithm used to determine which backend Droplet will be selected by a client. It must be either \"round_robin\" or \"least_connections\"."},
			{Name: "created_at", Type: proto.ColumnType_DATETIME, Description: "Time when the load balancer was created."},
			{Name: "droplet_ids", Type: proto.ColumnType_JSON, Description: "An array containing the IDs of the Droplets assigned to the load balancer."},
			{Name: "enable_backend_keepalive", Type: proto.ColumnType_BOOL, Description: "A boolean value indicating whether HTTP keepalive connections are maintained to target Droplets."},
			{Name: "enable_proxy_protocol", Type: proto.ColumnType_BOOL, Description: "A boolean value indicating whether PROXY Protocol is in use."},
			{Name: "forwarding_rules", Type: proto.ColumnType_JSON, Description: "An object specifying the forwarding rules for a load balancer."},
			{Name: "health_check_healthy_threshold", Type: proto.ColumnType_INT, Transform: transform.FromField("HealthCheck.HealthyThreshold"), Description: "The number of times a health check must pass for a backend Droplet to be marked \"healthy\" and be re-added to the pool."},
			{Name: "health_check_interval_seconds", Type: proto.ColumnType_INT, Transform: transform.FromField("HealthCheck.CheckIntervalSeconds"), Description: "The number of seconds between between two consecutive health checks."},
			{Name: "health_check_path", Type: proto.ColumnType_STRING, Transform: transform.FromField("HealthCheck.Path"), Description: "The path on the backend Droplets to which the load balancer instance will send a request."},
			{Name: "health_check_port", Type: proto.ColumnType_INT, Transform: transform.FromField("HealthCheck.Port"), Description: "An integer representing the port on the backend Droplets on which the health check will attempt a connection."},
			{Name: "health_check_protocol", Type: proto.ColumnType_STRING, Transform: transform.FromField("HealthCheck.Protocol"), Description: "The protocol used for health checks sent to the backend Droplets. The possible values are \"http\", \"https\", or \"tcp\"."},
			{Name: "health_check_response_timeout_seconds", Type: proto.ColumnType_INT, Transform: transform.FromField("HealthCheck.ResponseTimeoutSeconds"), Description: "The number of seconds the load balancer instance will wait for a response until marking a health check as failed."},
			{Name: "health_check_unhealthy_threshold", Type: proto.ColumnType_INT, Transform: transform.FromField("HealthCheck.UnhealthyThreshold"), Description: "The number of times a health check must fail for a backend Droplet to be marked \"unhealthy\" and be removed from the pool."},
			{Name: "ip", Type: proto.ColumnType_IPADDR, Description: "An attribute containing the public-facing IP address of the load balancer."},
			{Name: "redirect_http_to_https", Type: proto.ColumnType_BOOL, Description: "A boolean value indicating whether HTTP requests to the load balancer on port 80 will be redirected to HTTPS on port 443."},
			{Name: "region_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Region.Slug"), Description: "The unique slug identifier for the region the load balancer is deployed in."},
			{Name: "region_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Region.Name"), Description: "The name of the region the load balancer is deployed in."},
			{Name: "size", Type: proto.ColumnType_STRING, Description: "The size of the load balancer. The available sizes are \"lb-small\", \"lb-medium\", or \"lb-large\"."},
			{Name: "status", Type: proto.ColumnType_STRING, Description: "A status string indicating the current state of the load balancer. This can be \"new\", \"active\", or \"errored\"."},
			{Name: "sticky_sessions_cookie_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("StickySessions.CookieName").NullIfZero(), Description: "The name of the cookie sent to the client. This attribute is only returned when using \"cookies\" for the sticky sessions type."},
			{Name: "sticky_sessions_cookie_ttl_seconds", Type: proto.ColumnType_INT, Transform: transform.FromField("StickySessions.CookieTtlSeconds").NullIfZero(), Description: "The number of seconds until the cookie set by the load balancer expires. This attribute is only returned when using \"cookies\" for the sticky sessions type."},
			{Name: "sticky_sessions_type", Type: proto.ColumnType_STRING, Transform: transform.FromField("StickySessions.Type"), Description: "An attribute indicating how and if requests from a client will be persistently served by the same backend Droplet. The possible values are \"cookies\" or \"none\"."},
			{Name: "tag", Type: proto.ColumnType_STRING, Description: "The name of a Droplet tag corresponding to Droplets assigned to the load balancer."},
			{Name: "urn", Type: proto.ColumnType_STRING, Transform: transform.FromValue().Transform(loadBalancerToURN), Description: "The uniform resource name (URN) for the load balancer."},
			{Name: "vpc_uuid", Type: proto.ColumnType_STRING, Description: "A string specifying the UUID of the VPC to which the load balancer is assigned."},

			// Resource interface
			{Name: "akas", Type: proto.ColumnType_JSON, Transform: transform.FromValue().Transform(loadBalancerToURN).Transform(ensureStringArray), Description: resourceInterfaceDescription("akas")},
			{Name: "tags", Type: proto.ColumnType_JSON, Transform: transform.FromField("Tags").Transform(labelsToTagsMap), Description: resourceInterfaceDescription("tags")},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Name"), Description: resourceInterfaceDescription("title")},
		},
	}
}

func listLoadBalancer(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_load_balancer.listLoadBalancer", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		loadBalancers, resp, err := conn.LoadBalancers.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_load_balancer.listLoadBalancer", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, t := range loadBalancers {
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

func getLoadBalancer(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_load_balancer.getLoadBalancer", "connection_error", err)
		return nil, err
	}
	quals := d.KeyColumnQuals
	id := quals["id"].GetStringValue()
	result, resp, err := conn.LoadBalancers.Get(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("digitalocean_load_balancer.getLoadBalancer", "not_found_error", err, "quals", quals, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_load_balancer.getLoadBalancer", "query_error", err, "quals", quals, "resp", resp)
		return nil, err
	}
	return *result, nil
}

func loadBalancerToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := d.Value.(godo.LoadBalancer)
	return "do:loadBalancer:" + i.ID, nil
}
