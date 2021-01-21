# Table: digitalocean_key

DigitalOcean allows you to add SSH public keys to the interface so that you can
embed your public key into a Droplet at the time of creation. Only the public
key is required to take advantage of this functionality.

## Examples

### List all keys

```sql
select
  *
from
  digitalocean_key;
```

### Get a Key by ID

```sql
select
  id,
  name,
  fingerprint
from
  digitalocean_key
where
  id = 29280599;
```

### Get a Key by Fingerprint

```sql
select
  id,
  name,
  fingerprint
from
  digitalocean_key
where
  fingerprint = '3a:84:c2:cc:77:e9:ea:95:5b:45:c3:4d:92:fc:4a:ac'
```
