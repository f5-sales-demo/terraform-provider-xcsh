---
page_title: "xcsh_geo_location_set"
subcategory: ""
description: "Manages Geolocation Set in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["geo location set"], "body_bytes": 1518, "body_sha256": "sha256:939aa66bcedf44c5e36fccb75c7269568098117315419109a3fe9aef199e9fcd", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:geo_location_set:reference", "xcsh-docs:resources:geo_location_set:examples", "xcsh-docs:resources:geo_location_set:import", "xcsh-docs:resources:geo_location_set:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:geo_location_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:geo_location_set:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/geo_location_set/index.md", "product": "distributed-cloud", "provider_name": "geo_location_set", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-3103011222330300-2233200322023210-0200023031210213-3231332213130232-1013021130033033-2013330101030012-3301030030300101-3321313130311100", "registry_path": "docs/resources/geo_location_set.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/geo_location_set/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Manages Geolocation Set in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["geo_location_setCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/geo_location_set/lifecycle/timeouts/)
