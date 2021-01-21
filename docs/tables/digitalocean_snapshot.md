# Table: digitalocean_snapshot

Snapshots are saved instances of a Droplet or a block storage snapshot.

## Examples

### List all snapshots

```sql
select
  *
from
  digitalocean_snapshot;
```

### Get a snapshot by ID

```sql
select
  *
from
  digitalocean_snapshot
where
  id = '12005676-5a92-11eb-a53a-0a58ac14663a';
```

### Droplet snapshots

```sql
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

```sql
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

```sql
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

```sql
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
