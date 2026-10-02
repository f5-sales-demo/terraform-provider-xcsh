---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_geo_location_set."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1328, "body_sha256": "sha256:ff4e9fcf40154ff4fba45cb067ad9c82d7cdd7641052750aa0a60d5ed4f0bbbc", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:geo_location_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fe5ac2dcce0034284edb5717f6d3da18db76cd12233f57cc1a3e9be26885c3ea", "source_path": "examples/data-sources/xcsh_geo_location_set/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:geo_location_set:example:data-source", "parent_id": "xcsh-docs:data-sources:geo_location_set:examples", "path": "documentation/data-sources/geo_location_set/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "geo_location_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0230123203110203-1010212223323020-0222020203013212-0311132300332123-1011213212232302-2130110222302011-0010223032211111-1121123111310321", "registry_path": "docs/guides/data-sources--geo_location_set--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/geo_location_set/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_geo_location_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["geo_location_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
