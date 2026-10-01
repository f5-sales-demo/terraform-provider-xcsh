---
page_title: "xcsh_malicious_user_mitigation examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_malicious_user_mitigation examples."
---

# xcsh_malicious_user_mitigation examples

<a id="canonical-bc550721c1f7b8962ff6e6e12da3d5ce2a5e3dd5e0e1105f9540738bee71be27"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-afca644282115db5008fb5ae0a817288da057c9d154548f90eb65b8d52e14b16"></a>

## Examples — Examples / 38f798a974b8 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)
- Examples

<a id="canonical-abde7de6e95035d796f2c36d60fcc2b1883cc1d7cb4028d69b201fdf90c92df1"></a>

## Complete configurations — Examples / 38f798a974b8 / 3

- [Data source](data-sources--malicious_user_mitigation--examples--group-001.md#canonical-851fe37b02305209a869e30645e387a4a167a33384e30b6302efb89ebb8cf7a6): valid configuration.

<a id="canonical-157f7d9d19e88212f1f8d6961f568a914151484dd63edbbc78e4890e431bd865"></a>

## Next pages — Examples / 38f798a974b8 / 4

- [Data source](data-sources--malicious_user_mitigation--examples--group-001.md#canonical-851fe37b02305209a869e30645e387a4a167a33384e30b6302efb89ebb8cf7a6)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)

<a id="canonical-851fe37b02305209a869e30645e387a4a167a33384e30b6302efb89ebb8cf7a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb574d3b18746835ceae9cddf5c3fbc780e2e31d66f5855b6c88e2126685fe67"></a>

## Data source — Data source / b5a4b0f39762 / 2

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)
- [Examples](data-sources--malicious_user_mitigation--examples--group-001.md#canonical-bc550721c1f7b8962ff6e6e12da3d5ce2a5e3dd5e0e1105f9540738bee71be27)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_malicious_user_mitigation/data-source.tf`; digest `sha256:48191839abab419f70a338b5b737300e10e8068606e35e5a0829c82bf6fe5846`.

```terraform
# MaliciousUserMitigation Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing MaliciousUserMitigation by name
data "xcsh_malicious_user_mitigation" "example" {
  name      = "example-malicious-user-mitigation"
  namespace = "staging"
}

output "malicious_user_mitigation_id" {
  value = data.xcsh_malicious_user_mitigation.example.id
}
```

<a id="canonical-1c573d2df4defe2111bb3b82f33388ffc1da6a8d4f1db62f3025a15b78a2898b"></a>

## Next pages — Data source / b5a4b0f39762 / 3

- [Examples](data-sources--malicious_user_mitigation--examples--group-001.md#canonical-bc550721c1f7b8962ff6e6e12da3d5ce2a5e3dd5e0e1105f9540738bee71be27)
- [xcsh_malicious_user_mitigation](../data-sources/malicious_user_mitigation.md#canonical-66a761b26a6c9f489198abfa5c12f059643a1ca2b96cf552db770749ec38dba5)
