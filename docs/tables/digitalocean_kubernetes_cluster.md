# Table: digitalocean_kubernetes_cluster

DigitalOcean Kubernetes (DOKS) is a managed Kubernetes service that lets you deploy Kubernetes clusters without the complexities of handling the control plane and containerized infrastructure. Clusters are compatible with standard Kubernetes toolchains and integrate natively with DigitalOcean Load Balancers and block storage volumes.

## Examples

### Basic info

```sql
select
  id,
  name,
  cluster_subnet,
  ipv4
from
  digitalocean_kubernetes_cluster;
```

### List clusters that are not running

```sql
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

```sql
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
