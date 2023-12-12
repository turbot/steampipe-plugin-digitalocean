---
title: "Steampipe Table: digitalocean_bill - Query DigitalOcean Billing Information using SQL"
description: "Allows users to query Billing Information in DigitalOcean, specifically the monthly usage, cost, and other billing details, providing insights into resource consumption and expenditure."
---

# Table: digitalocean_bill - Query DigitalOcean Billing Information using SQL

DigitalOcean Billing is a feature within DigitalOcean that enables users to track and manage their resource consumption and associated costs. It provides a detailed breakdown of the usage and costs for various DigitalOcean resources, including droplets, volumes, and more. Through DigitalOcean Billing, users can gain insights into their spending patterns and optimize their resource usage accordingly.

## Table Usage Guide

The `digitalocean_bill` table provides insights into the billing information within DigitalOcean. As a DevOps engineer or a finance manager, you can explore detailed billing information through this table, including the cost, usage, and other related details of different resources. Utilize it to uncover information about your spending patterns, identify high-cost resources, and optimize your resource usage to reduce costs.

## Examples

### List all bills
Explore your billing history on DigitalOcean to better understand your usage patterns and costs. This can help in budget management and predicting future expenses.

```sql+postgres
select
  *
from
  digitalocean_bill;
```

```sql+sqlite
select
  *
from
  digitalocean_bill;
```

### Amounts by year
Explore which years had the highest total payments in your DigitalOcean account. This can be useful for financial planning and budgeting purposes.

```sql+postgres
select
  extract(year from date) as year,
  sum(- to_number(amount,'L9G999g999.99')) as payment
from
  digitalocean_bill
where
  type = 'Payment'
group by
  year;
```

```sql+sqlite
select
  strftime('%Y', date) as year,
  sum(- cast(amount as decimal)) as payment
from
  digitalocean_bill
where
  type = 'Payment'
group by
  year;
```