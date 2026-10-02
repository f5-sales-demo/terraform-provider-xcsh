---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_upgrade_status."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1561, "body_sha256": "sha256:a8c6329a036a59faad1573a7fb4f4b9d1b67b3dda0bf484444c3616551de2fb8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_upgrade_status:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3c8a578ab685be3268d5b0dab58926aa035dcddae8cb2d3ebaa0f6a17186a683", "source_path": "examples/data-sources/xcsh_site_upgrade_status/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_upgrade_status:example:data-source", "parent_id": "xcsh-docs:data-sources:site_upgrade_status:examples", "path": "documentation/data-sources/site_upgrade_status/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "site_upgrade_status", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2300222103311212-2122102320121120-1002232212000123-0112301033311311-3211020331002101-0123231331013330-0322323313202022-3000302011202323", "registry_path": "docs/guides/data-sources--site_upgrade_status--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_upgrade_status/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_site_upgrade_status.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/examples/)
- [xcsh_site_upgrade_status](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_upgrade_status/)
