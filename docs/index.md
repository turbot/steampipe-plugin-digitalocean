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

## Configuration Arguments

To use the API, you'll first generate a personal access token. Personal access tokens function like ordinary OAuth access tokens. You can use them to authenticate to the API by including one in a bearer-type Authorization header with your request.

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

## Configure API Token (Example configurations)

The default connection. This uses standard Application Default Credentials (ADC) against the active account as configured for digitalocean

connection "digitalocean" {
plugin    = "digitalocean"
}
A connection to a specific account, using standard ADC Credentials.

connection "digitalocean_my_account" {
plugin    = "digitalocean"
token   = "xoxp-2556146250-EXAMPLE-1646968370949-df954218b5da5b8614c85cc454136b27"
}

