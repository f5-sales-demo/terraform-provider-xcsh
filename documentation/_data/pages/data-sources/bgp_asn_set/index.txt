---
page_title: "xcsh_bgp_asn_set"
subcategory: ""
description: "Reads BGP Asn Set information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["bgp asn set"], "body_bytes": 1304, "body_sha256": "sha256:46cbdfa5024b666693cd52a30b654f88aa7b0b725f52ae9c638133a6ac06ec57", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:bgp_asn_set:reference", "xcsh-docs:data-sources:bgp_asn_set:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bgp_asn_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bgp_asn_set:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/bgp_asn_set/index.md", "product": "distributed-cloud", "provider_name": "bgp_asn_set", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-0101132211310203-2233232230121033-0201030033302303-2112033202323033-1200310311200222-2223131321000021-3211120032230011-1102013013320333", "registry_path": "docs/data-sources/bgp_asn_set.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bgp_asn_set/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Reads BGP Asn Set information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["bgp_asn_setCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bgp_asn_set

Breadcrumbs:

- xcsh_bgp_asn_set

Reads BGP Asn Set information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGPAsnSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BGPAsnSet by name
data "xcsh_bgp_asn_set" "example" {
  name      = "example-bgp-asn-set"
  namespace = "staging"
}

output "bgp_asn_set_id" {
  value = data.xcsh_bgp_asn_set.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_asn_set/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bgp_asn_set/examples/)
