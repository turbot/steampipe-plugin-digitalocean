# Table: digitalocean_image

A DigitalOcean image can be used to create a Droplet and may come in a number
of flavors. Currently, there are five types of images: snapshots, backups,
applications, distributions, and custom images.

## Examples

### List all images

```sql
select
  *
from
  digitalocean_image;
```

### List custom images

```sql
select
  slug,
  distribution,
  error_message
from
  digitalocean_image
where
  not public;
```

### Get Image by ID

```sql
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

```sql
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

```sql
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

### List all backups

```sql
select
  name,
  size_gigabytes,
  type
from
  digitalocean_image
where
  type = 'backup';
```
