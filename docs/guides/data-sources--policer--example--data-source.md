---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_policer."
xcsh_docs: {"aliases": [], "body_bytes": 1010, "body_sha256": "sha256:7e834230c5a46648991883008d2fa38a2d24962e98bd1cb17c6617c7e03404b4", "canonical_id": "xcsh-docs:data-sources:policer:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:policer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:bf10d06a3b33dc40f3d87d76f2625b6d1d04a36df75db4e7eceed8ef9f19362a", "source_path": "examples/data-sources/xcsh_policer/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:policer:example:data-source", "parent_id": "xcsh-docs:data-sources:policer:examples", "path": "docs/guides/data-sources--policer--example--data-source.md", "provider_name": "policer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policer/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_policer](../data-sources/policer.md)
- [Examples](data-sources--policer--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_policer/data-source.tf`; digest `sha256:bf10d06a3b33dc40f3d87d76f2625b6d1d04a36df75db4e7eceed8ef9f19362a`.

```terraform
# Policer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Policer by name
data "xcsh_policer" "example" {
  name      = "example-policer"
  namespace = "staging"
}

output "policer_id" {
  value = data.xcsh_policer.example.id
}
```

## Next pages

- [Examples](data-sources--policer--examples.md)
- [xcsh_policer](../data-sources/policer.md)
