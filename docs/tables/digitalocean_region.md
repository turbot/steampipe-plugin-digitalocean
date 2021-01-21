# Table: digitalocean_region

A region in DigitalOcean represents a datacenter where Droplets can be deployed
and images can be transferred. Each region represents a specific datacenter in
a geographic location. Some geographical locations may have multiple "regions"
available. This means that there are multiple datacenters available within that
area.

## Examples

### List all regions

```sql
select
  *
from
  digitalocean_region;
```

### New York regions

```sql
select
  *
from
  digitalocean_region
where
  slug like 'ny%';
```

### Regions available for new droplets

```sql
select
  slug,
  name,
  available
from
  digitalocean_region
where
  available;
```

### Regions where the storage feature is available

```sql
select
  slug,
  name,
  features
from
  digitalocean_region
where
  features ? 'storage';
```
