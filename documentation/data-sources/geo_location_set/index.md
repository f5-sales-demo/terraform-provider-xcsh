---
page_title: "xcsh_geo_location_set"
subcategory: ""
description: "Manages Geolocation Set in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["geo location set"], "body_bytes": 1319, "body_sha256": "sha256:e4266991310d74a55094fd677bdbcdf6280a599b546a9f7751eb703811fe546e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:geo_location_set:reference", "xcsh-docs:data-sources:geo_location_set:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:geo_location_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:geo_location_set:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/geo_location_set/index.md", "product": "distributed-cloud", "provider_name": "geo_location_set", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1001100000212311-3233013232003113-2232222223002331-1133222232220002-1101232133012211-1121021022203120-3302330011301312-1023310223132122", "registry_path": "docs/data-sources/geo_location_set.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/geo_location_set/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages Geolocation Set in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["geo_location_setCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/geo_location_set/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/geo_location_set/examples/)
