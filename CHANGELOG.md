## v0.15.1 [2023-10-04]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#87](https://github.com/turbot/steampipe-plugin-digitalocean/pull/87))

## v0.15.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#85](https://github.com/turbot/steampipe-plugin-digitalocean/pull/85))
- Recompiled plugin with Go version `1.21`. ([#85](https://github.com/turbot/steampipe-plugin-digitalocean/pull/85))

## v0.14.0 [2023-07-17]

_Enhancements_

- Updated the `docs/index.md` file to include multi-account configuration examples. ([#79](https://github.com/turbot/steampipe-plugin-digitalocean/pull/79))

## v0.13.0 [2023-06-20]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.5.0](https://github.com/turbot/steampipe-plugin-sdk/blob/v5.5.0/CHANGELOG.md#v550-2023-06-16) which significantly reduces API calls and boosts query performance, resulting in faster data retrieval. ([#77](https://github.com/turbot/steampipe-plugin-digitalocean/pull/77))

## v0.12.0 [2023-04-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#75](https://github.com/turbot/steampipe-plugin-digitalocean/pull/75))

## v0.11.0 [2023-01-25]

_What's new?_

- New tables added
  - [digitalocean_container_registry](https://hub.steampipe.io/plugins/turbot/digitalocean/tables/digitalocean_container_registry) ([#70](https://github.com/turbot/steampipe-plugin-digitalocean/pull/70))
  - [digitalocean_kubernetes_node_pool](https://hub.steampipe.io/plugins/turbot/digitalocean/tables/digitalocean_kubernetes_node_pool) ([#71](https://github.com/turbot/steampipe-plugin-digitalocean/pull/71))

_Bug fixes_

- Fixed the `digitalocean_snapshot` table to correctly return data instead of an error when an `id` is passed in the `where` clause. ([#69](https://github.com/turbot/steampipe-plugin-digitalocean/pull/69))

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.11](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v4111-2023-01-24) which fixes the issue of non-caching of all the columns of the queried table. ([#72](https://github.com/turbot/steampipe-plugin-digitalocean/pull/72))

## v0.10.0 [2022-09-27]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#63](https://github.com/turbot/steampipe-plugin-digitalocean/pull/63))
- Recompiled plugin with Go version `1.19`. ([#63](https://github.com/turbot/steampipe-plugin-digitalocean/pull/63))

## v0.9.0 [2022-07-13]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v3.3.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v332--2022-07-11) which includes several caching fixes. ([#61](https://github.com/turbot/steampipe-plugin-digitalocean/pull/61))

## v0.8.1 [2022-05-24]

_Bug fixes_

- Fixed the Slack community links in README and docs/index.md files. ([#58](https://github.com/turbot/steampipe-plugin-digitalocean/pull/58))

## v0.8.0 [2022-04-28]

_Enhancements_

- Added support for native Linux ARM and Mac M1 builds. ([#50](https://github.com/turbot/steampipe-plugin-digitalocean/pull/50))
- Recompiled plugin with [steampipe-plugin-sdk v3.1.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v310--2022-03-30) and Go version `1.18`. ([#51](https://github.com/turbot/steampipe-plugin-digitalocean/pull/51))

## v0.7.0 [2022-04-22]

_Enhancements_

- Added column `firewall_rules` to `digitalocean_database` table. ([#48](https://github.com/turbot/steampipe-plugin-digitalocean/pull/48))

_Bug fixes_

- Fixed columns `users` and `db_names` in `digitalocean_database` table to correctly return data. ([#48](https://github.com/turbot/steampipe-plugin-digitalocean/pull/48))
- Updated the data type of column `amount` in `digitalocean_bill` table from `double` to `string`. ([#46](https://github.com/turbot/steampipe-plugin-digitalocean/pull/46))

## v0.6.0 [2021-11-24]

_Enhancements_

- Recompiled plugin with [steampipe-plugin-sdk v1.8.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v182--2021-11-22) and Go version 1.17 ([#41](https://github.com/turbot/steampipe-plugin-digitalocean/pull/41))
- Updated the README.md file to the latest format ([#36](https://github.com/turbot/steampipe-plugin-digitalocean/pull/36))

_Bug fixes_

- Example query updated in `digitalocean_droplet` table ([#43](https://github.com/turbot/steampipe-plugin-digitalocean/pull/43))

## v0.5.0 [2021-08-05]

_What's new?_

- New tables added
  - [digitalocean_alert_policy](https://hub.steampipe.io/plugins/turbot/digitalocean/tables/digitalocean_alert_policy) ([#33](https://github.com/turbot/steampipe-plugin-digitalocean/pull/33))

_Enhancements_

- Updated: Recompiled plugin with [steampipe-plugin-sdk v1.4.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v141--2021-07-20) ([#22](https://github.com/turbot/steampipe-plugin-digitalocean/pull/22))

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
