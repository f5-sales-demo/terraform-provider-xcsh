---
page_title: "xcsh_dc_cluster_group landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dc_cluster_group landing."
---

# xcsh_dc_cluster_group landing

<a id="canonical-a0240865dc3b4d8fc779d80a3fc0926a687cf9901210cdea3e46d6b423dc41b7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-97f486d5e0a447d654f602e6c14c252931052ec4e1d6d551730584efc7722366"></a>

## xcsh_dc_cluster_group — xcsh_dc_cluster_group / 68f0e1ee928b / 2

Breadcrumbs:

- xcsh_dc_cluster_group

Manages DC Cluster group in given namespace in F5 Distributed Cloud.

<a id="canonical-e77299dcd6b47582d98a756db054d787ff04c90f4110a8afc110c15b4f7b0aba"></a>

## Prerequisites — xcsh_dc_cluster_group / 68f0e1ee928b / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-a3cc852ded47d5216e390b7696aab80754bbf7c4308bf2694c68d63982cff9af"></a>

## Minimal configuration — xcsh_dc_cluster_group / 68f0e1ee928b / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DcClusterGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DcClusterGroup by name
data "xcsh_dc_cluster_group" "example" {
  name      = "example-dc-cluster-group"
  namespace = "system"
}

output "dc_cluster_group_id" {
  value = data.xcsh_dc_cluster_group.example.id
}
```

<a id="canonical-eb852915e38b6f3504efb3f5504bd4cfaceab0618c388bdc8272da340d3a3795"></a>

## Root configuration — xcsh_dc_cluster_group / 68f0e1ee928b / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-581abc4ae02681e62786bf4cb6505f933accbce3218cb8c7aee22b8efec5d421"></a>

## Next pages — xcsh_dc_cluster_group / 68f0e1ee928b / 6

- [Property reference](../guides/data-sources--dc_cluster_group--reference--group-001.md#canonical-44fdcb354fbb0d410aa13e0c7edd48a9c9ad808c808a37edef7204a79b332e4b)
- [Examples](../guides/data-sources--dc_cluster_group--examples--group-001.md#canonical-4a9ea2aeea2342c033a4b640f81a52986acdf9e373a1343b192be86253d100c9)
