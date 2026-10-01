---
page_title: "Data source"
subcategory: "Infrastructure"
description: "Data source for xcsh_site_mesh_group."
xcsh_docs: {"aliases": [], "body_bytes": 1110, "body_sha256": "sha256:4d7fa9724c3ec966b57d0883c7768b67330179fc9deeceecc3326a3c0e71a991", "canonical_id": "xcsh-docs:data-sources:site_mesh_group:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_mesh_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:57e46715034cefba0a395123af3987e2b164522a8e7ac450c2dc1b9924e5f440", "source_path": "examples/data-sources/xcsh_site_mesh_group/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_mesh_group:example:data-source", "parent_id": "xcsh-docs:data-sources:site_mesh_group:examples", "path": "docs/guides/data-sources--site_mesh_group--example--data-source.md", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_mesh_group/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_site_mesh_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md)
- [Examples](data-sources--site_mesh_group--examples.md)
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

## Next pages

- [Examples](data-sources--site_mesh_group--examples.md)
- [xcsh_site_mesh_group](../data-sources/site_mesh_group.md)
