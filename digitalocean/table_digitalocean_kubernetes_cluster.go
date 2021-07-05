package digitalocean

import (
	"context"
	"fmt"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/plugin"
	"github.com/turbot/steampipe-plugin-sdk/plugin/transform"
)

func tableDigitalOceanKubernetesCluster(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_kubernetes_cluster",
		Description: "DigitalOcean Kubernetes (DOKS) is a managed Kubernetes service that lets you deploy Kubernetes clusters without the complexities of handling the control plane and containerized infrastructure. Clusters are compatible with standard Kubernetes toolchains and integrate natively with DigitalOcean Load Balancers and block storage volumes.",
		List: &plugin.ListConfig{
			Hydrate: listKubernetesClusters,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    getKubernetesCluster,
		},
		Columns: []*plugin.Column{
			// Top columns
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique universal identifier of this cluster.",
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The globally unique human-readable name for the cluster.",
			},

			// Other columns
			{
				Name:        "ClusterSubnet",
				Type:        proto.ColumnType_BOOL,
				Description: "If true, all resources will be added to this project if no project is specified.",
			},
			{
				Name:        "Endpoint",
				Type:        proto.ColumnType_STRING,
				Description: "The purpose of the project.",
			},
			{
				Name:        "IPv4",
				Type:        proto.ColumnType_STRING,
				Description: "The unique universal identifier of the project owner.",
			},
			{
				Name:        "RegionSlug",
				Type:        proto.ColumnType_STRING,
				Description: "The description of the project.",
			},
			{
				Name:        "ServiceSubnet",
				Type:        proto.ColumnType_INT,
				Description: "The integer id of the project owner.",
			},
			{
				Name:        "VersionSlug",
				Type:        proto.ColumnType_STRING,
				Description: "The environment of the project's resources.",
			},
			{
				Name:        "VPCUUID",
				Type:        proto.ColumnType_STRING,
				Description: "Time when the project was updated.",
			},

			// Resource interface
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("akas"),
				Transform:   transform.FromValue().Transform(clusterToURN).Transform(ensureStringArray),
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

func listKubernetesClusters(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("listKubernetesClusters", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		clusters, resp, err := conn.Kubernetes.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("listKubernetesClusters", "query_error", err, "opts", opts)
			return nil, err
		}
		for _, cluster := range clusters {
			d.StreamListItem(ctx, cluster)
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

func getKubernetesCluster(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getKubernetesCluster", "connection_error", err)
		return nil, err
	}

	id := d.KeyColumnQuals["id"].GetStringValue()

	result, resp, err := conn.Kubernetes.Get(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("getKubernetesCluster", "not_found_error", err, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("getKubernetesCluster", "query_error", err, "resp", resp)
		return nil, err
	}
	return result, nil
}

func clusterToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := *d.Value.(*godo.KubernetesCluster)
	return fmt.Sprintf("do:kubernetes:%s", i.ID), nil
}
