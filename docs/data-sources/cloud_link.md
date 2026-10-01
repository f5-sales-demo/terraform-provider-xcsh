---
page_title: "xcsh_cloud_link landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_cloud_link landing."
---

# xcsh_cloud_link landing

<a id="canonical-3ef122572f65916275a0f8fe0c7d1dc1b10b359850c532ebe64c0f31f1465013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-6edd9427770704acb4eb4aac02b4527829d2d3a96e97f45df8660c4feee0d651"></a>

## xcsh_cloud_link — xcsh_cloud_link / ddebdcd98326 / 2

Breadcrumbs:

- xcsh_cloud_link

Manages new CloudLink with configured parameters in F5 Distributed Cloud.

<a id="canonical-97fc3567177c6eb6ba76431c616c3137ffa835acce82e6b59ec1ddee2fd20b02"></a>

## Prerequisites — xcsh_cloud_link / ddebdcd98326 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-91bb4c315c84d60c27a7596ba3a8425d7977e3003071779eb148c92d48350361"></a>

## Minimal configuration — xcsh_cloud_link / ddebdcd98326 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudLink Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudLink by name
data "xcsh_cloud_link" "example" {
  name      = "example-cloud-link"
  namespace = "staging"
}

output "cloud_link_id" {
  value = data.xcsh_cloud_link.example.id
}
```

<a id="canonical-bbc647e43334cb72cf9f6cf28b40f297afbcdaa6ebe298056eb86c78c9008b48"></a>

## Root configuration — xcsh_cloud_link / ddebdcd98326 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-bbfd903e7fdbcf88aa90b09f9fc07fae1a087904f0fd5fc38e13215df67649e1"></a>

## Next pages — xcsh_cloud_link / ddebdcd98326 / 6

- [Property reference](../guides/data-sources--cloud_link--reference--group-001.md#canonical-670a826d39c9d89e1bf0dbc4088d3a0059f267b1d391c20761b085901989f886)
- [Examples](../guides/data-sources--cloud_link--examples--group-001.md#canonical-1eb9e5d5f9aa54cd66d121744031aeea05feb427e7cf46b447e7f37931955fe1)
