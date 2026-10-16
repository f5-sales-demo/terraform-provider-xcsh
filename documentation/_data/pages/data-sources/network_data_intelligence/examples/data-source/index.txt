---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_data_intelligence."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1126, "body_sha256": "sha256:2e75113647eee5a5543a2c32ca0ea2d54f2fab54e757f4a02d3efb70b1eed00a", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_data_intelligence:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:529d85fb608129abcf053afa1a4c7516767c8567633cbc06efe8d845d135d233", "source_path": "examples/data-sources/xcsh_network_data_intelligence/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_data_intelligence:example:data-source", "parent_id": "xcsh-docs:data-sources:network_data_intelligence:examples", "path": "documentation/data-sources/network_data_intelligence/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_data_intelligence", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-2310133133233022-2132112222002311-0131322010333011-3133102122000021-1223111121232020-3223130120320103-2312031313322200-1322132211030230", "registry_path": "docs/guides/data-sources--network_data_intelligence--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_data_intelligence/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Data source for xcsh_network_data_intelligence.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": [], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_data_intelligence](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_data_intelligence/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_data_intelligence/data-source.tf`; digest `sha256:529d85fb608129abcf053afa1a4c7516767c8567633cbc06efe8d845d135d233`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_data_intelligence" "us" {
  regions = ["us"]
}

output "data_intelligence_https_egress" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 443
    destinations = data.xcsh_network_data_intelligence.us.cidr_blocks
  }
}
```
