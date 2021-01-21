# Table: digitalocean_load_balancer

DigitalOcean Load Balancers provide a way to distribute traffic across multiple
Droplets.

## Examples

### List all load balancers

```sql
select
  *
from
  digitalocean_load_balancer;
```

### Get load balancer by ID

```sql
select
  *
from
  digitalocean_load_balancer
where
  id = 'fad76135-48bb-49c8-a274-a9db584e1dc3';
```

### List load balancers and their rules

```sql
select
  lb.name,
  rule.*
from
  digitalocean_load_balancer as lb,
  jsonb_array_elements(forwarding_rules) as rule
```

### Ensure HTTPS is used as the health check protocol

```sql
select
  name,
  health_check_protocol
from
  digitalocean_load_balancer
where
  health_check_protocol != 'https';
```

### Get Load Balancer and VPC information

```sql
select
  lb.name,
  vpc.name,
  vpc.ip_range
from
  digitalocean_load_balancer as lb,
  digitalocean_vpc as vpc
where
  lb.vpc_uuid = vpc.id;
```

