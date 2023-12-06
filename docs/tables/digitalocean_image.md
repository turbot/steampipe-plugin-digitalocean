---
title: "Steampipe Table: digitalocean_image - Query DigitalOcean Images using SQL"
description: "Allows users to query Images in DigitalOcean, specifically the details of the images such as ID, name, type, distribution, slug, and status."
---

# Table: digitalocean_image - Query DigitalOcean Images using SQL

DigitalOcean Images are pre-configured disk images that can be used to create Droplets or Kubernetes nodes. These images can be manually created by users, automatically generated from existing Droplets, or provided by DigitalOcean. They include base distribution images, one-click application images, snapshots, and backups.

## Table Usage Guide

The `digitalocean_image` table provides insights into Images within DigitalOcean. As a DevOps engineer, explore image-specific details through this table, including type, distribution, slug, and status. Utilize it to uncover information about images, such as those with specific distributions, the details of snapshots and backups, and the verification of image statuses.

## Examples

### List all images
Explore all the available images in your DigitalOcean environment to understand what resources are currently in use. This helps in managing resources efficiently by identifying unused or redundant images.

```sql+postgres
select
  *
from
  digitalocean_image;
```

```sql+sqlite
select
  *
from
  digitalocean_image;
```

### List custom images
Explore which custom images in your DigitalOcean account are not public. This can help maintain privacy and security by ensuring your custom images are not accessible to the general public.

```sql+postgres
select
  slug,
  distribution,
  error_message
from
  digitalocean_image
where
  not public;
```

```sql+sqlite
select
  slug,
  distribution,
  error_message
from
  digitalocean_image
where
  public = 0;
```

### Get Image by ID
Explore which digital ocean image is associated with a specific ID to better manage and organize your resources. This is particularly useful in scenarios where you need to quickly identify and access specific images based on their unique identifiers.

```sql+postgres
select
  id,
  slug,
  name,
  distribution
from
  digitalocean_image
where
  id = 29280599;
```

```sql+sqlite
select
  id,
  slug,
  name,
  distribution
from
  digitalocean_image
where
  id = 29280599;
```

### Get Image by Slug
Discover the segments that use a specific image in your DigitalOcean environment. This is useful to understand where and how frequently certain images are being deployed, aiding in resource management and optimization.

```sql+postgres
select
  id,
  slug,
  name,
  distribution
from
  digitalocean_image
where
  id = 'freebsd-11-x64-zfs';
```

```sql+sqlite
select
  id,
  slug,
  name,
  distribution
from
  digitalocean_image
where
  id = 'freebsd-11-x64-zfs';
```

### Public images by distribution
Analyze the number of public images available for each distribution in DigitalOcean to understand their popularity and usage trends. This can be helpful in making informed decisions about which distributions to support or use based on their community adoption.

```sql+postgres
select
  distribution,
  count(id)
from
  digitalocean_image
group by
  distribution
order by
  count desc;
```

```sql+sqlite
select
  distribution,
  count(id)
from
  digitalocean_image
group by
  distribution
order by
  count(id) desc;
```

### List all backups
Explore which digital images have been backed up, understanding their size and type to manage storage and ensure data security effectively.

```sql+postgres
select
  name,
  size_gigabytes,
  type
from
  digitalocean_image
where
  type = 'backup';
```

```sql+sqlite
select
  name,
  size_gigabytes,
  type
from
  digitalocean_image
where
  type = 'backup';
```