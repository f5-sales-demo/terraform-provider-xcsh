---
page_title: "xcsh_fast_acl_rule examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl_rule examples."
---

# xcsh_fast_acl_rule examples

<a id="canonical-e5b9c4665cbf18fbdc5dded70558f8fea91828cc91f8581313bd641cf26d9b2c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dcc96182f680a00bd58237c55a630023bdc2d83e96620bf5425fe1e2a81b8dd3"></a>

## Examples — Examples / c667c15f9ba9 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
- Examples

<a id="canonical-91f456b3a827534082a46926d28112d2b2834f916ed3115abe76327ff1664794"></a>

## Complete configurations — Examples / c667c15f9ba9 / 3

- [Data source](data-sources--fast_acl_rule--examples--group-001.md#canonical-52388dc25cdc86169c017aaec49a1bc903a0eb2fd521f698b0915c2ed1170a24): valid configuration.

<a id="canonical-2b347d0d4cf3ca471764831e937a3381ac34eda3a3c8c947f350363f809120fc"></a>

## Next pages — Examples / c667c15f9ba9 / 4

- [Data source](data-sources--fast_acl_rule--examples--group-001.md#canonical-52388dc25cdc86169c017aaec49a1bc903a0eb2fd521f698b0915c2ed1170a24)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)

<a id="canonical-52388dc25cdc86169c017aaec49a1bc903a0eb2fd521f698b0915c2ed1170a24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5af2852d2ab90f7059d72e29b6850c6985ece0367939f58e930ba1adc2dbe1f0"></a>

## Data source — Data source / 207503d35a04 / 2

Breadcrumbs:

- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
- [Examples](data-sources--fast_acl_rule--examples--group-001.md#canonical-e5b9c4665cbf18fbdc5dded70558f8fea91828cc91f8581313bd641cf26d9b2c)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_fast_acl_rule/data-source.tf`; digest `sha256:d3a85aa2d2c4a98927959b827b07db2533ed58dbfde3c28ad82e3e2235d3a843`.

```terraform
# FastACLRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FastACLRule by name
data "xcsh_fast_acl_rule" "example" {
  name      = "example-fast-acl-rule"
  namespace = "staging"
}

output "fast_acl_rule_id" {
  value = data.xcsh_fast_acl_rule.example.id
}
```

<a id="canonical-67645a606ffba3c37ef9f95eb03bb575c3786bc590adb066cc4bdead64316a19"></a>

## Next pages — Data source / 207503d35a04 / 3

- [Examples](data-sources--fast_acl_rule--examples--group-001.md#canonical-e5b9c4665cbf18fbdc5dded70558f8fea91828cc91f8581313bd641cf26d9b2c)
- [xcsh_fast_acl_rule](../data-sources/fast_acl_rule.md#canonical-ba96eb5b80ab20ba2cee953fc8858074c44adb8ee3a4d0ec70f6a52e8b76e10e)
