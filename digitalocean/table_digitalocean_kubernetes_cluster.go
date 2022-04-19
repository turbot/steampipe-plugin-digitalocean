package digitalocean

import (
	"context"
	"fmt"
	"strings"

	"github.com/digitalocean/godo"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

//// TABLE DEFINITION

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
			{
				Name:        "status",
				Type:        proto.ColumnType_STRING,
				Description: "A string indicating the current status of the cluster. Potential values include running, provisioning, and errored.",
				Transform:   transform.FromField("Status.State"),
			},
			{
				Name:        "created_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date and time when the Kubernetes cluster was created.",
			},
			{
				Name:        "auto_upgrade",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean value indicating whether the cluster will be automatically upgraded to new patch releases during its maintenance window.",
			},
			{
				Name:        "cluster_subnet",
				Type:        proto.ColumnType_STRING,
				Description: "The range of IP addresses in the overlay network of the Kubernetes cluster.",
			},
			{
				Name:        "endpoint",
				Type:        proto.ColumnType_STRING,
				Description: "The base URL of the API server on the Kubernetes master node.",
			},
			{
				Name:        "ipv4",
				Type:        proto.ColumnType_STRING,
				Description: "The public IPv4 address of the Kubernetes master node.",
				Transform:   transform.FromField("IPv4"),
			},
			{
				Name:        "region_slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug identifier for the region where the Kubernetes cluster will be created.",
				Transform:   transform.FromField("RegionSlug"),
			},
			{
				Name:        "registry_enabled",
				Type:        proto.ColumnType_BOOL,
				Description: "A boolean value indicating whether cluster integrated with container registry.",
			},
			{
				Name:        "service_subnet",
				Type:        proto.ColumnType_STRING,
				Description: "The range of assignable IP addresses for services running in the Kubernetes cluster.",
			},
			{
				Name:        "surge_upgrade",
				Type:        proto.ColumnType_BOOL,
				Description: "Enable/disable surge upgrades for a cluster.",
			},
			{
				Name:        "updated_at",
				Type:        proto.ColumnType_TIMESTAMP,
				Description: "The date and time when the Kubernetes cluster was last updated.",
			},
			{
				Name:        "urn",
				Type:        proto.ColumnType_STRING,
				Description: "The uniform resource name (URN) for the cluster.",
				Transform:   transform.FromValue().Transform(clusterToURN),
			},
			{
				Name:        "version_slug",
				Type:        proto.ColumnType_STRING,
				Description: "The slug identifier for the version of Kubernetes used for the cluster.",
				Transform:   transform.FromField("VersionSlug"),
			},
			{
				Name:        "vpc_uuid",
				Type:        proto.ColumnType_STRING,
				Description: "The ID of the VPC where the Kubernetes cluster will be located.",
				Transform:   transform.FromField("VPCUUID"),
			},
			{
				Name:        "maintenance_policy",
				Type:        proto.ColumnType_JSON,
				Description: "A block representing the cluster's maintenance window.",
			},
			{
				Name:        "node_pools",
				Type:        proto.ColumnType_JSON,
				Description: "The cluster's default node pool.",
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
				Transform:   transform.FromValue().Transform(clusterToURN).Transform(ensureStringArray),
			},
		},
	}
}

//// LIST FUNCTION

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

//// HYDRATE FUNCTIONS

func getKubernetesCluster(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getKubernetesCluster", "connection_error", err)
		return nil, err
	}

	id := d.KeyColumnQuals["id"].GetStringValue()

	// Handle empty id
	if id == "" {
		return nil, nil
	}

	result, resp, err := conn.Kubernetes.Get(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), ": 404") {
			plugin.Logger(ctx).Warn("getKubernetesCluster", "not_found_error", err, "resp", resp)
			return nil, nil
		}
		if strings.Contains(err.Error(), ": 400") {
			plugin.Logger(ctx).Warn("getKubernetesCluster", "invalid_id", err, "resp", resp)
			return nil, nil
		}
		plugin.Logger(ctx).Error("getKubernetesCluster", "query_error", err, "resp", resp)
		return nil, err
	}
	return result, nil
}

//// TRANSFORM FUNCTION

func clusterToURN(_ context.Context, d *transform.TransformData) (interface{}, error) {
	i := *d.Value.(*godo.KubernetesCluster)
	return fmt.Sprintf("do:kubernetesCluster:%s", i.ID), nil
}
