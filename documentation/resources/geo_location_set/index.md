---
page_title: "xcsh_geo_location_set"
subcategory: ""
description: "Manages Geolocation Set in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["geo location set"], "body_bytes": 1505, "body_sha256": "sha256:82c55d2f759365062e3c24ae5c6ee23d22d3331e3b01bc36c225eee94590becb", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:geo_location_set:reference", "xcsh-docs:resources:geo_location_set:examples", "xcsh-docs:resources:geo_location_set:import", "xcsh-docs:resources:geo_location_set:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:geo_location_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:geo_location_set:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/geo_location_set/index.md", "product": "distributed-cloud", "provider_name": "geo_location_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-3103011222330300-2233200322023210-0200023031210213-3231332213130232-1013021130033033-2013330101030012-3301030030300101-3321313130311100", "registry_path": "docs/resources/geo_location_set.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/geo_location_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages Geolocation Set in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["geo_location_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/lifecycle/timeouts/)
