# Table: digitalocean_project

Projects allow you to organize your resources into groups that fit the way you
work. You can group resources (like Droplets, Spaces, load balancers, domains,
and floating IPs) in ways that align with the applications you host on
DigitalOcean.

## Examples

### List all projects

```sql
select
  *
from
  digitalocean_project;
```

### Get a project by ID

```sql
select
  *
from
  digitalocean_project
where
  id = '59137997-528e-45ef-9521-db041f2c7d94';
```

### Get the default project

```sql
select
  *
from
  digitalocean_project
where
  is_default;
```

