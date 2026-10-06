---
page_title: "xcsh_dns_zone_add_cryptokey"
subcategory: ""
description: "Adds a cryptographic key to a DNS zone."
xcsh_docs: {"aliases": ["dns zone add cryptokey"], "body_bytes": 1281, "body_sha256": "sha256:c5ee061f640cd8edac05476cb0361c25f7a6675f796e0eb8aa1421b5dfbb1495", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:dns_zone_add_cryptokey:reference", "xcsh-docs:actions:dns_zone_add_cryptokey:examples", "xcsh-docs:actions:dns_zone_add_cryptokey:lifecycle"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:dns_zone_add_cryptokey:collection", "completeness": "complete", "id": "xcsh-docs:actions:dns_zone_add_cryptokey:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/dns_zone_add_cryptokey/index.md", "product": "distributed-cloud", "provider_name": "dns_zone_add_cryptokey", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "actions", "registry_anchor": "canonical-2011333313231130-3200122021220331-0010211322333131-1302203232313332-2231111130321232-3311022020312311-2322131130200322-0032011221212010", "registry_path": "docs/actions/dns_zone_add_cryptokey.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/dns_zone_add_cryptokey/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Adds a cryptographic key to a DNS zone.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_dns_zone_add_cryptokey

Breadcrumbs:

- xcsh_dns_zone_add_cryptokey

Adds a cryptographic key to a DNS zone.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSZoneAddCryptokey Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_dns_zone_add_cryptokey" "example" {
  config {
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_add_cryptokey/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_add_cryptokey/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/dns_zone_add_cryptokey/lifecycle/)
