---
page_title: "xcsh_network_customer_edge_egress examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_customer_edge_egress examples."
---

# xcsh_network_customer_edge_egress examples

<a id="canonical-7f463e71e192bcbb9a976ba4c7ba9f7598a671ac30aada9471885444abcfcfe0"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-66150de6d58df2170be0f35970460e278c1a7d8a436147284df7429dca7f8db6"></a>

## Examples — Examples / 32af69d9d70e / 2

Breadcrumbs:

- [xcsh_network_customer_edge_egress](../data-sources/network_customer_edge_egress.md#canonical-e8a53f08d4a40bf225466083460717f25dabd979bfa6a17a3cab3b3fbaa24138)
- Examples

<a id="canonical-2743974f0d650d212f4b0c768921c309d99dddc06560303c6df55c3fb794ea62"></a>

## Complete configurations — Examples / 32af69d9d70e / 3

- [Data source](data-sources--network_customer_edge_egress--examples--group-001.md#canonical-c48507940c0c62e1bb3061c50a2a066caa5199b519652c1eb41e57c8a10217a5): valid configuration.

<a id="canonical-11781cbc700dcf8dfff134448be3f76a1837c42f361c2d5034e40380536901d7"></a>

## Next pages — Examples / 32af69d9d70e / 4

- [Data source](data-sources--network_customer_edge_egress--examples--group-001.md#canonical-c48507940c0c62e1bb3061c50a2a066caa5199b519652c1eb41e57c8a10217a5)
- [xcsh_network_customer_edge_egress](../data-sources/network_customer_edge_egress.md#canonical-e8a53f08d4a40bf225466083460717f25dabd979bfa6a17a3cab3b3fbaa24138)

<a id="canonical-c48507940c0c62e1bb3061c50a2a066caa5199b519652c1eb41e57c8a10217a5"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-7f85463cc8c32624f30b89afb374c86fa3bfcf3aac58fb2c0474c3ea939219c4"></a>

## Data source — Data source / b7d6d55c45e8 / 2

Breadcrumbs:

- [xcsh_network_customer_edge_egress](../data-sources/network_customer_edge_egress.md#canonical-e8a53f08d4a40bf225466083460717f25dabd979bfa6a17a3cab3b3fbaa24138)
- [Examples](data-sources--network_customer_edge_egress--examples--group-001.md#canonical-7f463e71e192bcbb9a976ba4c7ba9f7598a671ac30aada9471885444abcfcfe0)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_customer_edge_egress/data-source.tf`; digest `sha256:1a958054768616b5beeebffa91db22fe699af5f80112367c4fcd9c1d753d7368`.

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

<a id="canonical-762fa1d390822aea1100dc55cf7f3b841fff518f695ff657ce8537e5b068271e"></a>

## Next pages — Data source / b7d6d55c45e8 / 3

- [Examples](data-sources--network_customer_edge_egress--examples--group-001.md#canonical-7f463e71e192bcbb9a976ba4c7ba9f7598a671ac30aada9471885444abcfcfe0)
- [xcsh_network_customer_edge_egress](../data-sources/network_customer_edge_egress.md#canonical-e8a53f08d4a40bf225466083460717f25dabd979bfa6a17a3cab3b3fbaa24138)
