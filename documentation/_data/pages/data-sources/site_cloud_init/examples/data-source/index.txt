---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_cloud_init."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1061, "body_sha256": "sha256:15be87b44c0ab30449327a55d063614a3f675c78415bebde9a258fdd221edb52", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_cloud_init:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:edce0cae6fa8845014088889818debc08fd6df550bae0c288c669538441379ae", "source_path": "examples/data-sources/xcsh_site_cloud_init/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_cloud_init:example:data-source", "parent_id": "xcsh-docs:data-sources:site_cloud_init:examples", "path": "documentation/data-sources/site_cloud_init/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "site_cloud_init", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3023031122223133-3022120121113110-2321333301133030-2300003101220121-2011303302310123-1320021012330331-3211001111311321-1210212212231330", "registry_path": "docs/guides/data-sources--site_cloud_init--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_cloud_init/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_site_cloud_init.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_cloud_init](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_cloud_init/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_cloud_init/data-source.tf`; digest `sha256:edce0cae6fa8845014088889818debc08fd6df550bae0c288c669538441379ae`.

```terraform
# SiteCloudInit DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_cloud_init" "example" {
  provider_ref = "example-value"
  site_name    = "example-value"
}

output "site_cloud_init_result" {
  value     = data.xcsh_site_cloud_init.example
  sensitive = true
}
```
