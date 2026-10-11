---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_data_group."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1031, "body_sha256": "sha256:9c39276353cf75448b86cbae9e560b6e181fa33fcac53065d3667c57619284fa", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:data_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:b9743d4eb532720e8a79b190fba83b8f94d4cbca361dec928076b2b257a940ad", "source_path": "examples/data-sources/xcsh_data_group/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:data_group:example:data-source", "parent_id": "xcsh-docs:data-sources:data_group:examples", "path": "documentation/data-sources/data_group/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "data_group", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-1120210212232021-2132232322312100-2212111023221221-3033222233313212-2020321132302010-0022221122012023-3033003011122303-3030101222330002", "registry_path": "docs/guides/data-sources--data_group--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_group/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_data_group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["data_groupCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_group/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_group/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_data_group/data-source.tf`; digest `sha256:b9743d4eb532720e8a79b190fba83b8f94d4cbca361dec928076b2b257a940ad`.

```terraform
# DataGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DataGroup by name
data "xcsh_data_group" "example" {
  name      = "example-data-group"
  namespace = "staging"
}

output "data_group_id" {
  value = data.xcsh_data_group.example.id
}
```
