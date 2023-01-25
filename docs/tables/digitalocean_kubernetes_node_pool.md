# Table: digitalocean_kubernetes_node_pool

A node pool is a group of nodes with the same configuration within a cluster.

Each node within the node pool has a Kubernetes node label which is the node pool’s name. The node pool may be resized up or down to accommodate workloads.

## Examples

### Basic info

```sql
select
  id,
  name,
  cluster_id
  auto_scale
from
  digitalocean_kubernetes_node_pool;
```

### List node pools with autoscaling disabled

```sql
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

### Count numbers of nodes per node pool

```sql
select
  id,
  name,
  cluster_id,
  count as node_count
from
  digitalocean_kubernetes_node_pool;
```

### Get node details of node pools

```sql
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

### Get the top five node pools with the most nodes

```sql
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
