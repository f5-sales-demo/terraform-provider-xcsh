---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1108, "body_sha256": "sha256:c9b5acdb021fb7b5dbd85795182e9eec5dbd2571cbc649fa3d5edfdf06e2fb8b", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:2abaac1742f086ab988ef94e8d836fe6f538c549072b6b6dd846213e69e59239", "source_path": "examples/data-sources/xcsh_securemesh_site_v2/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:securemesh_site_v2:example:data-source", "parent_id": "xcsh-docs:data-sources:securemesh_site_v2:examples", "path": "documentation/data-sources/securemesh_site_v2/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2320121111203001-0201203100230213-1333311001213133-0333212022121321-0231200132112320-3310031302030330-3321100321022321-2222131112131300", "registry_path": "docs/guides/data-sources--securemesh_site_v2--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_securemesh_site_v2.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_securemesh_site_v2](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site_v2/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_securemesh_site_v2/data-source.tf`; digest `sha256:2abaac1742f086ab988ef94e8d836fe6f538c549072b6b6dd846213e69e59239`.

```terraform
# SecuremeshSiteV2 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecuremeshSiteV2 by name
data "xcsh_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2"
  namespace = "system"
}

output "securemesh_site_v2_id" {
  value = data.xcsh_securemesh_site_v2.example.id
}
```
