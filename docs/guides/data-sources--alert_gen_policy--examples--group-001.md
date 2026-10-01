---
page_title: "xcsh_alert_gen_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_alert_gen_policy examples."
---

# xcsh_alert_gen_policy examples

<a id="canonical-d17d447e66dec82f267cc359c03ee17dc2e216a459c21d2dd2a3110f30fcc8e4"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a0250289844fa1f19dbc00a7dc943d2dd9f87fd2af47d46b0c66a1111613f22f"></a>

## Examples — Examples / 3e4ca0ea0b1f / 2

Breadcrumbs:

- [xcsh_alert_gen_policy](../data-sources/alert_gen_policy.md#canonical-66af5b4db37114bd99ef2798c10e6b2faf4f99657029d35c516b10fefc2f1aa8)
- Examples

<a id="canonical-18d2bc98c9226377106e9fa7754a5d91d006462a31b0983826b30f316f8ea7bd"></a>

## Complete configurations — Examples / 3e4ca0ea0b1f / 3

- [Data source](data-sources--alert_gen_policy--examples--group-001.md#canonical-532d48138a52cf5599c640f03f80434f87f14969b751f91e86ab4aa88e9f02b6): valid configuration.

<a id="canonical-408a41c5cdf99948caf959d97f56473d9180558b1d1be7403533a811df4c28c8"></a>

## Next pages — Examples / 3e4ca0ea0b1f / 4

- [Data source](data-sources--alert_gen_policy--examples--group-001.md#canonical-532d48138a52cf5599c640f03f80434f87f14969b751f91e86ab4aa88e9f02b6)
- [xcsh_alert_gen_policy](../data-sources/alert_gen_policy.md#canonical-66af5b4db37114bd99ef2798c10e6b2faf4f99657029d35c516b10fefc2f1aa8)

<a id="canonical-532d48138a52cf5599c640f03f80434f87f14969b751f91e86ab4aa88e9f02b6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6fcf7b441a7e07da5a530873bed2a75364cb2a7f9bc4985e8465ca09de500f12"></a>

## Data source — Data source / 9fd46eb79870 / 2

Breadcrumbs:

- [xcsh_alert_gen_policy](../data-sources/alert_gen_policy.md#canonical-66af5b4db37114bd99ef2798c10e6b2faf4f99657029d35c516b10fefc2f1aa8)
- [Examples](data-sources--alert_gen_policy--examples--group-001.md#canonical-d17d447e66dec82f267cc359c03ee17dc2e216a459c21d2dd2a3110f30fcc8e4)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_alert_gen_policy/data-source.tf`; digest `sha256:349cbe33e45f53c629e7ce220830862d4e4d7dd9b1fd6b5c5a42c3556d1a1bed`.

```terraform
# AlertGenPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AlertGenPolicy by name
data "xcsh_alert_gen_policy" "example" {
  name      = "example-alert-gen-policy"
  namespace = "staging"
}

output "alert_gen_policy_id" {
  value = data.xcsh_alert_gen_policy.example.id
}
```

<a id="canonical-1bd986441767db5d0bef86c064c2b6955b8506977b8f03095e983de3e48f0fd3"></a>

## Next pages — Data source / 9fd46eb79870 / 3

- [Examples](data-sources--alert_gen_policy--examples--group-001.md#canonical-d17d447e66dec82f267cc359c03ee17dc2e216a459c21d2dd2a3110f30fcc8e4)
- [xcsh_alert_gen_policy](../data-sources/alert_gen_policy.md#canonical-66af5b4db37114bd99ef2798c10e6b2faf4f99657029d35c516b10fefc2f1aa8)
