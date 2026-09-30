---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_fast_acl."
xcsh_docs: {"aliases": [], "body_bytes": 1127, "body_sha256": "sha256:886187ae587cfd0fde2a0a55c5a34c66698448e7793d6241b4005b8079226e57", "child_ids": [], "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8fad5fdbb88c4dc30f2314eb3a4717e4d1278f83525eb02233ddf9ff899b963a", "source_path": "examples/data-sources/xcsh_fast_acl/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:fast_acl:example:data-source", "parent_id": "xcsh-docs:data-sources:fast_acl:examples", "path": "documentation/data-sources/fast_acl/examples/data-source/index.md", "provider_name": "fast_acl", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_fast_acl.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fast_aclCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_fast_acl/data-source.tf`; digest `sha256:8fad5fdbb88c4dc30f2314eb3a4717e4d1278f83525eb02233ddf9ff899b963a`.

```terraform
# FastACL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FastACL by name
data "xcsh_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"
}

output "fast_acl_id" {
  value = data.xcsh_fast_acl.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/examples/)
- [xcsh_fast_acl](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/)
