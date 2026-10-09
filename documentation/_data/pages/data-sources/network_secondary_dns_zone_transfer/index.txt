---
page_title: "xcsh_network_secondary_dns_zone_transfer"
subcategory: ""
description: "Published Secondary DNS transfer and notify IPv4 addresses. The source does not distinguish their purposes. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest."
xcsh_docs: {"aliases": ["network secondary dns zone transfer"], "body_bytes": 1950, "body_sha256": "sha256:1934f645134af21c40b43d685866cbd7084e223ed5a894ee731cf434537fb93d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_secondary_dns_zone_transfer:reference", "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_secondary_dns_zone_transfer/index.md", "product": "distributed-cloud", "provider_name": "network_secondary_dns_zone_transfer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-0211000112320000-3032222023300022-3233300332112210-3202010232211322-0100212322311011-2201301123313213-3323022110111321-3110203223320302", "registry_path": "docs/data-sources/network_secondary_dns_zone_transfer.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_secondary_dns_zone_transfer/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Published Secondary DNS transfer and notify IPv4 addresses. The source does not distinguish their purposes. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_secondary_dns_zone_transfer

Breadcrumbs:

- xcsh_network_secondary_dns_zone_transfer

Published Secondary DNS transfer and notify IPv4 addresses. The source does not distinguish their
purposes. Values are bundled from the pinned OpenAPI release; this data source performs no network
request. Ports and traffic direction are not encoded in the manifest.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

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

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_secondary_dns_zone_transfer/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_secondary_dns_zone_transfer/examples/)
