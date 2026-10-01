---
page_title: "xcsh_nat_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nat_policy examples."
---

# xcsh_nat_policy examples

<a id="canonical-29569badf4ac6533a49f84111b8f371cb2a24d15313f0d6ee8386e1a348410c6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d55c07878f57a1734eb5b0c2d0f32f5ddfb2154061b698c9e7077499026519be"></a>

## Examples — Examples / b1b039eb9887 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- Examples

<a id="canonical-c10f616a92975e3581a80b10da6ee3ce6b373282d07f917a86b3e0fbfa2bc8ef"></a>

## Complete configurations — Examples / b1b039eb9887 / 3

- [Data source](data-sources--nat_policy--examples--group-001.md#canonical-90c5e62d94778a72f26a1d0f5e7e03f1f0634d0960fe7ec96d4e9db14e55e2f8): valid configuration.

<a id="canonical-702d1e9ceced4dadf834805277d27a774671ce6847761283b7b7a8266d3b1198"></a>

## Next pages — Examples / b1b039eb9887 / 4

- [Data source](data-sources--nat_policy--examples--group-001.md#canonical-90c5e62d94778a72f26a1d0f5e7e03f1f0634d0960fe7ec96d4e9db14e55e2f8)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)

<a id="canonical-90c5e62d94778a72f26a1d0f5e7e03f1f0634d0960fe7ec96d4e9db14e55e2f8"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-60f928e5b03299f22b517c9fccb0534fdcbfa692f6e6fa09ca3d16df751a48d5"></a>

## Data source — Data source / 0565abf06278 / 2

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
- [Examples](data-sources--nat_policy--examples--group-001.md#canonical-29569badf4ac6533a49f84111b8f371cb2a24d15313f0d6ee8386e1a348410c6)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nat_policy/data-source.tf`; digest `sha256:d086b598bab70cc64257bc21568a65534a135b8fdb72d6f6952421f88e8adafe`.

```terraform
# NATPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NATPolicy by name
data "xcsh_nat_policy" "example" {
  name      = "example-nat-policy"
  namespace = "staging"
}

output "nat_policy_id" {
  value = data.xcsh_nat_policy.example.id
}
```

<a id="canonical-c585f2fea2a3eb612db5a546076db6622e86556fae7d38c2192e3d4d26414d9f"></a>

## Next pages — Data source / 0565abf06278 / 3

- [Examples](data-sources--nat_policy--examples--group-001.md#canonical-29569badf4ac6533a49f84111b8f371cb2a24d15313f0d6ee8386e1a348410c6)
- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-a5fec5294bc8787088593c2d4f32bac4080e871efb271be14c772afab23ddd19)
