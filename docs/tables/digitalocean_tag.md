# Table: digitalocean_tag


## Examples

### Tags by name

```sql
select
  name,
  resource_count
from
  digitalocean_tag
order by
  name;
```

### Resource counts by tag

```sql
select
  name,
  coalesce((resources -> 'databases' -> 'count') :: int, 0) as database_count,
  coalesce((resources -> 'droplets' -> 'count') :: int, 0) as droplet_count,
  coalesce((resources -> 'images' -> 'count') :: int, 0) as image_count,
  coalesce((resources -> 'volumes' -> 'count') :: int, 0) as volume_count,
  coalesce((resources -> 'volume_snapshots' -> 'count') :: int, 0) as volume_snapshot_count
from
  digitalocean_tag;
```
