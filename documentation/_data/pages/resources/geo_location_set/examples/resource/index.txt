---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_geo_location_set."
xcsh_docs: {"aliases": [], "body_bytes": 1266, "body_sha256": "sha256:868128de4ff8e9527ac1245e27e036b23285ed138f715efb95b7f7b08989bbaf", "child_ids": [], "collection_id": "xcsh-docs:resources:geo_location_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0e7cc4dbb2f138697448312f9d3840d2e253e947097376d3ac1a000ac5c54714", "source_path": "examples/resources/xcsh_geo_location_set/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:geo_location_set:example:resource", "parent_id": "xcsh-docs:resources:geo_location_set:examples", "path": "documentation/resources/geo_location_set/examples/resource/index.md", "provider_name": "geo_location_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/geo_location_set/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_geo_location_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["geo_location_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_geo_location_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_geo_location_set/resource.tf`; digest `sha256:0e7cc4dbb2f138697448312f9d3840d2e253e947097376d3ac1a000ac5c54714`.

```terraform
# GeoLocationSet Resource Example
# Manages Geolocation Set in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic GeoLocationSet configuration
resource "xcsh_geo_location_set" "example" {
  name      = "example-geo-location-set"
  namespace = "system"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/examples/)
- [xcsh_geo_location_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/)
