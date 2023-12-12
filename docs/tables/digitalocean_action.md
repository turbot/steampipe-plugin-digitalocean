---
title: "Steampipe Table: digitalocean_action - Query DigitalOcean Actions using SQL"
description: "Allows users to query DigitalOcean Actions, specifically the status, type, and associated resources, providing insights into the operations performed on the DigitalOcean platform."
---

# Table: digitalocean_action - Query DigitalOcean Actions using SQL

DigitalOcean Actions are operations that you can perform on the various resources within the DigitalOcean platform. Actions include operations such as creating a Droplet, resizing a Droplet, and transferring a Droplet. Each action is associated with specific resources and has a distinct status indicating the progress and result of the operation.

## Table Usage Guide

The `digitalocean_action` table provides insights into the operations performed on the DigitalOcean platform. As a system administrator or DevOps engineer, explore action-specific details through this table, including the type of action, the status of the operation, and the resources associated with each action. Utilize it to monitor and track the operations performed on your DigitalOcean resources, ensuring smooth and efficient management of your digital infrastructure.

## Examples

### List Actions
Determine the areas in which actions have been taken within the DigitalOcean platform. This can help in tracking and understanding the various activities that have occurred, offering insights for future planning and decision-making.

```sql+postgres
select
  *
from
  digitalocean_action;
```

```sql+sqlite
select
  *
from
  digitalocean_action;
```

### Get an action by ID
Explore which actions have been performed in your DigitalOcean environment by searching for a specific action ID. This is useful for auditing and tracking changes, ensuring you have a clear understanding of all operations and their impacts.

```sql+postgres
select
  *
from
  digitalocean_action
where
  id = 1101196049;
```

```sql+sqlite
select
  *
from
  digitalocean_action
where
  id = 1101196049;
```

### Droplets created today
Explore which new digital resources, referred to as 'droplets', have been created in the last 24 hours. This is useful for keeping track of recent additions and changes in your digital environment.

```sql+postgres
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

```sql+sqlite
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
  and a.started_at > datetime('now', '-24 hours');
```

### Resources deleted in the last 3 days
Explore which resources have been deleted in the past three days. This can help in tracking and managing resource usage and changes.

```sql+postgres
select
  *
from
  digitalocean_action
where
  type = 'destroy'
  and a.started_at > now() - interval '3 days';
```

```sql+sqlite
select
  *
from
  digitalocean_action
where
  type = 'destroy'
  and a.started_at > datetime('now', '-3 days');
```

### Creations by resource type in the last week
Identify instances where resources were created in the past week on DigitalOcean. This helps in understanding the distribution and frequency of different resource types being utilized, assisting in resource management and planning.

```sql+postgres
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

```sql+sqlite
select
  resource_type,
  count(*)
from
  digitalocean_action
where
  type = 'create'
  and started_at > datetime('now', '-7 days')
group by
  resource_type
order by
  count(*) desc;
```