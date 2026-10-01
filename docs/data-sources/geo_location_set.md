---
page_title: "xcsh_geo_location_set"
subcategory: ""
description: "xcsh_geo_location_set for xcsh_geo_location_set."
xcsh_docs: {"aliases": [], "body_bytes": 1234, "body_sha256": "sha256:c55d9eb9fae48cb2790eae05bedcf50b45400e0bb7ebbb2e2507ed2e8e974b28", "canonical_id": "xcsh-docs:data-sources:geo_location_set:fundamentals", "child_ids": ["xcsh-docs:data-sources:geo_location_set:reference", "xcsh-docs:data-sources:geo_location_set:examples"], "collection_id": "xcsh-docs:data-sources:geo_location_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:geo_location_set:fundamentals", "parent_id": null, "path": "docs/data-sources/geo_location_set.md", "provider_name": "geo_location_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/geo_location_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_geo_location_set for xcsh_geo_location_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["geo_location_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_geo_location_set

Breadcrumbs:

- xcsh_geo_location_set

Manages Geolocation Set in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--geo_location_set--reference.md)
- [Examples](../guides/data-sources--geo_location_set--examples.md)
