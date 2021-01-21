---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/digitalocean.svg"
brand_color: "#008bcf"
display_name: "DigitalOcean"
name: "digitalocean"
description: "Steampipe plugin for querying DigitalOcean databases, networks, and other resources."
---

# DigitalOcean

Query your DigitalOcean infrastructure including droplets, databases, networks and more.

## Installation

Download and install the latest DigitalOcean plugin:

```bash
$ steampipe plugin install digitalocean
Installing plugin digitalocean...
$
```

## Configure API Token

[Create a personal access token to use the DigitalOcean plugin](https://www.digitalocean.com/docs/apis-clis/api/create-personal-access-token/).
Read scope is required (write is not).

Set DigitalOcean API token as an environment variable (Mac, Linux):

```bash
export DIGITALOCEAN_TOKEN="xoxp-2556146250-EXAMPLE-1646968370949-df954218b5da5b8614c85cc454136b27"
```

Similar to Terraform, API tokens are loaded from the environment in this order of precedence:

- `DIGITALOCEAN_TOKEN`
- `DIGITALOCEAN_ACCESS_TOKEN`

Steampipe does not yet automatically load `doctl` configuration files.

## Your first query

```bash
$ steampipe query
Welcome to Steampipe v0.0.14
Type ".inspect" for more information.
> select * from digitalocean_account;
+--------------------------+---------------+----------------+-------------------+--------+----------------+--------------------------------------+--------------+
|          email           | droplet_limit | email_verified | floating_ip_limit | status | status_message |                 uuid                 | volume_limit |
+--------------------------+---------------+----------------+-------------------+--------+----------------+--------------------------------------+--------------+
| dwight@dundermifflin.com |            25 | true           |                 3 | active |                | 1593cd23-4203-4bee-b87b-37189e1dcf96 |          100 |
+--------------------------+---------------+----------------+-------------------+--------+----------------+--------------------------------------+--------------+
```
