---
page_title: "xcsh_network_secondary_dns_zone_transfer landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_secondary_dns_zone_transfer landing."
---

# xcsh_network_secondary_dns_zone_transfer landing

<a id="canonical-25016e00cea8bc0aefc3e5a4e212e97a109bad45a1c5bde7fb294579d48ebe32"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c21df305415f23e8cc7614235d0351e94803cb8e940a0f75eb45097580c372c8"></a>

## xcsh_network_secondary_dns_zone_transfer — xcsh_network_secondary_dns_zone_transfer / 750038b75e6a / 2

Breadcrumbs:

- xcsh_network_secondary_dns_zone_transfer

Published Secondary DNS transfer and notify IPv4 addresses. The source does not distinguish their
purposes. Values are bundled from the pinned OpenAPI release; this data source performs no network
request. Ports and traffic direction are not encoded in the manifest.

<a id="canonical-ced39f5b6043367bd14be04ed3b18a0e993bcdf23dedeca3853aaa1c31ece807"></a>

## Prerequisites — xcsh_network_secondary_dns_zone_transfer / 750038b75e6a / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-5aa8533bc7db5d2f96c2bb2e456d34bd6fbeace3af1eec141dc03563cd6d3dde"></a>

## Minimal configuration — xcsh_network_secondary_dns_zone_transfer / 750038b75e6a / 4

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

<a id="canonical-b2101240f647a4d92c6aa6fd3f9805109db89015fb8166f01a6c3617cdbf6839"></a>

## Root configuration — xcsh_network_secondary_dns_zone_transfer / 750038b75e6a / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-3ec9376a6250475eb0cee86b3ad89743869635e796f59dfd4e54cc7cdb751579"></a>

## Next pages — xcsh_network_secondary_dns_zone_transfer / 750038b75e6a / 6

- [Property reference](../guides/data-sources--network_secondary_dns_zone_transfer--reference--group-001.md#canonical-c945fd07bc00a831ed39b45e5684e423a74eab2504bb40760e1b07243fd4e278)
- [Examples](../guides/data-sources--network_secondary_dns_zone_transfer--examples--group-001.md#canonical-bc67719214c86f2fffbd605245afc762550a9ca52160a01d1c861afb2aaf5d2f)
