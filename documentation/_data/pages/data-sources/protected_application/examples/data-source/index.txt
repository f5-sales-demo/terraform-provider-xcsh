---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 1396, "body_sha256": "sha256:eeb50f1bcc9493fa560bdd7e3d47198ef8da2ca1b8735cb364b7720f8ac35e92", "child_ids": [], "collection_id": "xcsh-docs:data-sources:protected_application:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a72cc964b57acaba96990190e74ffa0a1d2b62d15327b6423caa90a5e44e936b", "source_path": "examples/data-sources/xcsh_protected_application/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:protected_application:example:data-source", "parent_id": "xcsh-docs:data-sources:protected_application:examples", "path": "documentation/data-sources/protected_application/examples/data-source/index.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_application/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_protected_application/data-source.tf`; digest `sha256:a72cc964b57acaba96990190e74ffa0a1d2b62d15327b6423caa90a5e44e936b`.

```terraform
# ProtectedApplication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtectedApplication by name
data "xcsh_protected_application" "example" {
  name      = "example-protected-application"
  namespace = "staging"
}

output "protected_application_id" {
  value = data.xcsh_protected_application.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/examples/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/protected_application/)
