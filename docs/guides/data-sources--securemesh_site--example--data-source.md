---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1112, "body_sha256": "sha256:ad56e33ddcdffb7f6eaf5f193f09dadfecb6a9d5e14aa305c20026be5b535d42", "canonical_id": "xcsh-docs:data-sources:securemesh_site:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:476e0c14fbf185d65694b51a11030a2462b91f28a9d067a59c797b6328e138a4", "source_path": "examples/data-sources/xcsh_securemesh_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:securemesh_site:example:data-source", "parent_id": "xcsh-docs:data-sources:securemesh_site:examples", "path": "docs/guides/data-sources--securemesh_site--example--data-source.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_securemesh_site](../data-sources/securemesh_site.md)
- [Examples](data-sources--securemesh_site--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_securemesh_site/data-source.tf`; digest `sha256:476e0c14fbf185d65694b51a11030a2462b91f28a9d067a59c797b6328e138a4`.

```terraform
# SecuremeshSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecuremeshSite by name
data "xcsh_securemesh_site" "example" {
  name      = "example-securemesh-site"
  namespace = "staging"
}

output "securemesh_site_id" {
  value = data.xcsh_securemesh_site.example.id
}
```

## Next pages

- [Examples](data-sources--securemesh_site--examples.md)
- [xcsh_securemesh_site](../data-sources/securemesh_site.md)
