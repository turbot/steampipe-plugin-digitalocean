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
