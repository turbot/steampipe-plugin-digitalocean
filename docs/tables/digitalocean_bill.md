# Table: digitalocean_bill

Billing history is a record of billing events for your account. For example,
entries may include events like payments made, invoices issued, or credits
granted.

## Examples

### List all bills

```sql
select
  *
from
  digitalocean_bill;
```

### Amounts by year

```sql
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
