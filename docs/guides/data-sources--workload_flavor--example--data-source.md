---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_workload_flavor."
xcsh_docs: {"aliases": [], "body_bytes": 1111, "body_sha256": "sha256:bae22e087512ca68c1b9800966f21716ae68d2d3257aa5d840559f3382a52eb9", "canonical_id": "xcsh-docs:data-sources:workload_flavor:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:workload_flavor:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6a1b8a9b6022de8c3cc75f425f01b978df5bf741969fc4bc0998b22eb6f2b173", "source_path": "examples/data-sources/xcsh_workload_flavor/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:workload_flavor:example:data-source", "parent_id": "xcsh-docs:data-sources:workload_flavor:examples", "path": "docs/guides/data-sources--workload_flavor--example--data-source.md", "provider_name": "workload_flavor", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/workload_flavor/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_workload_flavor.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["workload_flavorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_workload_flavor](../data-sources/workload_flavor.md)
- [Examples](data-sources--workload_flavor--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_workload_flavor/data-source.tf`; digest `sha256:6a1b8a9b6022de8c3cc75f425f01b978df5bf741969fc4bc0998b22eb6f2b173`.

```terraform
# WorkloadFlavor Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing WorkloadFlavor by name
data "xcsh_workload_flavor" "example" {
  name      = "example-workload-flavor"
  namespace = "shared"
}

output "workload_flavor_id" {
  value = data.xcsh_workload_flavor.example.id
}
```

## Next pages

- [Examples](data-sources--workload_flavor--examples.md)
- [xcsh_workload_flavor](../data-sources/workload_flavor.md)
