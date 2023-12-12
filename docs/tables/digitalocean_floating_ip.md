---
title: "Steampipe Table: digitalocean_floating_ip - Query DigitalOcean Floating IPs using SQL"
description: "Allows users to query Floating IPs in DigitalOcean, specifically the public IP addresses that can be assigned to Droplets, providing insights into IP allocation and usage."
---

# Table: digitalocean_floating_ip - Query DigitalOcean Floating IPs using SQL

A Floating IP is a public IP address that can be statically assigned to a Droplet. Unlike the standard public IP addresses, Floating IPs are retained even when a Droplet is destroyed, and can be quickly remapped to another Droplet within the same datacenter. This feature is particularly useful for directing network traffic and ensuring minimal downtime.

## Table Usage Guide

The `digitalocean_floating_ip` table provides insights into the Floating IPs within DigitalOcean. As a network administrator, explore IP-specific details through this table, including the region, Droplet ID, and associated metadata. Utilize it to uncover information about IP allocation, such as which Droplets are associated with specific IPs, and the distribution of IPs across different regions.

## Examples

### List all Floating IPs
Discover the segments that have floating IP addresses assigned to them, including the associated droplet names and regions. This is useful to manage and monitor the distribution and usage of floating IPs across different regions and droplets.

```sql+postgres
select
  ip,
  droplet ->> 'name' as droplet_name,
  region ->> 'slug' as region_slug
from
  digitalocean_floating_ip;
```

```sql+sqlite
select
  ip,
  json_extract(droplet, '$.name') as droplet_name,
  json_extract(region, '$.slug') as region_slug
from
  digitalocean_floating_ip;
```

### Get a Floating IP by IP
Discover the segments that have a specific IP address by identifying the associated droplet name and region. This is useful to understand the distribution and usage of floating IPs within your digital ocean infrastructure.

```sql+postgres
select
  ip,
  droplet ->> 'name' as droplet_name,
  region ->> 'slug' as region_slug
from
  digitalocean_floating_ip
where
  ip = '161.35.249.180';
```

```sql+sqlite
select
  ip,
  json_extract(droplet, '$.name') as droplet_name,
  json_extract(region, '$.slug') as region_slug
from
  digitalocean_floating_ip
where
  ip = '161.35.249.180';
```

### List all Floating IPs in New York regions
Explore which floating IPs are associated with the New York regions in your DigitalOcean account. This allows you to easily manage and allocate your resources within specific geographic locations.

```sql+postgres
select
  ip,
  droplet ->> 'name' as droplet_name,
  region ->> 'slug' as region_slug
from
  digitalocean_floating_ip
where
  region ->> 'slug' like 'ny%';
```

```sql+sqlite
select
  ip,
  json_extract(droplet, '$.name') as droplet_name,
  json_extract(region, '$.slug') as region_slug
from
  digitalocean_floating_ip
where
  json_extract(region, '$.slug') like 'ny%';
```