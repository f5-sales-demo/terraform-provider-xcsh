---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_app_api_group."
xcsh_docs: {"aliases": [], "body_bytes": 1084, "body_sha256": "sha256:754e1d132a6d133b92eb995750c487e9df2a8d0664773072d375337ec0223ba2", "canonical_id": "xcsh-docs:data-sources:app_api_group:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:app_api_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ea62726df10ead3d5992d45e3c012bb70708bc94e4271b50bd5cb8013ce78312", "source_path": "examples/data-sources/xcsh_app_api_group/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:app_api_group:example:data-source", "parent_id": "xcsh-docs:data-sources:app_api_group:examples", "path": "docs/guides/data-sources--app_api_group--example--data-source.md", "provider_name": "app_api_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_api_group/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_app_api_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_app_api_group](../data-sources/app_api_group.md)
- [Examples](data-sources--app_api_group--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_app_api_group/data-source.tf`; digest `sha256:ea62726df10ead3d5992d45e3c012bb70708bc94e4271b50bd5cb8013ce78312`.

```terraform
# AppAPIGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppAPIGroup by name
data "xcsh_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "staging"
}

output "app_api_group_id" {
  value = data.xcsh_app_api_group.example.id
}
```

## Next pages

- [Examples](data-sources--app_api_group--examples.md)
- [xcsh_app_api_group](../data-sources/app_api_group.md)
