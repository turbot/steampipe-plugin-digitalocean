---
title: "Steampipe Table: digitalocean_snapshot - Query DigitalOcean Snapshots using SQL"
description: "Allows users to query DigitalOcean Snapshots, providing information about the snapshot's ID, name, regions, resource ID, and resource type."
---

# Table: digitalocean_snapshot - Query DigitalOcean Snapshots using SQL

DigitalOcean Snapshots are point-in-time backups of a DigitalOcean Droplet or Volume. They contain all the data and settings from the Droplet or Volume at the moment the Snapshot was taken. Snapshots can be used to restore a Droplet or Volume to the state it was in when the Snapshot was taken, or to create new Droplets or Volumes.

## Table Usage Guide

The `digitalocean_snapshot` table provides insights into the snapshots within DigitalOcean. As a DevOps engineer, explore snapshot-specific details through this table, including the snapshot's ID, name, regions, resource ID, and resource type. Utilize it to uncover information about snapshots, such as the regions they are available in, the resources they are associated with, and the state of those resources at the time the snapshot was taken.

## Examples

### List all snapshots
Explore all snapshots within your DigitalOcean environment to better manage your resources and understand your current usage. This helps in efficient resource allocation and aids in decision-making for future resource planning.

```sql+postgres
select
  *
from
  digitalocean_snapshot;
```

```sql+sqlite
select
  *
from
  digitalocean_snapshot;
```

### Get a snapshot by ID
Analyze the settings of a specific DigitalOcean snapshot to understand its configuration and details. This can be useful in scenarios where you need to assess the elements within a particular snapshot for troubleshooting or optimization purposes.

```sql+postgres
select
  *
from
  digitalocean_snapshot
where
  id = '12005676-5a92-11eb-a53a-0a58ac14663a';
```

```sql+sqlite
select
  *
from
  digitalocean_snapshot
where
  id = '12005676-5a92-11eb-a53a-0a58ac14663a';
```

### Droplet snapshots
Assess the elements within your DigitalOcean environment to understand the size and type of your droplet-based snapshots. This allows for better resource management and planning for storage requirements.

```sql+postgres
select
  name,
  resource_type,
  size_gigabytes
from
  digitalocean_snapshot
where
  resource_type = 'droplet';
```

```sql+sqlite
select
  name,
  resource_type,
  size_gigabytes
from
  digitalocean_snapshot
where
  resource_type = 'droplet';
```

### Largest snapshots
Discover the top ten largest snapshots on DigitalOcean, allowing you to identify potential areas for data management and storage optimization. This can be useful in managing your resources more efficiently and reducing costs.

```sql+postgres
select
  name,
  resource_type,
  size_gigabytes
from
  digitalocean_snapshot
order by
  size_gigabytes desc
limit
  10;
```

```sql+sqlite
select
  name,
  resource_type,
  size_gigabytes
from
  digitalocean_snapshot
order by
  size_gigabytes desc
limit
  10;
```

### Oldest snapshots
Explore which snapshots were created first on your DigitalOcean account to help manage or clean up old resources. This can be especially useful in maintaining storage efficiency and organization.

```sql+postgres
select
  name,
  resource_type,
  created_at
from
  digitalocean_snapshot
order by
  created_at
limit
  10;
```

```sql+sqlite
select
  name,
  resource_type,
  created_at
from
  digitalocean_snapshot
order by
  created_at
limit
  10;
```

### Snapshot with Droplet details
Discover the segments that link snapshots to specific droplets in DigitalOcean to better understand your resource allocation and usage.

```sql+postgres
select
  d.id as droplet_id,
  d.name as droplet_name,
  s.id as snapshot_id,
  s.name as snapshot_name
from
  digitalocean_snapshot as s,
  digitalocean_droplet as d
where
  s.resource_id :: bigint = d.id
  and s.resource_type = 'droplet';
```

```sql+sqlite
select
  d.id as droplet_id,
  d.name as droplet_name,
  s.id as snapshot_id,
  s.name as snapshot_name
from
  digitalocean_snapshot as s,
  digitalocean_droplet as d
where
  s.resource_id = d.id
  and s.resource_type = 'droplet';
```