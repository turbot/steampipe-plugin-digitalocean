![image](https://hub.steampipe.io/images/plugins/turbot/do-social-graphic.png)

# DigitalOcean Plugin for Steampipe

Use SQL to query infrastructure including servers, networks, identity and more from DigitalOcean.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/digitalocean)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/digitalocean/tables)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-digitalocean/issues)

## Quick Start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install digitalocean
```

Run a query:

```sql
select
  name,
  engine,
  version
from
  digitalocean_database;
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-digitalocean.git
cd steampipe-plugin-digitalocean
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/digitalocean.spc
```

Try it!

```
steampipe query
> .inspect digitalocean
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). Contributions to the plugin are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-digitalocean/blob/main/LICENSE). Contributions to the plugin documentation are subject to the [CC BY-NC-ND license](https://github.com/turbot/steampipe-plugin-digitalocean/blob/main/docs/LICENSE).

`help wanted` issues:

- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [DigitalOcean Plugin](https://github.com/turbot/steampipe-plugin-digitalocean/labels/help%20wanted)
