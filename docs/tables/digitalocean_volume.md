# Table: digitalocean_volume

DigitalOcean Block Storage Volumes provide expanded storage capacity for your
Droplets and can be moved between Droplets within a specific region. Volumes
function as raw block devices, meaning they appear to the operating system as
locally attached storage which can be formatted using any file system supported
by the OS.

## Examples

### List all volumes

```sql
select
  *
from
  digitalocean_volume;
```

### Get a volume by ID

```sql
select
  *
from
  digitalocean_volume
where
  id = '12005676-5a92-11eb-a53a-0a58ac14663a';
```

### Volumes by region

```sql
select
  region_name,
  count(id),
  sum(size_gigabytes) as size_gigabytes
from
  digitalocean_volume
group by
  region_name
order by
  region_name;
```

### Largest volumes

```sql
select
  name,
  region_name,
  size_gigabytes
from
  digitalocean_volume
order by
  size_gigabytes desc
limit
  10;
```

### Oldest volumes

```sql
select
  name,
  region_name,
  created_at
from
  digitalocean_volume
order by
  created_at
limit
  10;
```

### Volumes with tag "production"

```sql
select
  name,
  region_name,
  tags
from
  digitalocean_volume
where
  tags ? 'production';
```
