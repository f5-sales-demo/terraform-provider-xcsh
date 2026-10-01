---
page_title: "xcsh_network_secondary_dns_zone_transfer examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_secondary_dns_zone_transfer examples."
---

# xcsh_network_secondary_dns_zone_transfer examples

<a id="canonical-bc67719214c86f2fffbd605245afc762550a9ca52160a01d1c861afb2aaf5d2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-bb8ec401c0e56561c14f54afd333cde62db1bfe72ef6b95e07fbf1553cbba98e"></a>

## Examples — Examples / 36c08be6b576 / 2

Breadcrumbs:

- [xcsh_network_secondary_dns_zone_transfer](../data-sources/network_secondary_dns_zone_transfer.md#canonical-25016e00cea8bc0aefc3e5a4e212e97a109bad45a1c5bde7fb294579d48ebe32)
- Examples

<a id="canonical-2b2722428a8144b1dd4418549ef59bc9056eca38c0756e4cd0f42c2c8cb869ea"></a>

## Complete configurations — Examples / 36c08be6b576 / 3

- [Data source](data-sources--network_secondary_dns_zone_transfer--examples--group-001.md#canonical-b4bbc49b71a5fffcdc5f6800ac8682ba065f7a1e2c897644927e6787cba5f2f1): valid configuration.

<a id="canonical-910d695d5b3db7a8b4a57c6a0625e9379326f146ff2110b78408dd0356cede7a"></a>

## Next pages — Examples / 36c08be6b576 / 4

- [Data source](data-sources--network_secondary_dns_zone_transfer--examples--group-001.md#canonical-b4bbc49b71a5fffcdc5f6800ac8682ba065f7a1e2c897644927e6787cba5f2f1)
- [xcsh_network_secondary_dns_zone_transfer](../data-sources/network_secondary_dns_zone_transfer.md#canonical-25016e00cea8bc0aefc3e5a4e212e97a109bad45a1c5bde7fb294579d48ebe32)

<a id="canonical-b4bbc49b71a5fffcdc5f6800ac8682ba065f7a1e2c897644927e6787cba5f2f1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-87226733109530ddb07b5eb028bd7ae69d2fce9ee9cac79df70b4f3082a13963"></a>

## Data source — Data source / e62c506ebcf5 / 2

Breadcrumbs:

- [xcsh_network_secondary_dns_zone_transfer](../data-sources/network_secondary_dns_zone_transfer.md#canonical-25016e00cea8bc0aefc3e5a4e212e97a109bad45a1c5bde7fb294579d48ebe32)
- [Examples](data-sources--network_secondary_dns_zone_transfer--examples--group-001.md#canonical-bc67719214c86f2fffbd605245afc762550a9ca52160a01d1c861afb2aaf5d2f)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_secondary_dns_zone_transfer/data-source.tf`; digest `sha256:da49a1e4ce345336b0b68a09a44132872466f1849e8370d3d88d79c58a018362`.

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

data "xcsh_network_secondary_dns_zone_transfer" "authoritative_dns" {}

# The manifest combines transfer and notify sources, so both explicit DNS
# rules use the same published allowlist.
output "secondary_dns_rules" {
  value = [
    {
      direction = "ingress"
      protocol  = "tcp"
      port      = 53
      sources   = data.xcsh_network_secondary_dns_zone_transfer.authoritative_dns.cidr_blocks
    },
    {
      direction = "ingress"
      protocol  = "udp"
      port      = 53
      sources   = data.xcsh_network_secondary_dns_zone_transfer.authoritative_dns.cidr_blocks
    },
  ]
}
```

<a id="canonical-b0a398e7451b55a68532dec672acb86a49fad1c34b9e800be200f7cca3a77c67"></a>

## Next pages — Data source / e62c506ebcf5 / 3

- [Examples](data-sources--network_secondary_dns_zone_transfer--examples--group-001.md#canonical-bc67719214c86f2fffbd605245afc762550a9ca52160a01d1c861afb2aaf5d2f)
- [xcsh_network_secondary_dns_zone_transfer](../data-sources/network_secondary_dns_zone_transfer.md#canonical-25016e00cea8bc0aefc3e5a4e212e97a109bad45a1c5bde7fb294579d48ebe32)
