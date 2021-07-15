# Table: digitalocean_domain

Adding a domain you own to your DigitalOcean account lets you manage the domain’s DNS records with the control panel and API. Domains you manage on DigitalOcean also integrate with DigitalOcean Load Balancers and Spaces to streamline automatic SSL certificate management.

## Examples

### Basic info

```sql
select
  name,
  urn,
  ttl
from
  digitalocean_domain;
```
