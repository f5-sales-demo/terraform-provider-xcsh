---
page_title: "xcsh_site_registrations_by_site landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_registrations_by_site landing."
---

# xcsh_site_registrations_by_site landing

<a id="canonical-2111303110111030-3113020121012123-1313322201232310-1013111203201222-0203220021201023-0033331330113033-0303303330002221-2203131220023320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2210031203322011-1213203312333311-3113233013312022-2031111312112021-1331310031021332-3113232200103120-3112312221303313-1110121132111132"></a>

## xcsh_site_registrations_by_site — xcsh_site_registrations_by_site / 002231022331 / 2

Breadcrumbs:

- xcsh_site_registrations_by_site

List registrations for a Customer Edge site.

<a id="canonical-2002222000020030-2223233321330321-3313211013031122-1033321133020210-1211211023231100-3330323312221331-1312122211300203-3302223022322232"></a>

## Prerequisites — xcsh_site_registrations_by_site / 002231022331 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0321112223112311-1320030110313332-3101231110130213-1102223301310031-2302013302133113-2031103021200320-1121213021202232-3021031003130312"></a>

## Minimal configuration — xcsh_site_registrations_by_site / 002231022331 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-1030310312332303-2001022302312212-1301130330123131-3203313313223133-1013320202333023-0113030010112020-2322001203100000-1023210323121130"></a>

## Root configuration — xcsh_site_registrations_by_site / 002231022331 / 5

Required root properties: `site_name`. Full root flags and choices appear in the property reference.

<a id="canonical-2002122320003121-1220112231201103-3112221023231011-2313231010233332-3101111232032120-3233203301331020-3113331331210033-1322322033112323"></a>

## Next pages — xcsh_site_registrations_by_site / 002231022331 / 6

- [Property reference](../guides/data-sources--site_registrations_by_site--reference--group-001.md#canonical-0102102012000331-3300022232033212-0133030232322023-1233323203011300-3213211013132213-1000132330333132-1212223310311330-0012021313023000)
- [Examples](../guides/data-sources--site_registrations_by_site--examples--group-001.md#canonical-2131333210003322-0031231330102031-3322130131132010-3220101013333221-0021233201013123-2302102123320121-1112122112332133-0311303233010233)
