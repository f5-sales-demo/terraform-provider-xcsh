---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_geo_location_set."
xcsh_docs: {"aliases": [], "body_bytes": 1328, "body_sha256": "sha256:ff4e9fcf40154ff4fba45cb067ad9c82d7cdd7641052750aa0a60d5ed4f0bbbc", "child_ids": [], "collection_id": "xcsh-docs:data-sources:geo_location_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fe5ac2dcce0034284edb5717f6d3da18db76cd12233f57cc1a3e9be26885c3ea", "source_path": "examples/data-sources/xcsh_geo_location_set/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:geo_location_set:example:data-source", "parent_id": "xcsh-docs:data-sources:geo_location_set:examples", "path": "documentation/data-sources/geo_location_set/examples/data-source/index.md", "provider_name": "geo_location_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/geo_location_set/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_geo_location_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["geo_location_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_geo_location_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/geo_location_set/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/geo_location_set/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_geo_location_set/data-source.tf`; digest `sha256:fe5ac2dcce0034284edb5717f6d3da18db76cd12233f57cc1a3e9be26885c3ea`.

```terraform
# GeoLocationSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing GeoLocationSet by name
data "xcsh_geo_location_set" "example" {
  name      = "example-geo-location-set"
  namespace = "system"
}

output "geo_location_set_id" {
  value = data.xcsh_geo_location_set.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/geo_location_set/examples/)
- [xcsh_geo_location_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/geo_location_set/)
