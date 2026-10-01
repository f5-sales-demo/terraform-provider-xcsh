---
page_title: "xcsh_addon_service_activation_status examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_addon_service_activation_status examples."
---

# xcsh_addon_service_activation_status examples

<a id="canonical-d291fe604a5df6f19ada2116def1f779abd33918a13909945c2a13c9a8ccc44e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-8176b327f9d0cbdba516db6c08017fe7d329eec6610e403cf8a65ba7fd4621e8"></a>

## Examples — Examples / 89c1e6dba9c2 / 2

Breadcrumbs:

- [xcsh_addon_service_activation_status](../data-sources/addon_service_activation_status.md#canonical-d23bd9a2e3d652134f2d7eae64e927317103e2f5329e77bf01c22114d878a506)
- Examples

<a id="canonical-02fe4098d12515bdef6b83afc1d205772b93d00a1401dd8ab5f8bafc25b32c6f"></a>

## Complete configurations — Examples / 89c1e6dba9c2 / 3

- [Data source](data-sources--addon_service_activation_status--examples--group-001.md#canonical-e72ddea9110096a8144263d97e89fa8fdbd62a3e166fde156fae42b0a566a726): valid configuration.

<a id="canonical-e2ccd5ce07a047d69e52ff5e79b97b41ba13dcfeca4fa48f27981c5d981946ee"></a>

## Next pages — Examples / 89c1e6dba9c2 / 4

- [Data source](data-sources--addon_service_activation_status--examples--group-001.md#canonical-e72ddea9110096a8144263d97e89fa8fdbd62a3e166fde156fae42b0a566a726)
- [xcsh_addon_service_activation_status](../data-sources/addon_service_activation_status.md#canonical-d23bd9a2e3d652134f2d7eae64e927317103e2f5329e77bf01c22114d878a506)

<a id="canonical-e72ddea9110096a8144263d97e89fa8fdbd62a3e166fde156fae42b0a566a726"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-114a3926ceff578fcad8363e6feb00cd7679420d69b567fd46b91aebc33207aa"></a>

## Data source — Data source / 92a563d0e8c3 / 2

Breadcrumbs:

- [xcsh_addon_service_activation_status](../data-sources/addon_service_activation_status.md#canonical-d23bd9a2e3d652134f2d7eae64e927317103e2f5329e77bf01c22114d878a506)
- [Examples](data-sources--addon_service_activation_status--examples--group-001.md#canonical-d291fe604a5df6f19ada2116def1f779abd33918a13909945c2a13c9a8ccc44e)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_addon_service_activation_status/data-source.tf`; digest `sha256:73b6eb2a08b57df6e1279cca2e1bd3688fc9db605ed1feb8ab5a2418a8f37cc0`.

```terraform
# AddonServiceActivationStatus Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Check the tenant's Client-Side Defense subscription.
data "xcsh_addon_service_activation_status" "example" {
  addon_service = "f5xc-client-side-defense-standard"
}

output "addon_service_activation_state" {
  value = data.xcsh_addon_service_activation_status.example.state
}
```

<a id="canonical-97e03241079b0c23b7dcd940fa9e8a4a918121befb81e8aa53401bd32afe8cf3"></a>

## Next pages — Data source / 92a563d0e8c3 / 3

- [Examples](data-sources--addon_service_activation_status--examples--group-001.md#canonical-d291fe604a5df6f19ada2116def1f779abd33918a13909945c2a13c9a8ccc44e)
- [xcsh_addon_service_activation_status](../data-sources/addon_service_activation_status.md#canonical-d23bd9a2e3d652134f2d7eae64e927317103e2f5329e77bf01c22114d878a506)
