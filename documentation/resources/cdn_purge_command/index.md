---
page_title: "xcsh_cdn_purge_command"
subcategory: ""
description: "Manages a CDN Purge Command resource in F5 Distributed Cloud for cdn purge command specification. configuration."
xcsh_docs: {"aliases": ["cdn purge command"], "body_bytes": 1642, "body_sha256": "sha256:4d56dcd60b9acb1dd212579ef0775be69b2b28d80c15451b45636c07f1af4e52", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_purge_command:reference", "xcsh-docs:resources:cdn_purge_command:examples", "xcsh-docs:resources:cdn_purge_command:import", "xcsh-docs:resources:cdn_purge_command:timeouts"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:cdn_purge_command:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_purge_command:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/cdn_purge_command/index.md", "product": "distributed-cloud", "provider_name": "cdn_purge_command", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0023022220111331-3102221003101103-0013330220230321-0113210333131330-1333010132202103-2320102103332311-0323220121121123-1132022333100211", "registry_path": "docs/resources/cdn_purge_command.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_purge_command/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Manages a CDN Purge Command resource in F5 Distributed Cloud for cdn purge command specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["cdn_purge_commandCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_cdn_purge_command

Breadcrumbs:

- xcsh_cdn_purge_command

Manages a CDN Purge Command resource in F5 Distributed Cloud for cdn purge command specification.
configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CDNPurgeCommand Resource Example
# Manages a CDN Purge Command resource in F5 Distributed Cloud for cdn purge command specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNPurgeCommand configuration
resource "xcsh_cdn_purge_command" "example" {
  name      = "example-cdn-purge-command"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_purge_command/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_purge_command/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_purge_command/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_purge_command/lifecycle/timeouts/)
