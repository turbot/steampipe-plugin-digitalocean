---
title: "Steampipe Table: digitalocean_load_balancer - Query DigitalOcean Load Balancers using SQL"
description: "Allows users to query DigitalOcean Load Balancers, specifically providing details about each load balancer's configuration, status, and associated resources."
---

# Table: digitalocean_load_balancer - Query DigitalOcean Load Balancers using SQL

DigitalOcean Load Balancers are a fully-managed, highly available network load balancing service. They distribute incoming traffic across your infrastructure to increase your application's availability. Load Balancers are protocol agnostic and route traffic to backend Droplets by using a health check to ensure only healthy Droplets receive traffic.

## Table Usage Guide

The `digitalocean_load_balancer` table provides insights into Load Balancers within DigitalOcean. As a DevOps engineer, you can explore load balancer-specific details through this table, including configuration, status, and associated resources. Utilize it to manage and monitor the health and performance of your load balancers, ensuring optimal distribution of network traffic across your infrastructure.

## Examples

### List all load balancers
Explore all the load balancers in your DigitalOcean account to manage your applications' traffic, ensuring high availability and reliability. This overview can help in assessing the performance of your applications and in identifying any potential bottlenecks.

```sql+postgres
select
  *
from
  digitalocean_load_balancer;
```

```sql+sqlite
select
  *
from
  digitalocean_load_balancer;
```

### Get load balancer by ID
Explore which load balancers are associated with a specific ID to manage network traffic more effectively. This is useful in pinpointing the exact load balancer that needs to be modified or troubleshooted.

```sql+postgres
select
  *
from
  digitalocean_load_balancer
where
  id = 'fad76135-48bb-49c8-a274-a9db584e1dc3';
```

```sql+sqlite
select
  *
from
  digitalocean_load_balancer
where
  id = 'fad76135-48bb-49c8-a274-a9db584e1dc3';
```

### List load balancers and their rules
Discover the segments that associate load balancers with their rules, providing a comprehensive view of how traffic is directed within your DigitalOcean environment. This can help in analyzing the distribution of network loads and optimizing performance.

```sql+postgres
select
  lb.name,
  rule.*
from
  digitalocean_load_balancer as lb,
  jsonb_array_elements(forwarding_rules) as rule
```

```sql+sqlite
select
  lb.name,
  rule.*
from
  digitalocean_load_balancer as lb,
  json_each(lb.forwarding_rules) as rule
```

### Ensure HTTPS is used as the health check protocol
Analyze the settings to understand if your load balancers are using HTTPS for health checks. This is crucial for maintaining secure and encrypted communication in your digitalocean infrastructure.

```sql+postgres
select
  name,
  health_check_protocol
from
  digitalocean_load_balancer
where
  health_check_protocol != 'https';
```

```sql+sqlite
select
  name,
  health_check_protocol
from
  digitalocean_load_balancer
where
  health_check_protocol != 'https';
```

### Get Load Balancer and VPC information
Explore the association between your load balancer and virtual private cloud (VPC) within the DigitalOcean platform. This query helps you understand the network settings of your load balancer, including the VPC it is linked to and the IP range of that VPC.

```sql+postgres
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

```sql+sqlite
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