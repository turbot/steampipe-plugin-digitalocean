package digitalocean

import (
	"context"
	"fmt"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

//// TABLE DEFINITION

func tableDigitalOceanKubernetesNodePool(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "digitalocean_kubernetes_node_pool",
		Description: "DigitalOcean Kubernetes Node Pool",
		List: &plugin.ListConfig{
			ParentHydrate: listKubernetesClusters,
			Hydrate:       listKubernetesNodePools,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"id", "cluster_id"}),
			Hydrate:    getKubernetesNodePool,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique universal identifier of this node pool.",
				Transform:   transform.FromField("ID"),
			},
			{
				Name:        "cluster_id",
				Type:        proto.ColumnType_STRING,
				Description: "The unique universal identifier of the associated cluster.",
				Transform:   transform.FromField("ClusterID"),
			},
			{
				Name:        "name",
				Type:        proto.ColumnType_STRING,
				Description: "The globally unique human-readable name for the node pool.",
				Transform:   transform.FromField("Name"),
			},
			{
				Name:        "auto_scale",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean value indicating whether the node pool has autoscaling enabled.",
				Transform:   transform.FromField("AutoScale"),
			},
			{
				Name:        "count",
				Type:        proto.ColumnType_INT,
				Description: "The number of nodes in the node pool.",
				Transform:   transform.FromField("Count"),
			},
			{
				Name:        "max_nodes",
				Type:        proto.ColumnType_INT,
				Description: "The maximum number of nodes allowed in the node pool.",
				Transform:   transform.FromField("MaxNodes"),
			},
			{
				Name:        "min_nodes",
				Type:        proto.ColumnType_INT,
				Description: "The minimum number of nodes allowed in the node pool.",
				Transform:   transform.FromField("MinNodes"),
			},
			{
				Name:        "size",
				Type:        proto.ColumnType_STRING,
				Description: "The size of the node pool.",
				Transform:   transform.FromField("Size"),
			},
			{
				Name:        "urn",
				Type:        proto.ColumnType_STRING,
				Description: "The uniform resource name (URN) for the node pool.",
				Transform:   transform.FromValue().Transform(nodePoolToURN),
			},

			{
				Name:        "labels",
				Type:        proto.ColumnType_JSON,
				Description: "The labels for the node pool.",
				Transform:   transform.FromField("Labels"),
			},
			{
				Name:        "nodes",
				Type:        proto.ColumnType_JSON,
				Description: "The nodes available in the node pool.",
				Transform:   transform.FromField("Nodes"),
			},
			{
				Name:        "taints",
				Type:        proto.ColumnType_JSON,
				Description: "The taints of the node pool.",
				Transform:   transform.FromField("Taints"),
			},

			// Steampipe standard columns
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
			{
				Name:        "akas",
				Type:        proto.ColumnType_JSON,
				Description: resourceInterfaceDescription("akas"),
				Transform:   transform.FromValue().Transform(nodePoolToURN).Transform(ensureStringArray),
			},
		}),
	}
}

type KubernetesNodePoolInfo struct {
	godo.KubernetesNodePool
	ClusterID string
}

//// LIST FUNCTION

func listKubernetesNodePools(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	cluster := h.Item.(*godo.KubernetesCluster)

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_kubernetes_node_pool.listKubernetesNodePools", "connection_error", err)
		return nil, err
	}
	opts := &godo.ListOptions{
		Page:    1,
		PerPage: 100,
	}
	for {
		nodePools, resp, err := conn.Kubernetes.ListNodePools(ctx, cluster.ID, opts)
		if err != nil {
			plugin.Logger(ctx).Error("digitalocean_kubernetes_node_pool.listKubernetesNodePools", "query_error", err, "id", cluster.ID)
			return nil, err
		}
		for _, nodePool := range nodePools {
			d.StreamListItem(ctx, &KubernetesNodePoolInfo{*nodePool, cluster.ID})
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

func getKubernetesNodePool(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	id := d.EqualsQuals["id"].GetStringValue()
	clusterId := d.EqualsQuals["cluster_id"].GetStringValue()

	// Handle empty id
	if id == "" || clusterId == "" {
		return nil, nil
	}

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("digitalocean_kubernetes_node_pool.getKubernetesNodePool", "connection_error", err)
		return nil, err
	}

	result, resp, err := conn.Kubernetes.GetNodePool(ctx, clusterId, id)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			return nil, nil
		}
		plugin.Logger(ctx).Error("digitalocean_kubernetes_node_pool.getKubernetesNodePool", "query_error", err, "resp", resp)
		return nil, err
	}
	return &KubernetesNodePoolInfo{*result, clusterId}, nil
}

//// TRANSFORM FUNCTION

func nodePoolToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := *d.Value.(*KubernetesNodePoolInfo)
	return fmt.Sprintf("do:kubernetesNodePool:%s", i.ID), nil
}
