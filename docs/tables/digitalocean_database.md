# Table: digitalocean_database

DigitalOcean's managed database service simplifies the creation and management
of highly available database clusters. Currently, it offers support for
PostgreSQL, Redis, and MySQL.

## Examples

### List all databases

```sql
select
  *
from
  digitalocean_database;
```

### Get database by ID

```sql
select
  *
from
  digitalocean_database
where
  id = 'fad76135-48bb-49c8-a274-a9db584e1dc3';
```

### All database users by instance

```sql
select
  db.name as db_name,
  u ->> 'name' as user_name,
  u ->> 'role' as user_role
from
  digitalocean_database as db,
  jsonb_array_elements(users) as u
```

### Databases not using SSL

```sql
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

WARNING: DigitalOcean returns the database password as metadata. Use with care!

```sql
select
  name,
  connection_uri,
  private_connection_uri
from
  digitalocean_database;
```

### Databases by engine version

```sql
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

### Get database firewall trusted sources

```sql
select 
  name as "Name",
  firewall ->> 'type' as "Firewall Source",
  firewall ->> 'value' as "Source ID" 
from 
  digitalocean_database, 
  jsonb_array_elements(firewall_rules) as firewall;
```