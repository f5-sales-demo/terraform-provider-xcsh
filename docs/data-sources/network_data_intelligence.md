---
page_title: "xcsh_network_data_intelligence landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_data_intelligence landing."
---

# xcsh_network_data_intelligence landing

<a id="canonical-db1c6e7c4254880692807159269c6d76d21575ca6575b838319cf6c66703a647"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-59e126ba9b135ea80cbda0c0529eced68485382e62b4c15306b9533d2933da85"></a>

## xcsh_network_data_intelligence — xcsh_network_data_intelligence / fd62344ac446 / 2

Breadcrumbs:

- xcsh_network_data_intelligence

Regional Data Intelligence IPv4 destinations. Values are bundled from the pinned OpenAPI release;
this data source performs no network request. Ports and traffic direction are not encoded in the
manifest.

<a id="canonical-e511a3d17540fe8c1f3856025adf9162b0753eb7bf63e2f3d3a2b3fd57ba4a53"></a>

## Prerequisites — xcsh_network_data_intelligence / fd62344ac446 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-f3f06f22277a5d882e256aaccde30c436b22888c2140db8cc91dc2091e340b3d"></a>

## Minimal configuration — xcsh_network_data_intelligence / fd62344ac446 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_data_intelligence" "us" {
  regions = ["us"]
}

output "data_intelligence_https_egress" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 443
    destinations = data.xcsh_network_data_intelligence.us.cidr_blocks
  }
}
```

<a id="canonical-9690c14b478d056ae81009fe947567f4d787c49cf8fd72c7056701e11efd0716"></a>

## Root configuration — xcsh_network_data_intelligence / fd62344ac446 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-a6526b2d1f4e5ad87f66273768eca1067b0f291db7234763b9c86cb36d6523bc"></a>

## Next pages — xcsh_network_data_intelligence / fd62344ac446 / 6

- [Property reference](../guides/data-sources--network_data_intelligence--reference--group-001.md#canonical-c658a6fbfe8dbbe5fd2dab0eb03fbfc881b85553026886291a106a251c844753)
- [Examples](../guides/data-sources--network_data_intelligence--examples--group-001.md#canonical-3b7b4a98c89d8243a85c0a9412ae43f816de2f4deca4d3e685794d6ea3227e8b)
