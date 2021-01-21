# Table: digitalocean_vpc

VPCs (virtual private clouds) are virtual networks containing resources that
can communicate with each other in full isolation using private IP addresses.

## Examples

### List all vpcs

```sql
select
  *
from
  digitalocean_vpc;
```

### Get a VPC by ID

```sql
select
  id,
  name,
  ip_range
from
  digitalocean_vpc
where
  id = '411728c1-f29d-474c-bd9c-8c4f261c7904';
```

### Find the VPC that contains an IP

```sql
select
  id,
  name,
  ip_range
from
  digitalocean_vpc
where
  ip_range >> '10.106.0.123';
```

### Default VPCs

```sql
select
  id,
  name,
  ip_range
from
  digitalocean_vpc
where
  is_default;
```
