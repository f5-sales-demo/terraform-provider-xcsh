---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_geo_location_set."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1032, "body_sha256": "sha256:51e4c8dbdb1b03542ba890b47633eb697e0ff90e9680fddf0ebe00f65739dd61", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:geo_location_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0e7cc4dbb2f138697448312f9d3840d2e253e947097376d3ac1a000ac5c54714", "source_path": "examples/resources/xcsh_geo_location_set/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:geo_location_set:example:resource", "parent_id": "xcsh-docs:resources:geo_location_set:examples", "path": "documentation/resources/geo_location_set/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "geo_location_set", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0330213121232311-1231200102323110-2202312321201323-3121322002131301-2333220120310312-0213030330132220-3230200133202333-1333002123232002", "registry_path": "docs/guides/resources--geo_location_set--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/geo_location_set/examples/resource/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Resource for xcsh_geo_location_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["geo_location_setCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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
