---
page_title: "xcsh_alert_policy examples"
subcategory: "Monitoring"
description: "Complete grouped canonical reference for xcsh_alert_policy examples."
---

# xcsh_alert_policy examples

<a id="canonical-3d81f623ab4dcd5d770dbd48b8bef571212d883a505f979b3fabf28a252d1469"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8093bae7bd81a119b062460efd28820bd5c775608ce4cf7db5a658730cc7688b"></a>

## Examples — Examples / fa1529296c8d / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- Examples

<a id="canonical-0126bef2761c33c0cf3bf2d54e27b03c8806197367c25c9a243b74e1d5b28992"></a>

## Complete configurations — Examples / fa1529296c8d / 3

- [Data source](data-sources--alert_policy--examples--group-001.md#canonical-acfc878132e2705e87ae9dbf3b9d6277d227fc13836cdda9afd6133299384e7b): valid configuration.

<a id="canonical-e560fb43b577ea14b83d1a3b0d48b13c264c2f922b502b81eb538df62c64f50f"></a>

## Next pages — Examples / fa1529296c8d / 4

- [Data source](data-sources--alert_policy--examples--group-001.md#canonical-acfc878132e2705e87ae9dbf3b9d6277d227fc13836cdda9afd6133299384e7b)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)

<a id="canonical-acfc878132e2705e87ae9dbf3b9d6277d227fc13836cdda9afd6133299384e7b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65f63ad38c79f98b16d6f3ef544581069e84022df5acd784518af9155731e38a"></a>

## Data source — Data source / 6fa580f6b908 / 2

Breadcrumbs:

- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
- [Examples](data-sources--alert_policy--examples--group-001.md#canonical-3d81f623ab4dcd5d770dbd48b8bef571212d883a505f979b3fabf28a252d1469)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_alert_policy/data-source.tf`; digest `sha256:c6152f124e4da7e8b1a6f94adbb961ca4985954776a9f55f7df6b34b8d892657`.

```terraform
# AlertPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertPolicy by name
data "xcsh_alert_policy" "example" {
  name      = "example-alert-policy"
  namespace = "staging"
}

output "alert_policy_id" {
  value = data.xcsh_alert_policy.example.id
}
```

<a id="canonical-ab1f58fc80a7a82ed73815dc95df901b3eaf0a5bc591725663bea19f475aa330"></a>

## Next pages — Data source / 6fa580f6b908 / 3

- [Examples](data-sources--alert_policy--examples--group-001.md#canonical-3d81f623ab4dcd5d770dbd48b8bef571212d883a505f979b3fabf28a252d1469)
- [xcsh_alert_policy](../data-sources/alert_policy.md#canonical-3839c02b07d955e3e93e0d2e71298906f0e75421074c8fd38c1a7d4e4f56055c)
