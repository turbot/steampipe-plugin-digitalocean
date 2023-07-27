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

## Multi-Account Connections

You may create multiple digitalocean connections:

```hcl
connection "do_dev" {
  plugin    = "digitalocean"
  token     = "1646968370949-df954218b5da5b8614c85cc4541abcde"
}

connection "do_qa" {
  plugin    = "digitalocean"
  token     = "1646968370949-df954218b5da5b8614c85cc4541fghij"
}

connection "do_prod" {
  plugin    = "digitalocean"
  token     = "1646968370949-df954218b5da5b8614c85cc4541klmno"
}
```

Each connection is implemented as a distinct [Postgres schema](https://www.postgresql.org/docs/current/ddl-schemas.html). As such, you can use qualified table names to query a specific connection:

```sql
select * from do_qa.digitalocean_project
```

You can create multi-account connections by using an [**aggregator** connection](https://steampipe.io/docs/using-steampipe/managing-connections#using-aggregators). Aggregators allow you to query data from multiple connections for a plugin as if they are a single connection.

```hcl
connection "do_all" {
  plugin      = "digitalocean"
  type        = "aggregator"
  connections = ["do_dev", "do_qa", "do_prod"]
}
```

Querying tables from this connection will return results from the `do_dev`, `do_qa`, and `do_prod` connections:

```sql
select * from do_all.digitalocean_project
```

Alternatively, you can use an unqualified name and it will be resolved according to the [Search Path](https://steampipe.io/docs/guides/search-path). It's a good idea to name your aggregator first alphabetically so that it is the first connection in the search path (i.e. `do_all` comes before `do_dev`):

```sql
select * from digitalocean_project
```

Steampipe supports the `*` wildcard in the connection names. For example, to aggregate all the digitalocean plugin connections whose names begin with `do_`:

```hcl
connection "do_all" {
  type        = "aggregator"
  plugin      = "digitalocean"
  connections = ["do_*"]
}
```

## Get Involved

* Open source: https://github.com/turbot/steampipe-plugin-digitalocean
* Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
