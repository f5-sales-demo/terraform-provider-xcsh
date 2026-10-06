---
page_title: "xcsh_site_upgrade_sw examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_upgrade_sw examples."
---

# xcsh_site_upgrade_sw examples

<a id="canonical-1133112323300000-3332003131302131-2300100101022322-3033010311201200-2203211331101220-1232230303131300-3331212323021122-2322203021032110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_site_upgrade_sw](../actions/site_upgrade_sw.md#canonical-0120130312312320-0223010000032113-2233123110100201-3130110122132200-2323212123211333-3120313220122233-0023002110221112-0320032112111211)
- Examples

<a id="canonical-2031010301103312-0333022102212311-2110333013302122-2011033002303030-0102313110322210-0311202032212302-1001200033311333-2131122012310101"></a>

### Complete configurations for `xcsh_site_upgrade_sw`

- [Action](actions--site_upgrade_sw--examples--group-001.md#canonical-1323323103032321-2101300321023232-2032302101012302-2313102031131130-1311333332310113-3013030213333101-0223213103012110-0013302000322232): valid configuration.

<a id="canonical-1323323103032321-2101300321023232-2032302101012302-2313102031131130-1311333332310113-3013030213333101-0223213103012110-0013302000322232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Action example

Breadcrumbs:

- [xcsh_site_upgrade_sw](../actions/site_upgrade_sw.md#canonical-0120130312312320-0223010000032113-2233123110100201-3130110122132200-2323212123211333-3120313220122233-0023002110221112-0320032112111211)
- [Examples](actions--site_upgrade_sw--examples--group-001.md#canonical-1133112323300000-3332003131302131-2300100101022322-3033010311201200-2203211331101220-1232230303131300-3331212323021122-2322203021032110)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_site_upgrade_sw/action.tf`; digest `sha256:f543d404a49e0f6ff11128bf33aab9d859e0913de2e4fcdcef829bf992c11087`.

```terraform
# SiteUpgradeSw Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# The API accepts the upgrade request immediately; convergence is asynchronous.
# This action does not reconcile a site's pinned software_settings.
action "xcsh_site_upgrade_sw" "example" {
  config {
    site             = "example-value"
    software_version = "example-value"
  }
}
```
