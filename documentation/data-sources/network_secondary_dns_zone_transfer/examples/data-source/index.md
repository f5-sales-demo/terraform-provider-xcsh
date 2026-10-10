---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_secondary_dns_zone_transfer."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1490, "body_sha256": "sha256:b57ee2527bf642b222cdf580c95a4dafd1933a5d92e6f7d8d8478d25ef06c19f", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:da49a1e4ce345336b0b68a09a44132872466f1849e8370d3d88d79c58a018362", "source_path": "examples/data-sources/xcsh_network_secondary_dns_zone_transfer/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:example:data-source", "parent_id": "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:examples", "path": "documentation/data-sources/network_secondary_dns_zone_transfer/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_secondary_dns_zone_transfer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2310232330102123-1301221133333330-3130113312200000-2230201220022322-0012113313220132-0230202113121010-2102133212132013-3023221133023301", "registry_path": "docs/guides/data-sources--network_secondary_dns_zone_transfer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_secondary_dns_zone_transfer/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_network_secondary_dns_zone_transfer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_secondary_dns_zone_transfer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_secondary_dns_zone_transfer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_secondary_dns_zone_transfer/examples/)
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
