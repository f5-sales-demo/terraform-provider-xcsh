---
page_title: "xcsh_mitigated_domain examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_mitigated_domain examples."
---

# xcsh_mitigated_domain examples

<a id="canonical-6c61924190f11f3700744be4ef278107ba6451dcbc41246a1e2d9701afbbda24"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2f6c1c31dc60a0664a2105bbd9d257d9a9e9944cdb7bb2b2270791719606d2b4"></a>

## Examples — Examples / d8ce071ecd10 / 2

Breadcrumbs:

- [xcsh_mitigated_domain](../data-sources/mitigated_domain.md#canonical-efcfbe08e10eae8cd34e645079b3b2ddf4a333bb259bfc7f7c6325c8868a5fb4)
- Examples

<a id="canonical-09d87d5ac15fb44def18dc6f9535b6834b9ec03cd3dd6a335425fbf8bfa3001c"></a>

## Complete configurations — Examples / d8ce071ecd10 / 3

- [Data source](data-sources--mitigated_domain--examples--group-001.md#canonical-f7313658f187e0771348f257e2b9cfcc650608ce5c331cb7f3bb44150324e8a6): valid configuration.

<a id="canonical-1197000fcc26abd1fdf892d0beb8412df71250c691bb79b316f32f98daba0fda"></a>

## Next pages — Examples / d8ce071ecd10 / 4

- [Data source](data-sources--mitigated_domain--examples--group-001.md#canonical-f7313658f187e0771348f257e2b9cfcc650608ce5c331cb7f3bb44150324e8a6)
- [xcsh_mitigated_domain](../data-sources/mitigated_domain.md#canonical-efcfbe08e10eae8cd34e645079b3b2ddf4a333bb259bfc7f7c6325c8868a5fb4)

<a id="canonical-f7313658f187e0771348f257e2b9cfcc650608ce5c331cb7f3bb44150324e8a6"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d616851deac63f381990c6bbc4c1ca188c2f918a9ef7a6ae2ab3cd4875ae8314"></a>

## Data source — Data source / e28b113a77d7 / 2

Breadcrumbs:

- [xcsh_mitigated_domain](../data-sources/mitigated_domain.md#canonical-efcfbe08e10eae8cd34e645079b3b2ddf4a333bb259bfc7f7c6325c8868a5fb4)
- [Examples](data-sources--mitigated_domain--examples--group-001.md#canonical-6c61924190f11f3700744be4ef278107ba6451dcbc41246a1e2d9701afbbda24)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_mitigated_domain/data-source.tf`; digest `sha256:bc0c9edeff4c50185f9f5d99bbe4c09ceec7f010ca7ea7432a8786d884607a16`.

```terraform
# MitigatedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing MitigatedDomain by name
data "xcsh_mitigated_domain" "example" {
  name      = "example-mitigated-domain"
  namespace = "staging"
}

output "mitigated_domain_id" {
  value = data.xcsh_mitigated_domain.example.id
}
```

<a id="canonical-1704ba82da8203e751315a1a43fe4a42bf3d91725232b934fd555a2297b242c0"></a>

## Next pages — Data source / e28b113a77d7 / 3

- [Examples](data-sources--mitigated_domain--examples--group-001.md#canonical-6c61924190f11f3700744be4ef278107ba6451dcbc41246a1e2d9701afbbda24)
- [xcsh_mitigated_domain](../data-sources/mitigated_domain.md#canonical-efcfbe08e10eae8cd34e645079b3b2ddf4a333bb259bfc7f7c6325c8868a5fb4)
