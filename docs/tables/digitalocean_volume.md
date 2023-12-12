---
title: "Steampipe Table: digitalocean_volume - Query DigitalOcean Volumes using SQL"
description: "Allows users to query DigitalOcean Volumes, specifically providing insights into volume details, including size, region, and droplet attachments."
---

# Table: digitalocean_volume - Query DigitalOcean Volumes using SQL

DigitalOcean Volumes are highly available and reliable storage volumes that you can attach to your DigitalOcean Droplets. Volumes are network-based block storage that provide additional data storage for your Droplets. They are independent entities that you can move between Droplets within the same datacenter.

## Table Usage Guide

The `digitalocean_volume` table provides insights into the storage volumes within DigitalOcean. As a DevOps engineer, explore volume-specific details through this table, including size, region, and droplet attachments. Utilize it to manage and optimize your storage resources, such as identifying unattached volumes or volumes with low utilization.

## Examples

### List all volumes
Explore all the storage volumes available in your DigitalOcean account to understand your current storage usage and plan for future needs. This query is useful for managing your resources effectively and avoiding potential storage shortages.

```sql+postgres
select
  *
from
  digitalocean_volume;
```

```sql+sqlite
select
  *
from
  digitalocean_volume;
```

### Get a volume by ID
Discover the details of a specific storage volume in your DigitalOcean environment using its unique ID. This can be useful for troubleshooting or auditing purposes, to understand the settings and configuration of a particular volume.

```sql+postgres
select
  *
from
  digitalocean_volume
where
  id = '12005676-5a92-11eb-a53a-0a58ac14663a';
```

```sql+sqlite
select
  *
from
  digitalocean_volume
where
  id = '12005676-5a92-11eb-a53a-0a58ac14663a';
```

### Volumes by region
Analyze the distribution of storage volumes across different regions to understand resource allocation and usage patterns. This can aid in identifying regions with high storage usage and help in strategic planning for resource provisioning.

```sql+postgres
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

```sql+sqlite
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
Explore which digital ocean volumes are the largest in terms of gigabytes across different regions. This can be useful for managing storage resources and identifying areas that may require capacity adjustments.

```sql+postgres
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

```sql+sqlite
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
Identify the oldest storage volumes in your DigitalOcean account to assess whether they're still needed or if they can be deleted to save costs. This query helps in managing resources effectively by highlighting potential areas for cleanup and cost savings.

```sql+postgres
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

```sql+sqlite
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
Discover the segments that are tagged as 'production' within the DigitalOcean platform, allowing you to focus on areas of your business that are in active use or deployment.

```sql+postgres
select
  name,
  region_name,
  tags
from
  digitalocean_volume
where
  tags ? 'production';
```

```sql+sqlite
Error: SQLite does not support the '?' operator for JSON objects.
```