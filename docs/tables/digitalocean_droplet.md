# Table: digitalocean_droplet

A Droplet is a DigitalOcean virtual machine.

## Examples

### List all droplets

```sql
select
  *
from
  digitalocean_droplet;
```

### Get a droplet by ID

```sql
select
  *
from
  digitalocean_droplet
where
  id = 227211874;
```

### Droplets by region_slug

```sql
select
  region_slug,
  count(id),
  sum(memory) as total_memory
from
  digitalocean_droplet
group by
  region_slug
order by
  region_slug;
```

### Droplets that do not have backups enabled

```sql
select
  name,
  region_slug,
  features
from
  digitalocean_droplet
where
  not features ? 'backups';
```

### Droplet network addresses

```sql
select
  name,
  region_slug,
  private_ipv4,
  public_ipv4,
  public_ipv6
from
  digitalocean_droplet;
```

### Largest droplets

```sql
select
  name,
  region_slug,
  memory
from
  digitalocean_droplet
order by
  memory desc
limit
  10;
```

### Oldest droplets

```sql
select
  name,
  region_slug,
  created_at
from
  digitalocean_droplet
order by
  created_at
limit
  10;
```

### Droplets with tag "production"

```sql
select
  name,
  region_slug,
  tags
from
  digitalocean_droplet
where
  tags ? 'production';
```
