---
title: "Steampipe Table: digitalocean_kubernetes_cluster - Query DigitalOcean Kubernetes Clusters using SQL"
description: "Allows users to query DigitalOcean Kubernetes Clusters, specifically the cluster's ID, name, region, version, and other metadata, providing insights into the cluster's configuration and status."
---

# Table: digitalocean_kubernetes_cluster - Query DigitalOcean Kubernetes Clusters using SQL

DigitalOcean Kubernetes (DOKS) is a managed Kubernetes service that lets you deploy, manage, and scale containerized applications using Kubernetes. It provides developers with the flexibility to ship and scale applications without the overhead of managing the underlying infrastructure. DOKS is integrated with the DigitalOcean developer cloud stack, offering seamless management of Kubernetes clusters.

## Table Usage Guide

The `digitalocean_kubernetes_cluster` table provides insights into Kubernetes clusters within the DigitalOcean cloud platform. As a DevOps engineer, explore cluster-specific details through this table, including cluster version, status, and associated metadata. Utilize it to uncover information about your Kubernetes deployments, such as the clusters' current version, their location, and their current running status.

## Examples

### Basic info
Explore which DigitalOcean Kubernetes clusters are currently active, along with their associated subnet and IP address details. This can be useful for managing and monitoring your cloud resources effectively.

```sql+postgres
select
  id,
  name,
  cluster_subnet,
  ipv4
from
  digitalocean_kubernetes_cluster;
```

```sql+sqlite
select
  id,
  name,
  cluster_subnet,
  ipv4
from
  digitalocean_kubernetes_cluster;
```

### List clusters that are not running
Explore which DigitalOcean Kubernetes clusters are not currently running. This can be useful for identifying potential issues or managing resource allocation.

```sql+postgres
select
  id,
  name,
  cluster_subnet,
  ipv4
from
  digitalocean_kubernetes_cluster
where
  status <> 'running';
```

```sql+sqlite
select
  id,
  name,
  cluster_subnet,
  ipv4
from
  digitalocean_kubernetes_cluster
where
  status <> 'running';
```

### List clusters with auto-upgrade not enabled
Analyze the settings to understand which Kubernetes clusters on DigitalOcean have not enabled the auto-upgrade feature. This can help in ensuring that your systems are always up-to-date with the latest features and security patches.

```sql+postgres
select
  id,
  name,
  cluster_subnet,
  ipv4
from
  digitalocean_kubernetes_cluster
where
  not auto_upgrade;
```

```sql+sqlite
select
  id,
  name,
  cluster_subnet,
  ipv4
from
  digitalocean_kubernetes_cluster
where
  auto_upgrade = 0;
```