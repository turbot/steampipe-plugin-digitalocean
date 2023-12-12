---
title: "Steampipe Table: digitalocean_tag - Query DigitalOcean Tags using SQL"
description: "Allows users to query Tags in DigitalOcean, specifically the tag name and resource count, providing insights into resource tagging patterns and potential management needs."
---

# Table: digitalocean_tag - Query DigitalOcean Tags using SQL

A Tag in DigitalOcean is a label that can be applied to various resources to organize and manage them more effectively. Tags can be applied to droplets, images, volumes, volume snapshots, databases, and load balancers. They help to categorize and filter resources based on user-defined parameters. 

## Table Usage Guide

The `digitalocean_tag` table provides insights into the tags within DigitalOcean. As an IT administrator, explore tag-specific details through this table, including the tag names and resource counts. Utilize it to uncover information about tags, such as those with a high number of associated resources, enabling better resource management and categorization.

## Examples

### Tags by name
Discover the segments that are tagged within your DigitalOcean resources and analyze the frequency of each tag's usage. This can be useful in understanding how your resources are categorized and managed.

```sql+postgres
select
  name,
  resource_count
from
  digitalocean_tag
order by
  name;
```

```sql+sqlite
select
  name,
  resource_count
from
  digitalocean_tag
order by
  name;
```

### Resource counts by tag
Determine the areas in which resources are tagged and gain insights into the distribution of different resource types such as databases, droplets, images, volumes, and volume snapshots. This can aid in resource management and allocation within a DigitalOcean environment.

```sql+postgres
select
  name,
  coalesce((resources -> 'databases' -> 'count') :: int, 0) as database_count,
  coalesce((resources -> 'droplets' -> 'count') :: int, 0) as droplet_count,
  coalesce((resources -> 'images' -> 'count') :: int, 0) as image_count,
  coalesce((resources -> 'volumes' -> 'count') :: int, 0) as volume_count,
  coalesce((resources -> 'volume_snapshots' -> 'count') :: int, 0) as volume_snapshot_count
from
  digitalocean_tag;
```

```sql+sqlite
select
  name,
  coalesce(cast(json_extract(resources, '$.databases.count') as integer), 0) as database_count,
  coalesce(cast(json_extract(resources, '$.droplets.count') as integer), 0) as droplet_count,
  coalesce(cast(json_extract(resources, '$.images.count') as integer), 0) as image_count,
  coalesce(cast(json_extract(resources, '$.volumes.count') as integer), 0) as volume_count,
  coalesce(cast(json_extract(resources, '$.volume_snapshots.count') as integer), 0) as volume_snapshot_count
from
  digitalocean_tag;
```