---
page_title: "xcsh_rate_limiter_policy examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_rate_limiter_policy examples."
---

# xcsh_rate_limiter_policy examples

<a id="canonical-9af59c23111ef40dd083e7310441f58a2c7e0e1ad5fb8e562709dd4ea7f2db9f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7a15787e95c9011924140391deb8394c8f8817d2280d91b54a59aeb04ae2a3fd"></a>

## Examples — Examples / d8144d9f5022 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- Examples

<a id="canonical-185ac85dea4cbfa594ac22954e6588e898b9df80c24d77e9ef87856eb416969e"></a>

## Complete configurations — Examples / d8144d9f5022 / 3

- [Data source](data-sources--rate_limiter_policy--examples--group-001.md#canonical-ee781cab2246deb3990c22fc8e012d6de9137dbe53690ee56ec73851735a9557): valid configuration.

<a id="canonical-c18b91d2c4ef291da7202d14ca4d8e3935246212d5db1f236c693199d0b01d93"></a>

## Next pages — Examples / d8144d9f5022 / 4

- [Data source](data-sources--rate_limiter_policy--examples--group-001.md#canonical-ee781cab2246deb3990c22fc8e012d6de9137dbe53690ee56ec73851735a9557)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)

<a id="canonical-ee781cab2246deb3990c22fc8e012d6de9137dbe53690ee56ec73851735a9557"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6294f5a4ed0d3f76a0fab32e8e723b34f5c3be09628badee151e82f00e0f39bb"></a>

## Data source — Data source / 44ca5713e397 / 2

Breadcrumbs:

- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
- [Examples](data-sources--rate_limiter_policy--examples--group-001.md#canonical-9af59c23111ef40dd083e7310441f58a2c7e0e1ad5fb8e562709dd4ea7f2db9f)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_rate_limiter_policy/data-source.tf`; digest `sha256:631ebdd1844abfdc3c403666a18c881b536cd51ce00cfb0a5d316458e5a24bba`.

```terraform
# RateLimiterPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing RateLimiterPolicy by name
data "xcsh_rate_limiter_policy" "example" {
  name      = "example-rate-limiter-policy"
  namespace = "staging"
}

output "rate_limiter_policy_id" {
  value = data.xcsh_rate_limiter_policy.example.id
}
```

<a id="canonical-687c62aba3a7eeb0750c455a1f8bb0b32c48fc73500bc5f816ba6e85f1e3e4c2"></a>

## Next pages — Data source / 44ca5713e397 / 3

- [Examples](data-sources--rate_limiter_policy--examples--group-001.md#canonical-9af59c23111ef40dd083e7310441f58a2c7e0e1ad5fb8e562709dd4ea7f2db9f)
- [xcsh_rate_limiter_policy](../data-sources/rate_limiter_policy.md#canonical-d85e7ee161a0d5147d4a3d4ca16bb93ca846a6e3f718c829bcd6aed4d7147bd4)
