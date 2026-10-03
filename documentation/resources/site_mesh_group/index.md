---
page_title: "xcsh_site_mesh_group"
subcategory: "Infrastructure"
description: "Manages Site Mesh Group in system namespace of user in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["site mesh group"], "body_bytes": 1678, "body_sha256": "sha256:7da54125aa2d023edc05667429e974c8a2f4aadadc25f04ed97c7647ff87cdbd", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:site_mesh_group:reference", "xcsh-docs:resources:site_mesh_group:examples", "xcsh-docs:resources:site_mesh_group:import", "xcsh-docs:resources:site_mesh_group:timeouts"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/site_mesh_group/index.md", "product": "distributed-cloud", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1110101130122031-1113132001032222-1100113001200302-1121013020212322-2003100312321332-2303302030212320-0302000323201302-1200320112202313", "registry_path": "docs/resources/site_mesh_group.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages Site Mesh Group in system namespace of user in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [{"name": "site", "source": "receipt-pinned-dependency:required"}], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site_mesh_group

Breadcrumbs:

- xcsh_site_mesh_group

Manages Site Mesh Group in system namespace of user in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `site`.

- site: Sites to include in mesh connectivity

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteMeshGroup Resource Example
# Manages Site Mesh Group in system namespace of user in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SiteMeshGroup configuration
resource "xcsh_site_mesh_group" "example" {
  name      = "example-site-mesh-group"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/lifecycle/timeouts/)
