---
page_title: "xcsh_tunnel"
subcategory: ""
description: "Manages tunnel in a given namespace. If one already exist it will give a error in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["tunnel"], "body_bytes": 1479, "body_sha256": "sha256:eca50c50f3eceb265983dfd91d1374fc8e618db9699ef6726e36795e65931e2f", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:reference", "xcsh-docs:resources:tunnel:examples", "xcsh-docs:resources:tunnel:import", "xcsh-docs:resources:tunnel:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/tunnel/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122", "registry_path": "docs/resources/tunnel.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Manages tunnel in a given namespace. If one already exist it will give a error in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["tunnelCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_tunnel

Breadcrumbs:

- xcsh_tunnel

Manages tunnel in a given namespace. If one already exist it will give a error in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Tunnel Resource Example
# Manages tunnel in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Tunnel configuration
resource "xcsh_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/lifecycle/timeouts/)
