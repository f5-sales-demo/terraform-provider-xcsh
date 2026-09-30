---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1091, "body_sha256": "sha256:e028651a39c07d1e1aec3692134ce74c2c5e23b6c8142d35a0aabb4662f1e57e", "child_ids": [], "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6d30c4d1a598a1f7f950e3ba739fc0d0cef6e7cc4d4b114700efeb50ec416c17", "source_path": "examples/data-sources/xcsh_fleet/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:fleet:example:data-source", "parent_id": "xcsh-docs:data-sources:fleet:examples", "path": "documentation/data-sources/fleet/examples/data-source/index.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_fleet/data-source.tf`; digest `sha256:6d30c4d1a598a1f7f950e3ba739fc0d0cef6e7cc4d4b114700efeb50ec416c17`.

```terraform
# Fleet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Fleet by name
data "xcsh_fleet" "example" {
  name      = "example-fleet"
  namespace = "staging"
}

output "fleet_id" {
  value = data.xcsh_fleet.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/examples/)
- [xcsh_fleet](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fleet/)
