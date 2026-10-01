---
page_title: "xcsh_device_intelligence_summary landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_summary landing."
---

# xcsh_device_intelligence_summary landing

<a id="canonical-713f059bbd4c71f5f978f47874a8258e6f5fed5f744224aa8b9189615842fff9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-a7b2c5365e7266142fce0f7137b55b108510880da757ad6596b957d955cd7273"></a>

## xcsh_device_intelligence_summary — xcsh_device_intelligence_summary / 3f3fe450c7d3 / 2

Breadcrumbs:

- xcsh_device_intelligence_summary

Resource creation operation.

<a id="canonical-96db13bbb93694dd54dc9c2b244d42c625ae2d199079d3cd35dea2f38f8267d5"></a>

## Prerequisites — xcsh_device_intelligence_summary / 3f3fe450c7d3 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-4ec2d3d309edfc167d616f9a79e3e52cb23a6053b741ef1e0412ad91d63bf50d"></a>

## Minimal configuration — xcsh_device_intelligence_summary / 3f3fe450c7d3 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceSummary DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_summary" "example" {
  namespace = "example-value"
}

output "device_intelligence_summary_result" {
  value = data.xcsh_device_intelligence_summary.example
}
```

<a id="canonical-83fd355cc873f019f67a4fece5033aff14793ab1f0a8ed643fa9d892d7416bf8"></a>

## Root configuration — xcsh_device_intelligence_summary / 3f3fe450c7d3 / 5

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-dc2016264d48b1f9833e29cccb655ab83f0fac98f226611ac2769bc4a908642a"></a>

## Next pages — xcsh_device_intelligence_summary / 3f3fe450c7d3 / 6

- [Property reference](../guides/data-sources--device_intelligence_summary--reference--group-001.md#canonical-1c701a3bcd569781ea4b09b023243540a7da6f24b4b2c73c27e1755ab7ac2705)
- [Examples](../guides/data-sources--device_intelligence_summary--examples--group-001.md#canonical-f97a8f9c74d388e7b38af72d6e8c5345b93515d764c0d104c74440b774b7228c)
