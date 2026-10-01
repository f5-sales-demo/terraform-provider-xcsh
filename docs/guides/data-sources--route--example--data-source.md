---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 984, "body_sha256": "sha256:a5486e01ed8710b30c84fd81f0cfe17fac8b2ea661d566a6af1da72913bd8d21", "canonical_id": "xcsh-docs:data-sources:route:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:route:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:2077e687d5b6480f640ada5b5b9db38081cc538c7b97d40a3dfe626debd65a3f", "source_path": "examples/data-sources/xcsh_route/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:route:example:data-source", "parent_id": "xcsh-docs:data-sources:route:examples", "path": "docs/guides/data-sources--route--example--data-source.md", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/route/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_route](../data-sources/route.md)
- [Examples](data-sources--route--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_route/data-source.tf`; digest `sha256:2077e687d5b6480f640ada5b5b9db38081cc538c7b97d40a3dfe626debd65a3f`.

```terraform
# Route Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Route by name
data "xcsh_route" "example" {
  name      = "example-route"
  namespace = "staging"
}

output "route_id" {
  value = data.xcsh_route.example.id
}
```

## Next pages

- [Examples](data-sources--route--examples.md)
- [xcsh_route](../data-sources/route.md)
