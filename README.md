![image](https://hub.steampipe.io/images/plugins/turbot/do-social-graphic.png)

# Digital Ocean Plugin for Steampipe
<p align="center">
  <a aria-label="Steampipe logo" href="https://steampipe.io">
    <img src="https://steampipe.io/images/steampipe_logo_wordmark_padding.svg" height="28">
  </a>
  <a aria-label="License" href="LICENSE">
    <img alt="" src="https://img.shields.io/static/v1?label=license&message=Apache-2.0&style=for-the-badge&labelColor=777777&color=F3F1F0">
  </a>
</p>

Use SQL to query infrastructure including servers, networks, identity and more from Digital Ocean.

- **[Get started →](https://hub.steampipe.io/plugins/turbot/digitalocean)**
- Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/digitalocean/tables)
- Community: [Slack Channel](https://join.slack.com/t/steampipe/shared_invite/zt-oij778tv-lYyRTWOTMQYBVAbtPSWs3g)
- Get involved: [Issues](https://github.com/turbot/steampipe-plugin-digitalocean/issues)

## Quick Start
Install the plugin with [Steampipe](https://steampipe.io):
```shell
steampipe plugin install digitalocean
```

Run a Query:
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

## Documentation
- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Community and Contribution
### Community
The Steampipe community can be found on [GitHub Discussions](https://github.com/turbot/steampipe/discussions), where you can ask questions, voice ideas, and share your projects.

Our [Code of Conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md) applies to all Steampipe community channels.

### Contributing
Please see [CONTRIBUTING.md](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md).

