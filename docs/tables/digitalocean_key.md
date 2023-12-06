---
title: "Steampipe Table: digitalocean_key - Query DigitalOcean SSH Keys using SQL"
description: "Allows users to query DigitalOcean SSH Keys, providing insights into the SSH keys used to access Droplets."
---

# Table: digitalocean_key - Query DigitalOcean SSH Keys using SQL

DigitalOcean SSH Keys are a secure way of logging into a virtual private server with SSH than using a password. With SSH keys, users can log into a server without a password. They provide a more secure way of logging into a server with SSH than using a password alone.

## Table Usage Guide

The `digitalocean_key` table provides insights into SSH keys within DigitalOcean. As a DevOps engineer, explore key-specific details through this table, including key fingerprint, public key, and associated metadata. Utilize it to uncover information about keys, such as their names, ids, and the Droplets they are associated with.

## Examples

### List all keys
Explore all the keys available in your DigitalOcean account to better manage your resources and security settings. This can assist in identifying unused or unnecessary keys, thereby helping to streamline operations and enhance security.

```sql+postgres
select
  *
from
  digitalocean_key;
```

```sql+sqlite
select
  *
from
  digitalocean_key;
```

### Get a Key by ID
Analyze the settings to understand the specific details of a digital key using its unique identifier. This can be useful in managing and validating your digital keys for secure operations.

```sql+postgres
select
  id,
  name,
  fingerprint
from
  digitalocean_key
where
  id = 29280599;
```

```sql+sqlite
select
  id,
  name,
  fingerprint
from
  digitalocean_key
where
  id = 29280599;
```

### Get a Key by Fingerprint
Explore the specific digital key associated with a given fingerprint, which is useful for identifying and verifying the ownership of resources in a DigitalOcean environment. This could be particularly beneficial in scenarios where you need to audit resource access or troubleshoot permission issues.

```sql+postgres
select
  id,
  name,
  fingerprint
from
  digitalocean_key
where
  fingerprint = '3a:84:c2:cc:77:e9:ea:95:5b:45:c3:4d:92:fc:4a:ac';
```

```sql+sqlite
select
  id,
  name,
  fingerprint
from
  digitalocean_key
where
  fingerprint = '3a:84:c2:cc:77:e9:ea:95:5b:45:c3:4d:92:fc:4a:ac';
```