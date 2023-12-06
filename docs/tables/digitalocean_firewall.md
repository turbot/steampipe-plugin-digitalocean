---
title: "Steampipe Table: digitalocean_firewall - Query DigitalOcean Firewalls using SQL"
description: "Allows users to query DigitalOcean Firewalls, providing insights into firewall configurations and rules."
---

# Table: digitalocean_firewall - Query DigitalOcean Firewalls using SQL

DigitalOcean Firewalls are a security feature that controls the traffic to your Droplet. Firewalls place a barrier between your servers and other machines on the network to protect them from external attacks. Firewalls can be customized to only allow traffic to certain ports and addresses.

## Table Usage Guide

The `digitalocean_firewall` table provides insights into firewall configurations within DigitalOcean. As a DevOps engineer, explore firewall-specific details through this table, including inbound and outbound rules, associated Droplets, and tags. Utilize it to uncover information about firewall rules, the Droplets they apply to, and the overall security of your network.

## Examples

### Basic info
Explore which firewalls have unrestricted inbound access, potentially posing a security risk. This is useful for identifying and mitigating potential vulnerabilities in your network's security.

```sql+postgres
select
  id,
  name,
  created_at,
  status
from
  digitalocean_firewall;
```

```sql+sqlite
select
  id,
  name,
  created_at,
  status
from
  digitalocean_firewall;
```

## List firewalls whose inbound access is not restricted

```sql+postgres
select
  id,
  name,
  created_at,
  status
from
  digitalocean_firewall,
  jsonb_array_elements(inbound_rules) as i
where
  i -> 'sources' -> 'addresses' = '["0.0.0.0/0","::/0"]';
```

```sql+sqlite
select
  digitalocean_firewall.id,
  digitalocean_firewall.name,
  digitalocean_firewall.created_at,
  digitalocean_firewall.status
from
  digitalocean_firewall,
  json_each(inbound_rules) as i
where
  json_extract(i.value, '$.sources.addresses') = '["0.0.0.0/0","::/0"]';
```

### List failed firewalls
Identify instances where firewall creation attempts have been unsuccessful. This could be useful in troubleshooting and ensuring the security of your digital assets.

```sql+postgres
select
  id,
  name,
  created_at,
  status
from
  digitalocean_firewall
where
  status = 'failed';
```

```sql+sqlite
select
  id,
  name,
  created_at,
  status
from
  digitalocean_firewall
where
  status = 'failed';
```