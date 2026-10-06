---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_upgrade_status."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1312, "body_sha256": "sha256:6eb21a8b23938d9034241e55ba62a0ccd2ba76bcfca91d1fb69455693fe68924", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_upgrade_status:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3c8a578ab685be3268d5b0dab58926aa035dcddae8cb2d3ebaa0f6a17186a683", "source_path": "examples/data-sources/xcsh_site_upgrade_status/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_upgrade_status:example:data-source", "parent_id": "xcsh-docs:data-sources:site_upgrade_status:examples", "path": "documentation/data-sources/site_upgrade_status/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "site_upgrade_status", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-2300222103311212-2122102320121120-1002232212000123-0112301033311311-3211020331002101-0123231331013330-0322323313202022-3000302011202323", "registry_path": "docs/guides/data-sources--site_upgrade_status--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_upgrade_status/examples/data-source/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Data source for xcsh_site_upgrade_status.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": [], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_upgrade_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_upgrade_status/data-source.tf`; digest `sha256:3c8a578ab685be3268d5b0dab58926aa035dcddae8cb2d3ebaa0f6a17186a683`.

```terraform
# Observe upgrade eligibility or wait for supplied software and OS targets to
# be installed with the site back ONLINE.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 7.3.0"
    }
  }
}

data "xcsh_site_upgrade_status" "site" {
  site = "example-smsv2-site"

  expected_software_version = "crt-20260201-0179"
  expected_os_version       = "9.2026.17"
  wait                      = true
  timeout_seconds           = 7200
  poll_interval_seconds     = 30
}

output "upgrade_converged" {
  value = data.xcsh_site_upgrade_status.site.target_converged
}
```
