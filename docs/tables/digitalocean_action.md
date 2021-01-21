# Table: digitalocean_action

Actions are records of events that have occurred on the resources in your DigitalOcean account. These can be things like rebooting a Action, or transferring an image to a new region.

## Examples

### List Actions

```sql
select
  *
from
  digitalocean_action;
```

### Get an action by ID

```sql
select
  *
from
  digitalocean_action
where
  id = 1101196049;
```

### Droplets created today

```sql
select
  d.name,
  d.private_ipv4,
  a.resource_type,
  a.type,
  a.id as action_id
from
  digitalocean_action as a,
  digitalocean_droplet as d
where
  a.resource_id = d.id
  and a.type = 'create'
  and a.resource_type = 'droplet'
  and a.started_at > now() - interval '24 hours';
```

### Resources deleted in the last 3 days

```sql
select
  *
from
  digitalocean_action
where
  type = 'destroy'
  and a.started_at > now() - interval '3 days';
```

### Creations by resource type in the last week

```sql
select
  resource_type,
  count(*)
from
  digitalocean_action
where
  type = 'create'
  and started_at > now() - interval '1 week'
group by
  resource_type
order by
  count;
```
