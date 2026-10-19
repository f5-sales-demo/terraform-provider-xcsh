---
page_title: "Data source"
subcategory: "Infrastructure"
description: "Data source for xcsh_site_mesh_group."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1079, "body_sha256": "sha256:014094bd21b8848b950edc9db0c314cb8a88185d1cca578a67a132325e52f20a", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site_mesh_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:57e46715034cefba0a395123af3987e2b164522a8e7ac450c2dc1b9924e5f440", "source_path": "examples/data-sources/xcsh_site_mesh_group/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_mesh_group:example:data-source", "parent_id": "xcsh-docs:data-sources:site_mesh_group:examples", "path": "documentation/data-sources/site_mesh_group/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-2030332011320001-0321202110110120-2212212021302132-1021131111021000-1202322131123031-2222333313312121-1023310021330003-3230031010103301", "registry_path": "docs/guides/data-sources--site_mesh_group--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_mesh_group/examples/data-source/index.txt", "spec_pin_digest": "sha256:ea0d80973b5efe05eb3575be6beda9c96b86701d264dc1aee3972e6a23ebb3fb", "summary": "Data source for xcsh_site_mesh_group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.3", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "6ee1e85b2e917d58078f38a270b790077b3fba12"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_mesh_group/data-source.tf`; digest `sha256:57e46715034cefba0a395123af3987e2b164522a8e7ac450c2dc1b9924e5f440`.

```terraform
# SiteMeshGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SiteMeshGroup by name
data "xcsh_site_mesh_group" "example" {
  name      = "example-site-mesh-group"
  namespace = "staging"
}

output "site_mesh_group_id" {
  value = data.xcsh_site_mesh_group.example.id
}
```
