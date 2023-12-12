---
title: "Steampipe Table: digitalocean_account - Query DigitalOcean Accounts using SQL"
description: "Allows users to query DigitalOcean Accounts, specifically the account's details such as email, status, droplet limit, and more."
---

# Table: digitalocean_account - Query DigitalOcean Accounts using SQL

DigitalOcean is a Cloud Infrastructure Provider that offers scalable compute platforms with add-on storage, security, and monitoring capabilities. Users can create, automate, manage, and scale cloud applications efficiently on DigitalOcean. The account resource in DigitalOcean provides information about the user's account.

## Table Usage Guide

The `digitalocean_account` table provides insights into the user accounts in DigitalOcean. As a cloud engineer or system administrator, you can explore account-specific details through this table, such as account status, droplet limit, email, and more. Utilize it to manage and monitor your DigitalOcean accounts effectively and efficiently.

## Examples

### Account information
Explore your DigitalOcean account details to better understand your account settings and configurations. This can be beneficial when managing your resources or when you need to verify your account information.

```sql+postgres
select
  *
from
  digitalocean_account;
```

```sql+sqlite
select
  *
from
  digitalocean_account;
```

### Check current status of your account
Explore the current status of your DigitalOcean account to understand any potential issues or updates. This could be useful for troubleshooting or maintaining account health.

```sql+postgres
select
  status,
  status_message
from
  digitalocean_account;
```

```sql+sqlite
select
  status,
  status_message
from
  digitalocean_account;
```

### Check usage of limits in your account
Analyze the utilization of your account's resources to understand how close you are to hitting your limits. This is useful for efficient resource management and avoiding potential disruptions due to exceeding limits.

```sql+postgres
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

```sql+sqlite
select
  (
    select
      count(urn)
    from
      digitalocean_droplet
  ) as droplet_count,
  droplet_limit,
  round(
    100.0 * (
      select
        count(urn)
      from
        digitalocean_droplet
    ) / droplet_limit,
    1
  ) as droplet_usage_percent,
  (
    select
      count(urn)
    from
      digitalocean_volume
  ) as volume_count,
  volume_limit,
  round(
    100.0 * (
      select
        count(urn)
      from
        digitalocean_volume
    ) / volume_limit,
    1
  ) as volume_usage_percent,
  (
    select
      count(urn)
    from
      digitalocean_floating_ip
  ) as floating_ip_count,
  floating_ip_limit,
  round(
    100.0 * (
      select
        count(urn)
      from
        digitalocean_floating_ip
    ) / floating_ip_limit,
    1
  ) as floating_ip_usage_percent
from
  digitalocean_account;
```