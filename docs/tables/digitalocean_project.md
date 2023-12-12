---
title: "Steampipe Table: digitalocean_project - Query DigitalOcean Projects using SQL"
description: "Allows users to query DigitalOcean Projects, specifically providing insights into project details, including name, description, purpose, environment, and associated resources."
---

# Table: digitalocean_project - Query DigitalOcean Projects using SQL

A DigitalOcean Project is a high-level organizational unit within the DigitalOcean platform that allows users to group and manage their resources (like Droplets, Spaces, and load balancers) based on their respective workflows. Projects help users to manage their infrastructure more efficiently by providing a way to view and organize their resources in a way that aligns with their application architecture or business requirements. It provides a way to manage resources in a structured manner, making it easier to find and manage resources.

## Table Usage Guide

The `digitalocean_project` table provides insights into projects within DigitalOcean. As a DevOps engineer, explore project-specific details through this table, including names, descriptions, purposes, environments, and associated resources. Utilize it to uncover information about projects, such as their purpose, the environment they are associated with, and the resources they contain.

## Examples

### List all projects
Explore all your projects on DigitalOcean to gain an overview of your current work. This can help you manage resources efficiently and plan future projects effectively.

```sql+postgres
select
  *
from
  digitalocean_project;
```

```sql+sqlite
select
  *
from
  digitalocean_project;
```

### Get a project by ID
This example demonstrates how to locate a specific project within DigitalOcean's resources. It's particularly useful when you need to quickly access or review the details of a specific project, identified by its unique ID.

```sql+postgres
select
  *
from
  digitalocean_project
where
  id = '59137997-528e-45ef-9521-db041f2c7d94';
```

```sql+sqlite
select
  *
from
  digitalocean_project
where
  id = '59137997-528e-45ef-9521-db041f2c7d94';
```

### Get the default project
Explore the default project settings in your DigitalOcean account to understand its configuration and operational parameters. This can help you assess the current setup and make necessary adjustments for optimal performance.

```sql+postgres
select
  *
from
  digitalocean_project
where
  is_default;
```

```sql+sqlite
select
  *
from
  digitalocean_project
where
  is_default = 1;
```