---
title: "Steampipe Table: digitalocean_database - Query DigitalOcean Databases using SQL"
description: "Allows users to query DigitalOcean Databases, providing details about the databases available in the DigitalOcean resources. The table provides insights into the database configurations, regions, users, and other associated metadata."
---

# Table: digitalocean_database - Query DigitalOcean Databases using SQL

DigitalOcean Databases is a service offered by DigitalOcean that provides fully managed databases. It provides a scalable, reliable, and secure environment for developers to manage and scale their databases. The service supports multiple database engines, including PostgreSQL, MySQL, and Redis.

## Table Usage Guide

The `digitalocean_database` table provides insights into the databases in DigitalOcean. As a database administrator or developer, you can explore specific details about each database, including its configurations, associated users, regions, and more. Utilize this table to manage and monitor your databases effectively, ensuring optimal performance and security.

## Examples

### List all databases
Explore the full range of databases within your DigitalOcean environment. This can aid in assessing overall usage and identifying any databases that may need attention or adjustment for optimal performance.

```sql+postgres
select
  *
from
  digitalocean_database;
```

```sql+sqlite
select
  *
from
  digitalocean_database;
```

### Get database by ID
Explore the specific details of a particular DigitalOcean database by identifying it through its unique ID. This can be useful for understanding the properties and configurations of a specific database within your DigitalOcean environment.

```sql+postgres
select
  *
from
  digitalocean_database
where
  id = 'fad76135-48bb-49c8-a274-a9db584e1dc3';
```

```sql+sqlite
select
  *
from
  digitalocean_database
where
  id = 'fad76135-48bb-49c8-a274-a9db584e1dc3';
```

### All database users by instance
Explore which users are associated with each instance in your DigitalOcean database. This is useful for understanding user roles and responsibilities within your database management system.

```sql+postgres
select
  db.name as db_name,
  u ->> 'name' as user_name,
  u ->> 'role' as user_role
from
  digitalocean_database as db,
  jsonb_array_elements(users) as u;
```

```sql+sqlite
select
  db.name as db_name,
  json_extract(u.value, '$.name') as user_name,
  json_extract(u.value, '$.role') as user_role
from
  digitalocean_database as db,
  json_each(db.users) as u;
```

### Databases not using SSL
Explore which DigitalOcean databases are not using SSL, helping to identify potential security vulnerabilities in your database connections. This can be beneficial in enhancing your data security measures.

```sql+postgres
select
  name,
  connection_ssl
from
  digitalocean_database
where
  not connection_ssl
  or not private_connection_ssl;
```

```sql+sqlite
select
  name,
  connection_ssl
from
  digitalocean_database
where
  not connection_ssl
  or not private_connection_ssl;
```

### Get database connection
Explore which DigitalOcean databases you are connected to, by obtaining the connection URIs. This can be beneficial in managing your connections and ensuring the security of your private connections.
WARNING: DigitalOcean returns the database password as metadata. Use with care!


```sql+postgres
select
  name,
  connection_uri,
  private_connection_uri
from
  digitalocean_database;
```

```sql+sqlite
select
  name,
  connection_uri,
  private_connection_uri
from
  digitalocean_database;
```

### Databases by engine version
Explore the distribution of your DigitalOcean databases by identifying the number of databases operating on different engine versions. This can help you manage and plan upgrades, ensuring your systems remain up-to-date and secure.

```sql+postgres
select
  engine,
  version,
  count(id)
from
  digitalocean_database
group by
  engine,
  version
order by
  count desc;
```

```sql+sqlite
select
  engine,
  version,
  count(id)
from
  digitalocean_database
group by
  engine,
  version
order by
  count(id) desc;
```

### Get database firewall trusted sources
Explore which trusted sources are allowed through your database firewall. This is useful for reviewing your security settings and ensuring only authorized sources have access.

```sql+postgres
select 
  name as "Name",
  firewall ->> 'type' as "Firewall Source",
  firewall ->> 'value' as "Source ID" 
from 
  digitalocean_database, 
  jsonb_array_elements(firewall_rules) as firewall;
```

```sql+sqlite
select 
  name as "Name",
  json_extract(firewall.value, '$.type') as "Firewall Source",
  json_extract(firewall.value, '$.value') as "Source ID" 
from 
  digitalocean_database, 
  json_each(firewall_rules) as firewall;
```