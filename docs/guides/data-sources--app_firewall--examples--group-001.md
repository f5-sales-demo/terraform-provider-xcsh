---
page_title: "xcsh_app_firewall examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_app_firewall examples."
---

# xcsh_app_firewall examples

<a id="canonical-dd79e5b08d935c0b78b3339f9166535800c9d5021a648923f2d47769f1d7a913"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-685f384c1c44362727505320cc908155bfe9b825806a3894d41eed72668cdb0a"></a>

## Examples — Examples / 37ae4c54e73e / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- Examples

<a id="canonical-919b9adf72efd713e3310b4ce9b0ee56d8ebcc31efa587a179449866801e13fe"></a>

## Complete configurations — Examples / 37ae4c54e73e / 3

- [Data source](data-sources--app_firewall--examples--group-001.md#canonical-1a5b7fdea64cca0fb7310d088714d46e7fa70f59f66f6bcd9508c04022b7505b): valid configuration.

<a id="canonical-32444c508b6a0affbd93c9c30a2c7559a2aee0a3bd2ff0b900dc6a9f3344ede1"></a>

## Next pages — Examples / 37ae4c54e73e / 4

- [Data source](data-sources--app_firewall--examples--group-001.md#canonical-1a5b7fdea64cca0fb7310d088714d46e7fa70f59f66f6bcd9508c04022b7505b)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)

<a id="canonical-1a5b7fdea64cca0fb7310d088714d46e7fa70f59f66f6bcd9508c04022b7505b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ff5439052244ec3f6fc10d5b6285c819cd5514faeb2561c03c6eee30cb47288b"></a>

## Data source — Data source / 31aba66db8c0 / 2

Breadcrumbs:

- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
- [Examples](data-sources--app_firewall--examples--group-001.md#canonical-dd79e5b08d935c0b78b3339f9166535800c9d5021a648923f2d47769f1d7a913)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_firewall/data-source.tf`; digest `sha256:3f7d455a2134b0926217d55dc017e3580269822e7d265dbe83b586db534e3931`.

```terraform
# AppFirewall Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppFirewall by name
data "xcsh_app_firewall" "example" {
  name      = "example-app-firewall"
  namespace = "staging"
}

output "app_firewall_id" {
  value = data.xcsh_app_firewall.example.id
}
```

<a id="canonical-dbbce4a461d70c984be31042a9c56508d12c752972a64920ce20eda94a2c613e"></a>

## Next pages — Data source / 31aba66db8c0 / 3

- [Examples](data-sources--app_firewall--examples--group-001.md#canonical-dd79e5b08d935c0b78b3339f9166535800c9d5021a648923f2d47769f1d7a913)
- [xcsh_app_firewall](../data-sources/app_firewall.md#canonical-382e1559c072471efc20c44b05ca9ec952e2d80c4bff0fa01ac43dc03116e4f6)
