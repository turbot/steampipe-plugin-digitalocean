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

### Droplets by region

```sql
select
  region,
  count(id),
  sum(size_gigabytes) as size_gigabytes
from
  digitalocean_droplet
group by
  region
order by
  region;
```

### Droplets that do not have backups enabled

```sql
select
  name,
  region,
  features
from
  digitalocean_droplet
where
  not features ? 'backups';
```


### Largest droplets

```sql
select
  name,
  region,
  size_gigabytes
from
  digitalocean_droplet
order by
  size_gigabytes desc
limit
  10;
```

### Oldest droplets

```sql
select
  name,
  region,
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
  region,
  tags
from
  digitalocean_droplet
where
  tags ? 'production';
```
