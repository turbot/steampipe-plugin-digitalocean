# Table: digitalocean_account

Query information about your current account.

## Examples

### Account information

```sql
select
  *
from
  digitalocean_account;
```

### Check current status of your account

```sql
select
  status,
  status_message
from
  digitalocean_account;
```

### Check usage of limits in your account

```sql
with droplets as (
  select
    count(urn)
  from
    digitalocean_droplet
),
floating_ips as (
  select
    count(urn)
  from
    digitalocean_floating_ip
),
volumes as (
  select
    count(urn)
  from
    digitalocean_volume
)
select
  (
    select
      count
    from
      droplets
  ) as droplet_count,
  droplet_limit,
  round(
    100.0 * (
      select
        count
      from
        droplets
    ) / droplet_limit,
    1
  ) as droplet_usage_percent,
  (
    select
      count
    from
      volumes
  ) as volume_count,
  volume_limit,
  round(
    100.0 * (
      select
        count
      from
        volumes
    ) / volume_limit,
    1
  ) as volume_usage_percent,
  (
    select
      count
    from
      floating_ips
  ) as floating_ip_count,
  floating_ip_limit,
  round(
    100.0 * (
      select
        count
      from
        floating_ips
    ) / floating_ip_limit,
    1
  ) as floating_ip_usage_percent
from
  digitalocean_account;
```
