---
title: "Steampipe Table: digitalocean_kubernetes_node_pool - Query DigitalOcean Kubernetes Node Pools using SQL"
description: "Allows users to query DigitalOcean Kubernetes Node Pools, providing detailed information about each node pool within the Kubernetes clusters."
---

# Table: digitalocean_kubernetes_node_pool - Query DigitalOcean Kubernetes Node Pools using SQL

A Kubernetes Node Pool in DigitalOcean is a subset of machines within a Kubernetes cluster that have similar configurations. Node Pools enable you to create and manage multiple groups of nodes within the same Kubernetes cluster, each tailored to the specific compute requirements of your workloads. It provides a way to manage different types of workloads in a Kubernetes cluster more efficiently.

## Table Usage Guide

The `digitalocean_kubernetes_node_pool` table provides insights into Kubernetes Node Pools within DigitalOcean. As a DevOps engineer, you can explore node pool-specific details through this table, including configurations, node count, and associated metadata. Use it to manage and monitor the health and performance of your node pools, ensuring optimal operation of your Kubernetes workloads.


## Examples

### Basic info
Analyze the settings to understand the configuration of DigitalOcean Kubernetes node pools. This can help in managing resources and scaling strategies more effectively.

```sql+postgres
select
  id,
  name,
  cluster_id
  auto_scale
from
  digitalocean_kubernetes_node_pool;
```

```sql+sqlite
select
  id,
  name,
  cluster_id,
  auto_scale
from
  digitalocean_kubernetes_node_pool;
```

### List node pools with autoscaling disabled
Explore which node pools within your Kubernetes cluster have autoscaling disabled. This is beneficial for understanding your resource utilization and identifying potential areas for optimization.

```sql+postgres
select
  id,
  name,
  cluster_id,
  auto_scale
from
  digitalocean_kubernetes_node_pool
where
  not auto_scale;
```

```sql+sqlite
select
  id,
  name,
  cluster_id,
  auto_scale
from
  digitalocean_kubernetes_node_pool
where
  auto_scale = 0;
```

### Count numbers of nodes per node pool
Explore the distribution of nodes within your DigitalOcean Kubernetes clusters. This query helps to balance workload by showing the number of nodes in each node pool.

```sql+postgres
select
  id,
  name,
  cluster_id,
  count as node_count
from
  digitalocean_kubernetes_node_pool;
```

```sql+sqlite
select
  id,
  name,
  cluster_id,
  count as node_count
from
  digitalocean_kubernetes_node_pool;
```

### Get node details of node pools
Discover the details of your node pools, including creation and update times, to understand their status and manage them more effectively. This can be particularly useful in managing resources and troubleshooting issues within your digital ocean Kubernetes environment.

```sql+postgres
select
  p.id,
  p.name,
  n ->> 'created_at' as node_created_at,
  n ->> 'droplet_id' as node_droplet_id,
  n ->> 'id' as node_id,
  n ->> 'name' as node_name,
  n ->> 'status' as node_status,
  n ->> 'updated_at' as node_updated_at
from
  digitalocean_kubernetes_node_pool as p,
  jsonb_array_elements(nodes) as n;
```

```sql+sqlite
select
  p.id,
  p.name,
  json_extract(n.value, '$.created_at') as node_created_at,
  json_extract(n.value, '$.droplet_id') as node_droplet_id,
  json_extract(n.value, '$.id') as node_id,
  json_extract(n.value, '$.name') as node_name,
  json_extract(n.value, '$.status') as node_status,
  json_extract(n.value, '$.updated_at') as node_updated_at
from
  digitalocean_kubernetes_node_pool as p,
  json_each(p.nodes) as n;
```

### Get the top five node pools with the most nodes
Analyze the settings to understand which five node pools in your DigitalOcean Kubernetes service have the highest number of nodes. This can help manage resources by identifying areas of high resource concentration.

```sql+postgres
select
  id,
  name,
  cluster_id,
  max_nodes
from
  digitalocean_kubernetes_node_pool
order by
  max_nodes desc
limit 5;
```

```sql+sqlite
select
  id,
  name,
  cluster_id,
  max_nodes
from
  digitalocean_kubernetes_node_pool
order by
  max_nodes desc
limit 5;
```

### Get cluster details for the node pools
Determine the status and endpoint details of your DigitalOcean Kubernetes clusters by examining the associated node pools. This can aid in monitoring cluster health and connectivity.

```sql+postgres
select
  n.id as node_pool_id,
  n.name,
  n.cluster_id,
  c.status as cluster_status,
  c.cluster_subnet,
  c.endpoint as cluster_endpoint
from
  digitalocean_kubernetes_node_pool as n,
  digitalocean_kubernetes_cluster as c
where
  c.id = cluster_id;
```

```sql+sqlite
select
  n.id as node_pool_id,
  n.name,
  n.cluster_id,
  c.status as cluster_status,
  c.cluster_subnet,
  c.endpoint as cluster_endpoint
from
  digitalocean_kubernetes_node_pool as n,
  digitalocean_kubernetes_cluster as c
where
  c.id = n.cluster_id;
````