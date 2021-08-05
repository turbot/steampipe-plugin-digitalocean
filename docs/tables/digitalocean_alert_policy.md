# Table: digitalocean_alert_policy

DigitalOcean Monitoring is a free, opt-in service that gathers metrics about Droplet-level resource utilization. It provides additional Droplet graphs and supports configurable metrics alert policies with integrated email Slack notifications to help you track the operational health of your infrastructure.

## Examples

### Basic info

```sql
select
  uuid,
  enabled,
  type
from
  digitalocean_alert_policy;
```

### List disabled alert policies

```sql
select
  uuid,
  enabled,
  type
from
  digitalocean_alert_policy
where
  not enabled;
```
