---
title: "Steampipe Table: digitalocean_size - Query DigitalOcean Droplet Sizes using SQL"
description: "Allows users to query DigitalOcean Droplet Sizes, providing insights into available droplet configurations and their specifications."
---

# Table: digitalocean_size - Query DigitalOcean Droplet Sizes using SQL

DigitalOcean Droplet Sizes represent different configurations of CPU, memory, and storage that can be used for Droplets. These configurations determine the hardware of the host machine and have different costs associated with them. Droplet Sizes are predefined and cannot be customized.

## Table Usage Guide

The `digitalocean_size` table provides insights into available configurations for DigitalOcean Droplets. As a system administrator or DevOps engineer, explore droplet size-specific details through this table, including memory, vCPUs, disk size, and transfer limits. Utilize it to understand the specifications and costs of different droplet configurations, aiding in informed decision making for resource allocation and cost management.

## Examples

### List all sizes
Explore the different available options in your DigitalOcean environment, including memory, disk, and CPU specifications, to better understand your current resource utilization and plan for future needs. This can help you manage your resources more effectively and ensure your applications have the resources they need to run smoothly.

```sql+postgres
select
  *
from
  digitalocean_size;
```

```sql+sqlite
select
  *
from
  digitalocean_size;
```

### Most expensive sizes
Analyze the settings to understand the most costly configurations in terms of monthly expenses. This query can be used to identify potential areas for cost optimization by pinpointing the top ten most expensive sizes.

```sql+postgres
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

```sql+sqlite
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
Explore which size options are available in a specific region, like Bangalore, to understand the monthly pricing and make informed decisions for resource allocation. This is helpful in planning your budget and operational needs in the given region.

```sql+postgres
select
  slug,
  price_monthly
from
  digitalocean_size
where
  regions ? 'blr1'
  and available;
```

```sql+sqlite
select
  slug,
  price_monthly
from
  digitalocean_size
where
  json_extract(regions, '$.blr1') is not null
  and available = 1;
```