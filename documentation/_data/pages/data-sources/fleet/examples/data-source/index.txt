---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_fleet."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 983, "body_sha256": "sha256:fd047f649c0bafaba950943584c55e0d37d3c67de0b3277fe1bde3a6c02e9d8d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fleet:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6d30c4d1a598a1f7f950e3ba739fc0d0cef6e7cc4d4b114700efeb50ec416c17", "source_path": "examples/data-sources/xcsh_fleet/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:fleet:example:data-source", "parent_id": "xcsh-docs:data-sources:fleet:examples", "path": "documentation/data-sources/fleet/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1323010013013231-3003101211233233-3010120120301000-0230121132123333-1331011333002031-1211102213323202-2302012333020033-2033013320121233", "registry_path": "docs/guides/data-sources--fleet--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fleet/examples/data-source/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Data source for xcsh_fleet.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["fleetCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
