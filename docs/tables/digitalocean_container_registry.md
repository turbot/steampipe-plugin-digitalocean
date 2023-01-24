# Table: digitalocean_container_registry

The DigitalOcean Container Registry (DOCR) is a private Docker image registry with additional tooling support that enables integration with your Docker environment and DigitalOcean Kubernetes clusters. DOCR registries are private and co-located in the datacenters where DigitalOcean Kubernetes clusters are operated for secure, stable, and performant rollout of images to your clusters.

## Examples

### Basic info

```sql
select
  name,
  urn,
  created_at
from
  digitalocean_container_registry;
```

### Get container registry details created in last 30 days

```sql
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