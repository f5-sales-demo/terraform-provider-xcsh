---
page_title: "Data source"
subcategory: "Infrastructure"
description: "Data source for xcsh_site_mesh_group."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1316, "body_sha256": "sha256:a1d7111a52ef4677d25b9f60ed1e274e838a79fab90216235a32751002cf0cce", "capabilities": ["networking"], "category": "networking", "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:site_mesh_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:57e46715034cefba0a395123af3987e2b164522a8e7ac450c2dc1b9924e5f440", "source_path": "examples/data-sources/xcsh_site_mesh_group/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_mesh_group:example:data-source", "parent_id": "xcsh-docs:data-sources:site_mesh_group:examples", "path": "documentation/data-sources/site_mesh_group/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "site_mesh_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2030332011320001-0321202110110120-2212212021302132-1021131111021000-1202322131123031-2222333313312121-1023310021330003-3230031010103301", "registry_path": "docs/guides/data-sources--site_mesh_group--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_mesh_group/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_site_mesh_group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["site_mesh_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/examples/)
- [xcsh_site_mesh_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_mesh_group/)
