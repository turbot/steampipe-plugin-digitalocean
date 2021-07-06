# Table: digitalocean_firewall

DigitalOcean Cloud Firewalls are a network-based, stateful firewall service for Droplets provided at no additional cost. Cloud firewalls block all traffic that isn’t expressly permitted by a rule.

## Examples

### Basic info

```sql
select
  id,
  name,
  created_at,
  status
from
  digitalocean_firewall;
```

## List firewalls whose inbound access is not restricted

```sql
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

### List failed firewalls

```sql
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
