---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_data_group."
xcsh_docs: {"aliases": [], "body_bytes": 1047, "body_sha256": "sha256:72172b22ca99b694eb3ebce461b8b93b9954f7db78807d2fa33bdedb4ab434d3", "canonical_id": "xcsh-docs:data-sources:data_group:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:data_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:b9743d4eb532720e8a79b190fba83b8f94d4cbca361dec928076b2b257a940ad", "source_path": "examples/data-sources/xcsh_data_group/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:data_group:example:data-source", "parent_id": "xcsh-docs:data-sources:data_group:examples", "path": "docs/guides/data-sources--data_group--example--data-source.md", "provider_name": "data_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_group/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_data_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_data_group](../data-sources/data_group.md)
- [Examples](data-sources--data_group--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_data_group/data-source.tf`; digest `sha256:b9743d4eb532720e8a79b190fba83b8f94d4cbca361dec928076b2b257a940ad`.

```terraform
# DataGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DataGroup by name
data "xcsh_data_group" "example" {
  name      = "example-data-group"
  namespace = "staging"
}

output "data_group_id" {
  value = data.xcsh_data_group.example.id
}
```

## Next pages

- [Examples](data-sources--data_group--examples.md)
- [xcsh_data_group](../data-sources/data_group.md)
