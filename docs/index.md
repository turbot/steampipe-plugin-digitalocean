---
organization: Turbot
category: ["public cloud"]
icon_url: "/images/plugins/turbot/digitalocean.svg"
brand_color: "#008bcf"
display_name: "DigitalOcean"
short_name: "digitalocean"
description: "Steampipe plugin for querying DigitalOcean databases, networks, and other resources."
---

# DigitalOcean + Steampipe

Query your DigitalOcean infrastructure including droplets, databases, networks and more.

[DigitalOcean](https://www.digitalocean.com/) provides scalable and on-demand cloud infrastructure solutions for hosting or storage needs.

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

For example: 

Need example query and table 

## Documentation

- **[Table definitions & examples →](https://hub.steampipe.io/plugins/turbot/digitalocean/tables)**

## Get started

### Install

Download and install the latest DigitalOcean plugin:

```bash
$ steampipe plugin install digitalocean
Installing plugin digitalocean...
$
```

### Credentials

Needs Content

### Scope

A DigitalOcean connection is scoped to a single DigitalOcean account, with a single set of credentials.

### Connection Configuration

Connection configuration is defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

Installing the latest DigitalOcean plugin will create a connection file (`~/.steampipe/config/digitalocean.spc`) with a single connection named `digitalocean`. You must modify this connection to include your Personal Access Token for DigitalOcean account.

```hcl
connection "digitalocean" {
  plugin  = "digitalocean"
  token   = "17ImlCYdfZ3WJIrGk96gCpJn1fi1pLwVdrb23kj4"
}
```

### Configuration Arguments

The DigitalOcean plugin allows you set static credentials with the `token` argument. Personal access tokens function like ordinary OAuth access tokens -- You can use them to authenticate to the API by including it in a bearer-type authorization header along with your request. 

To use the plugin, you will need to [create personal access token](https://www.digitalocean.com/docs/apis-clis/api/create-personal-access-token/).  Read scope is required (write is not).

If the `token` argument is not specified for a connection, the project will be determined as per the following order:
  - DIGITALOCEAN_TOKEN environment variable, if set; otherwise
  - DIGITALOCEAN_ACCESS_TOKEN environment variable, if set (this is deprecated).

#### Configuration Examples

- Using token to establish connection to a specific account.

  ```hcl
  connection "digitalocean_my_account" {
    plugin = "digitalocean"
    token  = "1646968370949-df954218b5da5b8614c85cc454136b27"
  }
  ```

- Common configuration having multiple connections to different DigitalOcean accounts.

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

## Configuring DigitalOCean Credentials

Needs content

## Get Involved

* Open source: https://github.com/turbot/steampipe-plugin-digitalocean
* Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
