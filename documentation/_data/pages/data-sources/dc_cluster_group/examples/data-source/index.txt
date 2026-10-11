---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_dc_cluster_group."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1088, "body_sha256": "sha256:5d58b31f419a482f2c5e05ed6ef8986b6275962fae072af1c9dce66dd34e2062", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:dc_cluster_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:82b040ec9e5742587a1d4d2feb953ac78e9dca3a538c42aff66a3020a25fe4c5", "source_path": "examples/data-sources/xcsh_dc_cluster_group/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:dc_cluster_group:example:data-source", "parent_id": "xcsh-docs:data-sources:dc_cluster_group:examples", "path": "documentation/data-sources/dc_cluster_group/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "dc_cluster_group", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-2233313011020010-1321013321103112-0202103311312302-2101102021231212-3111020232003032-3012321120300301-2201311000111311-3302202102010002", "registry_path": "docs/guides/data-sources--dc_cluster_group--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dc_cluster_group/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_dc_cluster_group.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["dc_cluster_groupCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_dc_cluster_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dc_cluster_group/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dc_cluster_group/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dc_cluster_group/data-source.tf`; digest `sha256:82b040ec9e5742587a1d4d2feb953ac78e9dca3a538c42aff66a3020a25fe4c5`.

```terraform
# DcClusterGroup Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DcClusterGroup by name
data "xcsh_dc_cluster_group" "example" {
  name      = "example-dc-cluster-group"
  namespace = "system"
}

output "dc_cluster_group_id" {
  value = data.xcsh_dc_cluster_group.example.id
}
```
