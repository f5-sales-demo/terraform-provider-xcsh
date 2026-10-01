---
page_title: "xcsh_certificate examples"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_certificate examples."
---

# xcsh_certificate examples

<a id="canonical-3edd44319babaa4ce9d9f018bf81a77c5532f1bef794cec91c914fb336d52473"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-65f2bc06fc6634d9caf80dfc6b6a1388e9b457c73e170ae4be4a0d44debaaa99"></a>

## Examples — Examples / db38f3fb0418 / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)
- Examples

<a id="canonical-314e8cff9d0da565c0eea97be42935440b6be92b9757351af919c2ca8b857684"></a>

## Complete configurations — Examples / db38f3fb0418 / 3

- [Data source](data-sources--certificate--examples--group-001.md#canonical-2ceb677450ac955f870e0d98c130595ba486d0a7d1fcbbdb98ef65efd39a7e72): valid configuration.

<a id="canonical-fd27dad06fb8ea111eb9ffb56d437065971cd068d157acb2f70f455fa81351b8"></a>

## Next pages — Examples / db38f3fb0418 / 4

- [Data source](data-sources--certificate--examples--group-001.md#canonical-2ceb677450ac955f870e0d98c130595ba486d0a7d1fcbbdb98ef65efd39a7e72)
- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)

<a id="canonical-2ceb677450ac955f870e0d98c130595ba486d0a7d1fcbbdb98ef65efd39a7e72"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e5968d9ecd7bd798019a273028ef2f8f275183e464df3c8292926773164aa515"></a>

## Data source — Data source / 3bb68962fd86 / 2

Breadcrumbs:

- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)
- [Examples](data-sources--certificate--examples--group-001.md#canonical-3edd44319babaa4ce9d9f018bf81a77c5532f1bef794cec91c914fb336d52473)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_certificate/data-source.tf`; digest `sha256:ae1b4b8d1a231d986bf4b3d7767f79709ca2d85beaf03f7ea95f9a72d8a1d03e`.

```terraform
# Certificate Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Certificate by name
data "xcsh_certificate" "example" {
  name      = "example-certificate"
  namespace = "staging"
}

output "certificate_id" {
  value = data.xcsh_certificate.example.id
}
```

<a id="canonical-48ae76ba7730c2b32fdafe026e123893406d6a63544cbfd0710d9802008dc9c5"></a>

## Next pages — Data source / 3bb68962fd86 / 3

- [Examples](data-sources--certificate--examples--group-001.md#canonical-3edd44319babaa4ce9d9f018bf81a77c5532f1bef794cec91c914fb336d52473)
- [xcsh_certificate](../data-sources/certificate.md#canonical-c6e407e3aeb9050c5fa1ed5568b0cf55574fb4bfaf33eb473a67643d1944fe84)
