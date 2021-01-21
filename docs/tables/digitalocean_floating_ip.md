# Table: digitalocean_floating_ip

DigitalOcean Floating IPs are publicly-accessible static IP addresses that can
be mapped to one of your Droplets. They can be used to create highly available
setups or other configurations requiring movable addresses.

## Examples

### List all Floating IPs

```sql
select
  ip,
  droplet ->> 'name' as droplet_name,
  region ->> 'slug' as region_slug
from
  digitalocean_floating_ip;
```

### Get a Floating IP by IP

```sql
select
  ip,
  droplet ->> 'name' as droplet_name,
  region ->> 'slug' as region_slug
from
  digitalocean_floating_ip
where
  ip = '161.35.249.180';
```

### List all Floating IPs in New York regions

```sql
select
  ip,
  droplet ->> 'name' as droplet_name,
  region ->> 'slug' as region_slug
from
  digitalocean_floating_ip
where
  region ->> 'slug' like 'ny%';
```
