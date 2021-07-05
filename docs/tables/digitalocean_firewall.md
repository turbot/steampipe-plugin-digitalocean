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
