---
page_title: "xcsh_access_active_sessions"
subcategory: ""
description: "Resource retrieval operation."
xcsh_docs: {"aliases": ["access active sessions"], "body_bytes": 1275, "body_sha256": "sha256:0b5874fd6417d95bd7a39756454baa94fce214c510a8771d5878970ca48230b0", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:access_active_sessions:reference", "xcsh-docs:data-sources:access_active_sessions:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:access_active_sessions:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:access_active_sessions:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/access_active_sessions/index.md", "product": "distributed-cloud", "provider_name": "access_active_sessions", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3120002103123020-1021131301111100-3330103332001301-2030202210211031-3012203032110100-2013230030123100-0331113020331110-1012323022312133", "registry_path": "docs/data-sources/access_active_sessions.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/access_active_sessions/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Resource retrieval operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_access_active_sessions

Breadcrumbs:

- xcsh_access_active_sessions

Resource retrieval operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AccessActiveSessions DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_access_active_sessions" "example" {
  namespace = "example-value"
}

output "access_active_sessions_result" {
  value = data.xcsh_access_active_sessions.example
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/examples/)
