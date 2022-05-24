---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/digitalocean.svg"
brand_color: "#008bcf"
display_name: "DigitalOcean"
name: "digitalocean"
description: "Steampipe plugin for querying DigitalOcean databases, networks, and other resources."
og_description: "Query DigitalOcean with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/do-social-graphic.png"
---

# DigitalOcean + Steampipe

Query your DigitalOcean infrastructure including droplets, databases, networks, and more.

[DigitalOcean](https://www.digitalocean.com/) provides scalable and on-demand cloud infrastructure solutions for hosting or storage needs.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example:

```sql
select
  slug,
  name,
  available
from
  digitalocean_region
 ```
 
 ```
+------+-------------+-----------+
| slug | name        | available |
+------+-------------+-----------+
| nyc1 | New York 1  | true      |
| nyc3 | New York 3  | true      |
| ams2 | Amsterdam 2 | false     |
| sgp1 | Singapore 1 | true      |
| nyc2 | New York 2  | false     |
+------+-------------+-----------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/digitalocean/tables)**

## Get started

### Install

Download and install the latest DigitalOcean plugin:

```bash
steampipe plugin install digitalocean
```

### Configuration

Installing the latest DigitalOcean plugin will create a config file (`~/.steampipe/config/digitalocean.spc`) with a single connection named `digitalocean`:

```hcl
connection "digitalocean" {
  plugin  = "digitalocean"

  # Personal Access Token for your DigitalOcean account
  # Reference: https://www.digitalocean.com/docs/apis-clis/api/create-personal-access-token
  # Env variables (in order of precedence): DIGITALOCEAN_TOKEN, DIGITALOCEAN_ACCESS_TOKEN
  # token = "YOUR_DIGITALOCEAN_ACCESS_TOKEN"
}
```

### Example Configurations

- Connect to a single account:

  ```hcl
  connection "digitalocean_my_account" {
    plugin = "digitalocean"
    token  = "1646968370949-df954218b5da5b8614c85cc454136b27"
  }
  ```

- Create connections to multiple accounts:

  ```hcl
  connection "account_aaa" {
    plugin    = "digitalocean"
    token     = "1646968370949-df954218b5da5b8614c85cc4541abcde"
  }

  connection "account_bbb" {
    plugin    = "digitalocean"
    token     = "1646968370949-df954218b5da5b8614c85cc4541fghij"
  }

  connection "account_ccc" {
    plugin    = "digitalocean"
    token     = "1646968370949-df954218b5da5b8614c85cc4541klmno"
  }
  ```

## Get Involved

* Open source: https://github.com/turbot/steampipe-plugin-digitalocean
* Community: [Slack Channel](https://steampipe.io/community/join)
