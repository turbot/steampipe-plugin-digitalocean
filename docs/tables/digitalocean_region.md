---
title: "Steampipe Table: digitalocean_region - Query DigitalOcean Regions using SQL"
description: "Allows users to query DigitalOcean Regions, providing details such as slug, name, features, and sizes available in these regions."
---

# Table: digitalocean_region - Query DigitalOcean Regions using SQL

DigitalOcean Regions represent the geographical locations where your Droplets and other resources reside. Each region is a separate geographic area, and each region generally has multiple, isolated locations known as availability zones. Regions are designed to allow users to place resources, like Droplets and Spaces, closer to customers for reduced latency.

## Table Usage Guide

The `digitalocean_region` table provides insights into the Regions within DigitalOcean. As a DevOps engineer, explore region-specific details through this table, including the slug, name, features, and sizes available in these regions. Utilize it to understand the geographical distribution of your resources and the features available in each region.

## Examples

### List all regions
Explore the various regions available in your DigitalOcean account. This is useful for planning resource distribution and managing data residency requirements.

```sql+postgres
select
  *
from
  digitalocean_region;
```

```sql+sqlite
select
  *
from
  digitalocean_region;
```

### New York regions
Explore regions within the DigitalOcean service that are based in New York to better manage resources or optimize network latency for local users.

```sql+postgres
select
  *
from
  digitalocean_region
where
  slug like 'ny%';
```

```sql+sqlite
select
  *
from
  digitalocean_region
where
  slug like 'ny%';
```

### Regions available for new droplets
Explore which regions are currently available for deploying new droplets on DigitalOcean. This can be helpful for planning deployments and ensuring optimal location selection.

```sql+postgres
select
  slug,
  name,
  available
from
  digitalocean_region
where
  available;
```

```sql+sqlite
select
  slug,
  name,
  available
from
  digitalocean_region
where
  available = 1;
```

### Regions where the storage feature is available
Discover the regions where the storage feature is available. This can be useful for understanding where you can deploy resources that require this feature.

```sql+postgres
select
  slug,
  name,
  features
from
  digitalocean_region
where
  features ? 'storage';
```

```sql+sqlite
Error: SQLite does not support the '?' operator for JSON objects.
```