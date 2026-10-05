---
page_title: "Resource"
subcategory: "Infrastructure"
description: "Resource for xcsh_site_mesh_group."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1284, "body_sha256": "sha256:7d7b3b1599faa0a5e8a5e10a07a28cf891b4b4d3db98dc184894028f63539fb4", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:site_mesh_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9fd9c5f06a1b22b1d6bf2d57189b47e2ae5e2ce09701525907fff4345875b45e", "source_path": "examples/resources/xcsh_site_mesh_group/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:site_mesh_group:example:resource", "parent_id": "xcsh-docs:resources:site_mesh_group:examples", "path": "documentation/resources/site_mesh_group/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2013331331320111-3100102220130032-1121301310333001-1322300330132123-2323313112222030-0123032021312123-0110130030112030-3322110313311303", "registry_path": "docs/guides/resources--site_mesh_group--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/site_mesh_group/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_site_mesh_group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_site_mesh_group/resource.tf`; digest `sha256:9fd9c5f06a1b22b1d6bf2d57189b47e2ae5e2ce09701525907fff4345875b45e`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/examples/)
- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/site_mesh_group/)
