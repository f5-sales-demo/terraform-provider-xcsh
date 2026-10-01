---
page_title: "xcsh_site_signatures_update examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_signatures_update examples."
---

# xcsh_site_signatures_update examples

<a id="canonical-5f9409c365fddfafc467f46ef8caa853731689277c363e22f35bd55e440ba1d9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-13da37ef23e75246e6ca30edfb07cbcbdd6a544e17905cbc5c45c835c9ac1412"></a>

## Examples — Examples / 39bb4429e537 / 2

Breadcrumbs:

- [xcsh_site_signatures_update](../actions/site_signatures_update.md#canonical-997cfc6e0541d9b7078c9e8c892fe43f791342c6fe0d09cd7f128fc478b34edf)
- Examples

<a id="canonical-1bfce752fe3042b1a17e7818b4a4e85d1a6adf17230fb6a4981b3760499ecd8c"></a>

## Complete configurations — Examples / 39bb4429e537 / 3

- [Action](actions--site_signatures_update--examples--group-001.md#canonical-c8c3884502c30c709f6d9edeed3a977d5fa45dbe217da45eb35263703b2b9153): valid configuration.

<a id="canonical-ee5c582bebf55cf909a65524348dfc628f961ba70b329543090caf5e743c06c6"></a>

## Next pages — Examples / 39bb4429e537 / 4

- [Action](actions--site_signatures_update--examples--group-001.md#canonical-c8c3884502c30c709f6d9edeed3a977d5fa45dbe217da45eb35263703b2b9153)
- [xcsh_site_signatures_update](../actions/site_signatures_update.md#canonical-997cfc6e0541d9b7078c9e8c892fe43f791342c6fe0d09cd7f128fc478b34edf)

<a id="canonical-c8c3884502c30c709f6d9edeed3a977d5fa45dbe217da45eb35263703b2b9153"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d3eb614a3e70146d4d25e58686ab630bf90b34a8dbe9bea782ef0e7bf4efb1d9"></a>

## Action — Action / 0c5fae59bdbf / 2

Breadcrumbs:

- [xcsh_site_signatures_update](../actions/site_signatures_update.md#canonical-997cfc6e0541d9b7078c9e8c892fe43f791342c6fe0d09cd7f128fc478b34edf)
- [Examples](actions--site_signatures_update--examples--group-001.md#canonical-5f9409c365fddfafc467f46ef8caa853731689277c363e22f35bd55e440ba1d9)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_site_signatures_update/action.tf`; digest `sha256:b8fc93938ea82d6d38f39cc7c48dce93f107714b65d36ac9a46938636cbed4f9`.

```terraform
# SiteSignaturesUpdate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_site_signatures_update" "example" {
  config {
    namespace = "example-value"
  }
}
```

<a id="canonical-8b5f4729498b20513a3c979e7b28e0d7f620d25c00be8b52d4356f985dfd2d82"></a>

## Next pages — Action / 0c5fae59bdbf / 3

- [Examples](actions--site_signatures_update--examples--group-001.md#canonical-5f9409c365fddfafc467f46ef8caa853731689277c363e22f35bd55e440ba1d9)
- [xcsh_site_signatures_update](../actions/site_signatures_update.md#canonical-997cfc6e0541d9b7078c9e8c892fe43f791342c6fe0d09cd7f128fc478b34edf)
