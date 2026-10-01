---
page_title: "xcsh_site_mesh_group"
subcategory: "Infrastructure"
description: "xcsh_site_mesh_group for xcsh_site_mesh_group."
xcsh_docs: {"aliases": [], "body_bytes": 1489, "body_sha256": "sha256:ac8252ad122fcdf99640a1162c6f72278c23dbc6089e5c9670cd54d50ffe5062", "canonical_id": "xcsh-docs:resources:site_mesh_group:fundamentals", "child_ids": ["xcsh-docs:resources:site_mesh_group:reference", "xcsh-docs:resources:site_mesh_group:examples", "xcsh-docs:resources:site_mesh_group:import", "xcsh-docs:resources:site_mesh_group:timeouts"], "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:site_mesh_group:fundamentals", "parent_id": null, "path": "docs/resources/site_mesh_group.md", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_site_mesh_group for xcsh_site_mesh_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [Property reference](../guides/resources--site_mesh_group--reference.md)
- [Examples](../guides/resources--site_mesh_group--examples.md)
- [Import](../guides/resources--site_mesh_group--import.md)
- [Timeouts](../guides/resources--site_mesh_group--timeouts.md)
