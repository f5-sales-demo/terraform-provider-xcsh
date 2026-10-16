---
page_title: "xcsh_discovery"
subcategory: ""
description: "Reads Discovery information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["discovery"], "body_bytes": 1286, "body_sha256": "sha256:43374571254ed53ae4318549ec7796db9693697c450158a8b88279de80a0d6fd", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:discovery:reference", "xcsh-docs:data-sources:discovery:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:discovery:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/discovery/index.md", "product": "distributed-cloud", "provider_name": "discovery", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-2201232201303113-1031332010033212-2100120133111212-3203311223032333-0233333123230010-1020131020000301-1020330003101012-3030213333233222", "registry_path": "docs/data-sources/discovery.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/discovery/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Reads Discovery information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["discoveryCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_discovery

Breadcrumbs:

- xcsh_discovery

Reads Discovery information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Discovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Discovery by name
data "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}

output "discovery_id" {
  value = data.xcsh_discovery.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/discovery/examples/)
