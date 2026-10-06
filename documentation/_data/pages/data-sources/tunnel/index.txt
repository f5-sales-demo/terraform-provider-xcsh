---
page_title: "xcsh_tunnel"
subcategory: ""
description: "Reads Tunnel information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["tunnel"], "body_bytes": 1253, "body_sha256": "sha256:3c6867c66a2dd3b81e7788e6c6b517d7a9027779c096813ed4a1f92010440458", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:tunnel:reference", "xcsh-docs:data-sources:tunnel:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/tunnel/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001", "registry_path": "docs/data-sources/tunnel.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reads Tunnel information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_tunnel

Breadcrumbs:

- xcsh_tunnel

Reads Tunnel information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Tunnel Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Tunnel by name
data "xcsh_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "staging"
}

output "tunnel_id" {
  value = data.xcsh_tunnel.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/examples/)
