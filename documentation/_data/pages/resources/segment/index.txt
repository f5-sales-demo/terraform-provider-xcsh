---
page_title: "xcsh_segment"
subcategory: ""
description: "Manages a Segment resource in F5 Distributed Cloud for segment. configuration."
xcsh_docs: {"aliases": ["segment"], "body_bytes": 1477, "body_sha256": "sha256:5478e9b108c911fedd4ac244367409c7e1a85ae969a8eb5d30339d434aa1f294", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:segment:reference", "xcsh-docs:resources:segment:examples", "xcsh-docs:resources:segment:import", "xcsh-docs:resources:segment:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:segment:collection", "completeness": "complete", "id": "xcsh-docs:resources:segment:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/segment/index.md", "product": "distributed-cloud", "provider_name": "segment", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-0101112303102030-3002312202122100-1330003103220202-0000333232013010-1103110310200110-1220311303213310-0313003331122333-2030222122100100", "registry_path": "docs/resources/segment.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/segment/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Manages a Segment resource in F5 Distributed Cloud for segment. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["segmentCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_segment

Breadcrumbs:

- xcsh_segment

Manages a Segment resource in F5 Distributed Cloud for segment. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Segment Resource Example
# Manages a Segment resource in F5 Distributed Cloud for segment.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Segment configuration
resource "xcsh_segment" "example" {
  name      = "example-segment"
  namespace = "system"
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/segment/lifecycle/timeouts/)
