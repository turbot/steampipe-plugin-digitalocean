---
title: "Steampipe Table: digitalocean_app - Query DigitalOcean Apps using SQL"
description: "Allows users to query Apps in DigitalOcean, specifically to gather information about the app's ID, name, live URL, active deployment and region. This data can be used to monitor and manage the apps deployed on DigitalOcean."
---

# Table: digitalocean_app - Query DigitalOcean Apps using SQL

DigitalOcean App Platform is a Platform as a Service (PaaS) offering that allows developers to publish code directly to DigitalOcean servers without worrying about the underlying infrastructure. The service automatically analyzes the code, creates containers, and runs them on Kubernetes clusters. It supports several programming languages and frameworks and offers built-in database and caching for high scalability.

## Table Usage Guide

The `digitalocean_app` table provides insights into Apps deployed on the DigitalOcean App Platform. As a developer or DevOps engineer, explore app-specific details through this table, including the app's ID, name, live URL, active deployment and region. Utilize it to monitor and manage the apps deployed on DigitalOcean, and to ensure the smooth operation and performance of your applications.

## Examples

### Basic info
Explore the basic information about your DigitalOcean applications, such as their unique identifiers and creation dates, to gain a better understanding of your resources. This can be particularly useful for auditing purposes or when planning resource management strategies.

```sql+postgres
select
  id,
  name,
  urn,
  created_at
from
  digitalocean_app;
```

```sql+sqlite
select
  id,
  name,
  urn,
  created_at
from
  digitalocean_app;
```