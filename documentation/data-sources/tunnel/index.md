---
page_title: "xcsh_tunnel"
subcategory: ""
description: "Manages tunnel in a given namespace. If one already exist it will give a error in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["tunnel"], "body_bytes": 1292, "body_sha256": "sha256:e5549fe53672520e17b316478122f47d5136c302d8b6662999e9536a971f849f", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:tunnel:reference", "xcsh-docs:data-sources:tunnel:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tunnel:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/tunnel/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1022022122310010-0210230002101123-2000320113321202-1120011233123033-0103110000030232-0232323131002222-1030121030120312-3302012132002001", "registry_path": "docs/data-sources/tunnel.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tunnel/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages tunnel in a given namespace. If one already exist it will give a error in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tunnel/examples/)
