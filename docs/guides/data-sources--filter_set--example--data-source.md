---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_filter_set."
xcsh_docs: {"aliases": [], "body_bytes": 948, "body_sha256": "sha256:0a8b46662b34272922c3b56842923c2b7164549664dbfabbe07db506a6894eb4", "canonical_id": "xcsh-docs:data-sources:filter_set:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:filter_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:03a6f567d57f7b6dda2ca579596922f0b9b282f98f909656ba616ef93c767b0d", "source_path": "examples/data-sources/xcsh_filter_set/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:filter_set:example:data-source", "parent_id": "xcsh-docs:data-sources:filter_set:examples", "path": "docs/guides/data-sources--filter_set--example--data-source.md", "provider_name": "filter_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/filter_set/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_filter_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["filter_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_filter_set](../data-sources/filter_set.md)
- [Examples](data-sources--filter_set--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_filter_set/data-source.tf`; digest `sha256:03a6f567d57f7b6dda2ca579596922f0b9b282f98f909656ba616ef93c767b0d`.

```terraform
# FilterSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FilterSet by name
data "xcsh_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "staging"
}

output "filter_set_id" {
  value = data.xcsh_filter_set.example.id
}
```

## Next pages

- [Examples](data-sources--filter_set--examples.md)
- [xcsh_filter_set](../data-sources/filter_set.md)
