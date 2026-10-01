---
page_title: "xcsh_network_customer_edge_egress landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_customer_edge_egress landing."
---

# xcsh_network_customer_edge_egress landing

<a id="canonical-e8a53f08d4a40bf225466083460717f25dabd979bfa6a17a3cab3b3fbaa24138"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-dc31891e547e5cca0a4f4643b7fe8de85689056be07cdf163e59fc50a6907ccc"></a>

## xcsh_network_customer_edge_egress — xcsh_network_customer_edge_egress / 79cd67327bcf / 2

Breadcrumbs:

- xcsh_network_customer_edge_egress

Secure Mesh v2 registration IPv4 addresses and egress domains. Legacy Customer Edge values are
intentionally excluded. Values are bundled from the pinned OpenAPI release; this data source
performs no network request. Ports and traffic direction are not encoded in the manifest.

<a id="canonical-847fb8a4cdc13a37549f402626ad294e992ea6bf6188a2d457242de7e47db7cd"></a>

## Prerequisites — xcsh_network_customer_edge_egress / 79cd67327bcf / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-51488a15ef25df60c331b04c7ee84e99fb15900f288739ee69524c7cd95e0f34"></a>

## Minimal configuration — xcsh_network_customer_edge_egress / 79cd67327bcf / 4

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

data "xcsh_network_customer_edge_egress" "secure_mesh_v2" {}

# Use the domains with an FQDN-aware control. The legacy CE branch is not
# included in this data source.
output "secure_mesh_v2_https_egress" {
  value = {
    direction              = "egress"
    protocol               = "tcp"
    port                   = 443
    registration_addresses = data.xcsh_network_customer_edge_egress.secure_mesh_v2.registration_addresses
    domains                = data.xcsh_network_customer_edge_egress.secure_mesh_v2.domains
  }
}
```

<a id="canonical-aa1740bc39ce66198a119d6aab4d0aa5335a98a9c22cf9ace5b0a5019c6cda14"></a>

## Root configuration — xcsh_network_customer_edge_egress / 79cd67327bcf / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-148c83ae26e9778516deee4853b0e8bf48aa5d7c65e4877043376b5b40467b99"></a>

## Next pages — xcsh_network_customer_edge_egress / 79cd67327bcf / 6

- [Property reference](../guides/data-sources--network_customer_edge_egress--reference--group-001.md#canonical-fcf3f94ae406bb1faa761b4a9cdf3711b3d6f68f069675781ca1886e216a0b88)
- [Examples](../guides/data-sources--network_customer_edge_egress--examples--group-001.md#canonical-7f463e71e192bcbb9a976ba4c7ba9f7598a671ac30aada9471885444abcfcfe0)
