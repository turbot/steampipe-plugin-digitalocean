---
title: "Steampipe Table: digitalocean_droplet - Query DigitalOcean Droplets using SQL"
description: "Allows users to query DigitalOcean Droplets, providing detailed information about each droplet's configuration and status."
---

# Table: digitalocean_droplet - Query DigitalOcean Droplets using SQL

A DigitalOcean Droplet is a scalable compute platform with add-on storage, security, and monitoring capabilities. Choose the OS or application that best suits your needs and deploy Droplets on the datacenter region closest to your app or users. Droplets are virtual machines available in multiple configurations of CPU, memory, and SSD.

## Table Usage Guide

The `digitalocean_droplet` table provides insights into Droplets within DigitalOcean. As a DevOps engineer, explore Droplet-specific details through this table, including region, size, status, and associated metadata. Utilize it to uncover information about Droplets, such as their current status, the resources they're using, and their location.

## Examples

### List all droplets
Explore all the active droplets in your DigitalOcean account to gain insights into their status and configuration. This can help you manage resources more efficiently and identify any potential issues.

```sql+postgres
select
  *
from
  digitalocean_droplet;
```

```sql+sqlite
select
  *
from
  digitalocean_droplet;
```

### Get a droplet by ID
Determine the specifics of a particular DigitalOcean droplet based on its unique identifier. This is useful for gaining insights into the droplet's configuration and status.

```sql+postgres
select
  *
from
  digitalocean_droplet
where
  id = 227211874;
```

```sql+sqlite
select
  *
from
  digitalocean_droplet
where
  id = 227211874;
```

### Droplets by region_slug
Analyze the distribution of digital ocean droplets across various regions, providing insights into memory allocation and usage patterns in each region. This allows for better resource management and planning for future deployments.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that have not enabled backups on their digital ocean droplets. This is crucial for assessing the potential risk of data loss and implementing necessary safeguards.

```sql+postgres
select
  name,
  region_slug,
  features
from
  digitalocean_droplet
where
  not features ? 'backups';
```

```sql+sqlite
select
  name,
  region_slug,
  features
from
  digitalocean_droplet
where
  json_extract(features, '$.backups') is null;
```

### Droplet network addresses
Explore the network configurations of your DigitalOcean resources to gain insights into their geographical distribution and connectivity details. This helps in better understanding of your resource allocation and planning for future scaling or migration activities.

```sql+postgres
select
  name,
  region_slug,
  private_ipv4,
  public_ipv4,
  public_ipv6
from
  digitalocean_droplet;
```

```sql+sqlite
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
Discover the ten largest droplets in terms of memory within your DigitalOcean environment. This can help you manage resources more effectively and identify potential areas for optimization.

```sql+postgres
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

```sql+sqlite
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
Explore which DigitalOcean droplets were created first to better manage and prioritize system resources or updates. This can be useful in identifying older instances that may require upgrades or maintenance.

```sql+postgres
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

```sql+sqlite
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
Determine the areas in which your digital ocean droplets are tagged as 'production'. This can be useful to quickly identify and manage all production-related resources in your infrastructure.

```sql+postgres
select
  name,
  region_slug,
  tags
from
  digitalocean_droplet
where
  tags ? 'production';
```

```sql+sqlite
select
  name,
  region_slug,
  tags
from
  digitalocean_droplet
where
  json_extract(tags, '$.production') is not null;
```