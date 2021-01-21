# Table: digitalocean_size

The sizes objects represent different packages of hardware resources that can
be used for Droplets. This includes the amount of RAM, the number of virtual
CPUs, disk space, and transfer. The size object also includes the pricing
details and the regions that the size is available in.

## Examples

### List all sizes

```sql
select
  *
from
  digitalocean_size;
```

### Most expensive sizes

```sql
select
  slug,
  vcpus,
  memory,
  disk,
  price_hourly,
  price_monthly
from
  digitalocean_size
order by
  price_monthly desc
limit
  10;
```

### Sizes available in Bangalore

```sql
select
  slug,
  price_monthly
from
  digitalocean_size
where
  regions ? 'blr1'
  and available;
```
