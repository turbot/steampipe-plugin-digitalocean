---
title: "Steampipe Table: digitalocean_balance - Query DigitalOcean Balances using SQL"
description: "Allows users to query DigitalOcean Balances, providing insights into the account's balance and month-to-date usage."
---

# Table: digitalocean_balance - Query DigitalOcean Balances using SQL

DigitalOcean Balance is a feature that allows users to monitor their account's balance and usage. It provides a detailed view of the account's current balance, the amount used in the current billing month, and the projected amount for the next billing month. This feature is useful for managing expenses and understanding the cost implications of DigitalOcean resources.

## Table Usage Guide

The `digitalocean_balance` table provides insights into the account's balance and usage within DigitalOcean. As an account manager or financial analyst, explore balance-specific details through this table, including the current balance, month-to-date usage, and projected next month's usage. Utilize it to manage expenses, understand cost implications, and plan for future resource usage.

## Examples

### Get the balance
Explore your DigitalOcean account's current balance to understand your usage costs and manage your budget effectively. This helps you keep track of your expenses and plan for future resource allocation.

```sql+postgres
select
  *
from
  digitalocean_balance;
```

```sql+sqlite
select
  *
from
  digitalocean_balance;
```