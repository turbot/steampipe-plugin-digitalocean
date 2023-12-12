---
title: "Steampipe Table: digitalocean_vpc - Query DigitalOcean Virtual Private Clouds using SQL"
description: "Allows users to query DigitalOcean Virtual Private Clouds (VPCs), specifically providing insights into their configurations, associated resources, and other metadata."
---

# Table: digitalocean_vpc - Query DigitalOcean Virtual Private Clouds using SQL

A DigitalOcean Virtual Private Cloud (VPC) is a private network interface for collections of DigitalOcean resources. VPCs can help to manage the network's resources where these resources are kept and provide a layer of isolation and security. They allow users to control their network settings and handle the traffic to their resources.

## Table Usage Guide

The `digitalocean_vpc` table provides insights into the configuration and metadata of VPCs within DigitalOcean. As a network administrator, explore details about each VPC, including its IP range, region, and associated resources. Utilize it to uncover information about the network settings, manage the traffic to your resources, and maintain the isolation and security of your resources.

## Examples

### List all vpcs
Explore all the Virtual Private Clouds (VPCs) available in your DigitalOcean account to understand the network resources and services associated with each. This will aid in managing and optimizing your cloud infrastructure.

```sql+postgres
select
  *
from
  digitalocean_vpc;
```

```sql+sqlite
select
  *
from
  digitalocean_vpc;
```

### Get a VPC by ID
Discover the details of a specific Virtual Private Cloud (VPC) using its unique identifier. This can be useful when you need to understand the specific settings and configurations of a particular VPC within your DigitalOcean environment.

```sql+postgres
select
  id,
  name,
  ip_range
from
  digitalocean_vpc
where
  id = '411728c1-f29d-474c-bd9c-8c4f261c7904';
```

```sql+sqlite
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
Identify the specific Virtual Private Cloud (VPC) that includes a certain IP address. This is useful for pinpointing the location of a device or user within your network.

```sql+postgres
select
  id,
  name,
  ip_range
from
  digitalocean_vpc
where
  ip_range >> '10.106.0.123';
```

```sql+sqlite
Error: SQLite does not support CIDR operations.
```

### Default VPCs
Explore which Virtual Private Clouds (VPCs) are set as default in your DigitalOcean account. This can help in managing network configurations and understand potential security implications.

```sql+postgres
select
  id,
  name,
  ip_range
from
  digitalocean_vpc
where
  is_default;
```

```sql+sqlite
select
  id,
  name,
  ip_range
from
  digitalocean_vpc
where
  is_default;
```