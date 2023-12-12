---
title: "Steampipe Table: digitalocean_alert_policy - Query DigitalOcean Alert Policies using SQL"
description: "Allows users to query Alert Policies in DigitalOcean, specifically the alert policy configurations and associated details, providing insights into alert management and monitoring."
---

# Table: digitalocean_alert_policy - Query DigitalOcean Alert Policies using SQL

DigitalOcean Alert Policies is a feature within DigitalOcean Monitoring that enables users to set up alerts based on specific conditions related to their infrastructure. These alerts can notify users about any operational issues, performance anomalies, or resource constraints in their DigitalOcean resources. The feature allows users to proactively manage and maintain the health and performance of their applications and infrastructure.

## Table Usage Guide

The `digitalocean_alert_policy` table offers detailed information about the alert policies configured within DigitalOcean Monitoring. As a systems administrator or DevOps engineer, you can leverage this table to understand the alert configurations, associated thresholds, and target resources. This can assist in proactive infrastructure monitoring, anomaly detection, and incident management.

## Examples

### Basic info
Discover the segments that are active within your DigitalOcean alert policies. This can help you understand which types of alerts are currently enabled, allowing you to assess your alert system's functionality and coverage.

```sql+postgres
select
  uuid,
  enabled,
  type
from
  digitalocean_alert_policy;
```

```sql+sqlite
select
  uuid,
  enabled,
  type
from
  digitalocean_alert_policy;
```

### List disabled alert policies
Analyze the settings to understand which alert policies have been disabled on your DigitalOcean account. This can be beneficial in identifying potential gaps in your monitoring and alerting strategy.

```sql+postgres
select
  uuid,
  enabled,
  type
from
  digitalocean_alert_policy
where
  not enabled;
```

```sql+sqlite
select
  uuid,
  enabled,
  type
from
  digitalocean_alert_policy
where
  enabled = 0;
```