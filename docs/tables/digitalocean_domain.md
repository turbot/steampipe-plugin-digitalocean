---
title: "Steampipe Table: digitalocean_domain - Query DigitalOcean Domains using SQL"
description: "Allows users to query DigitalOcean Domains, providing insights into the DNS records and associated metadata."
---

# Table: digitalocean_domain - Query DigitalOcean Domains using SQL

A DigitalOcean Domain is a resource in the DigitalOcean cloud platform that represents a DNS zone file. It provides a way to manage DNS records for a particular domain, enabling users to point domain names to various servers and services. DigitalOcean Domains help in managing and configuring domain names for applications hosted on DigitalOcean.

## Table Usage Guide

The `digitalocean_domain` table provides insights into the DNS records within the DigitalOcean platform. As a system administrator or a DevOps engineer, explore domain-specific details through this table, including names, TTL values, and associated metadata. Utilize it to uncover information about domains, such as their zone file details, the IP addresses they are pointing to, and the verification of DNS configurations.

## Examples

### Basic info
Explore the various domains on DigitalOcean and understand their unique identifiers and time to live (TTL) settings. This can be helpful for managing and optimizing your domain configurations.

```sql+postgres
select
  name,
  urn,
  ttl
from
  digitalocean_domain;
```

```sql+sqlite
select
  name,
  urn,
  ttl
from
  digitalocean_domain;
```