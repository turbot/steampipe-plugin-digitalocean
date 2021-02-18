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

### Scope

A DigitalOcean connection is scoped to a single DigitalOcean account, with a single set of credentials.

## Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

Installing the latest digitalocean plugin will create a connection file (`~/.steampipe/config/digitalocean.spc`) with a single connection named `digitalocean`. You must modify this connection to include your Personal Access Token for Digital Ocean account.

```hcl
connection "digitalocean" {
  plugin  = "digitalocean"
  token   = "17ImlCYdfZ3WJIrGk96gCpJn1fi1pLwVdrb23kj4"
}
```

### Configuration Arguments

The DigitalOcean plugin allows you set static credentials with the `token` argument. Personal access tokens function like ordinary OAuth access tokens -- You can use them to authenticate to the API by including it in a bearer-type Authorization header with your request. 

To use the plugin, you'll first need to [create personal access token](https://www.digitalocean.com/docs/apis-clis/api/create-personal-access-token/).  Read scope is required (write is not).

If the `token` argument is not specified for a connection, the project will be determined in the following order:
  - The DIGITALOCEAN_TOKEN environment variable, if set; otherwise
  - The DIGITALOCEAN_ACCESS_TOKEN environment variable, if set (this is deprecated).

#### Example configurations

- A connection to a specific account, using token.

  ```hcl
  connection "digitalocean_my_account" {
    plugin = "digitalocean"
    token  = "1646968370949-df954218b5da5b8614c85cc454136b27"
  }
  ```

- A common configuration is to have multiple connections to different DigitalOcean accounts.
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
