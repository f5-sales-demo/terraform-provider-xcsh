---
page_title: "xcsh_network_secondary_dns_zone_transfer"
subcategory: ""
description: "Published Secondary DNS transfer and notify IPv4 addresses. The source does not distinguish their purposes. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest."
xcsh_docs: {"aliases": ["network secondary dns zone transfer"], "body_bytes": 1937, "body_sha256": "sha256:73ff955e2e7efcf1ac49dd5505bce2e16b8f751ed082fdf72babe580d6f697ab", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:network_secondary_dns_zone_transfer:reference", "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:examples"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_secondary_dns_zone_transfer:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_secondary_dns_zone_transfer/index.md", "product": "distributed-cloud", "provider_name": "network_secondary_dns_zone_transfer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0211000112320000-3032222023300022-3233300332112210-3202010232211322-0100212322311011-2201301123313213-3323022110111321-3110203223320302", "registry_path": "docs/data-sources/network_secondary_dns_zone_transfer.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_secondary_dns_zone_transfer/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Published Secondary DNS transfer and notify IPv4 addresses. The source does not distinguish their purposes. Values are bundled from the pinned OpenAPI release; this data source performs no network request. Ports and traffic direction are not encoded in the manifest.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_secondary_dns_zone_transfer/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_secondary_dns_zone_transfer/examples/)
