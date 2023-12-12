---
title: "Steampipe Table: digitalocean_container_registry - Query DigitalOcean Container Registries using SQL"
description: "Allows users to query Container Registries in DigitalOcean, particularly the details of each registry including its name, server URL, and creation time."
---

# Table: digitalocean_container_registry - Query DigitalOcean Container Registries using SQL

DigitalOcean Container Registry is a service within DigitalOcean that allows you to securely store and distribute Docker images. It provides a fully managed and scalable infrastructure for storing and accessing Docker images, allowing you to deploy new instances and services more quickly. Container Registry helps you manage images throughout the application lifecycle, maintain version control, and facilitate team collaboration.

## Table Usage Guide

The `digitalocean_container_registry` table provides insights into Container Registries within DigitalOcean. As a DevOps engineer, explore registry-specific details through this table, including the name, server URL, and creation time. Utilize it to manage and monitor your Docker images, ensure the security and accessibility of your registries, and streamline your application deployment process.

## Examples

### Basic info
Explore the creation timeline of your digital ocean container registries. This will help you understand when each registry was established, providing insights into your resource allocation and usage history.

```sql+postgres
select
  name,
  urn,
  created_at
from
  digitalocean_container_registry;
```

```sql+sqlite
select
  name,
  urn,
  created_at
from
  digitalocean_container_registry;
```

### Get container registry details created in last 30 days
Discover the details of newly created container registries within the past month. This is useful for tracking recent activity and understanding the storage usage of these registries.

```sql+postgres
select
  name,
  urn,
  created_at,
  storage_usage_bytes_updated_at,
  storage_usage_bytes
from
  digitalocean_container_registry
where
  created_at >= now() - interval '30' day;
```

```sql+sqlite
select
  name,
  urn,
  created_at,
  storage_usage_bytes_updated_at,
  storage_usage_bytes
from
  digitalocean_container_registry
where
  created_at >= datetime('now', '-30 day');
```