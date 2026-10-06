---
page_title: "xcsh_site_registrations_by_site examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations_by_site examples."
---

# xcsh_site_registrations_by_site examples

<a id="canonical-2131333210003322-0031231330102031-3322130131132010-3220101013333221-0021233201013123-2302102123320121-1112122112332133-0311303233010233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-2111303110111030-3113020121012123-1313322201232310-1013111203201222-0203220021201023-0033331330113033-0303303330002221-2203131220023320)
- Examples

<a id="canonical-2112031200122232-3132030122113101-0102021213112313-3121301323320330-3003222311202302-1311131121332023-3231330331330311-2110002000001112"></a>

### Complete configurations for `xcsh_site_registrations_by_site`

- [Data source](data-sources--site_registrations_by_site--examples--group-001.md#canonical-3331332010233120-2013201132100112-3203020223211211-1111333330212103-0032231210120130-0033023333220123-3213303212230122-0133100331112311): valid configuration.

<a id="canonical-3331332010233120-2013201132100112-3203020223211211-1111333330212103-0032231210120130-0033023333220123-3213303212230122-0133100331112311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_site_registrations_by_site](../data-sources/site_registrations_by_site.md#canonical-2111303110111030-3113020121012123-1313322201232310-1013111203201222-0203220021201023-0033331330113033-0303303330002221-2203131220023320)
- [Examples](data-sources--site_registrations_by_site--examples--group-001.md#canonical-2131333210003322-0031231330102031-3322130131132010-3220101013333221-0021233201013123-2302102123320121-1112122112332133-0311303233010233)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_registrations_by_site/data-source.tf`; digest `sha256:c5f91ec913e86f87149f588895f436b17703e7ae2ae9655f9226a6f3b281b896`.

```terraform
# SiteRegistrationsBySite DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations_by_site" "example" {
  site_name = "example-value"
}

output "site_registrations_by_site_result" {
  value = data.xcsh_site_registrations_by_site.example
}
```
