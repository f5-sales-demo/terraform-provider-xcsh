---
page_title: "xcsh_tunnel"
subcategory: ""
description: "Manages tunnel in a given namespace. If one already exist it will give a error in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["tunnel"], "body_bytes": 1466, "body_sha256": "sha256:ac63b94faf6a0f3dab7b6d32eba55712d382f0e15dfd03fac13d0c81237a2512", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:tunnel:reference", "xcsh-docs:resources:tunnel:examples", "xcsh-docs:resources:tunnel:import", "xcsh-docs:resources:tunnel:timeouts"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/tunnel/index.md", "product": "distributed-cloud", "provider_name": "tunnel", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122", "registry_path": "docs/resources/tunnel.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages tunnel in a given namespace. If one already exist it will give a error in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tunnelCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tunnel/lifecycle/timeouts/)
