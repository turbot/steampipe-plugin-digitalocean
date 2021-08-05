## v0.5.0 [2021-08-05]

_What's new?_

- New tables added
  - [digitalocean_alert_policy](https://hub.steampipe.io/plugins/turbot/digitalocean/tables/digitalocean_alert_policy) ([#33](https://github.com/turbot/steampipe-plugin-digitalocean/pull/33))

_Enhancements_
  - Updated: Recompiled plugin with [steampipe-plugin-sdk v1.4.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v131--2021-07-15) ([#22](https://github.com/turbot/steampipe-plugin-digitalocean/pull/22))

_Bug fixes_
  - Fixed: Example query updated in `digitalocean_volume` table ([#28](https://github.com/turbot/steampipe-plugin-digitalocean/pull/28))
  - Fixed: Querying data for columns `next_backup_window_start` and `next_backup_window_end` no longer causes queries to fail in the `digitalocean_droplet` table ([#24](https://github.com/turbot/steampipe-plugin-digitalocean/pull/24))

## v0.4.0 [2021-07-16]

_What's new?_

- New tables added
  - [digitalocean_app](https://hub.steampipe.io/plugins/turbot/digitalocean/tables/digitalocean_app) ([#18](https://github.com/turbot/steampipe-plugin-digitalocean/pull/18))
  - [digitalocean_domain](https://hub.steampipe.io/plugins/turbot/digitalocean/tables/digitalocean_domain) ([#17](https://github.com/turbot/steampipe-plugin-digitalocean/pull/17))
  - [digitalocean_firewall](https://hub.steampipe.io/plugins/turbot/digitalocean/tables/digitalocean_firewall) ([#15](https://github.com/turbot/steampipe-plugin-digitalocean/pull/15))
  - [digitalocean_kubernetes_cluster](https://hub.steampipe.io/plugins/turbot/digitalocean/tables/digitalocean_kubernetes_cluster) ([#16](https://github.com/turbot/steampipe-plugin-digitalocean/pull/16))

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.3.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v131--2021-07-15) ([#20](https://github.com/turbot/steampipe-plugin-digitalocean/pull/20))
- Updated plugin license to Apache 2.0 per [turbot/steampipe#488](https://github.com/turbot/steampipe/issues/488)

## v0.3.0 [2021-03-11]

_What's new?_

- New tables added
  - [digitalocean_action](https://hub.steampipe.io/plugins/turbot/digitalocean/tables/digitalocean_action)
  - [digitalocean_balance](https://hub.steampipe.io/plugins/turbot/digitalocean/tables/digitalocean_balance)
  - [digitalocean_bill](https://hub.steampipe.io/plugins/turbot/digitalocean/tables/digitalocean_bill)

_Enhancements_
  - Added `private_ipv4`, `public_ipv4`, `public_ipv6` columns to `digitalocean_droplet` table
  - Renamed column `size` to `size_slug` in `digitalocean_load_balancer` table
  - Renamed column `region` to `region_slug` in `digitalocean_load_balancer` table
  - Updated columns using deprecated `ColumnType_DATETIME` type to instead use `ColumnType_TIMESTAMP` type

## v0.2.1 [2021-02-25]

_Bug fixes_

- Recompiled plugin with latest [steampipe-plugin-sdk](https://github.com/turbot/steampipe-plugin-sdk) to resolve SDK issues:
  - Fix error for missing required quals [#40](https://github.com/turbot/steampipe-plugin-sdk/issues/42).
  - Queries fail with error socket: too many open files [#190](https://github.com/turbot/steampipe/issues/190)

## v0.2.0 [2021-02-18]

_What's new?_

- Added support for [connection configuration](https://github.com/turbot/steampipe-plugin-digitalocean/blob/main/docs/index.md#connection-configuration). You may specify digitalocean `token` for each connection in a configuration file. You can have multiple digitalocean connections, each configured for a different account.
