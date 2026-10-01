---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_upgrade_status."
xcsh_docs: {"aliases": [], "body_bytes": 1355, "body_sha256": "sha256:23350fef854a230fdf0531c8c77328195aa43db37d11f9d3cee5f5b64c0a248d", "canonical_id": "xcsh-docs:data-sources:site_upgrade_status:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_upgrade_status:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3c8a578ab685be3268d5b0dab58926aa035dcddae8cb2d3ebaa0f6a17186a683", "source_path": "examples/data-sources/xcsh_site_upgrade_status/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_upgrade_status:example:data-source", "parent_id": "xcsh-docs:data-sources:site_upgrade_status:examples", "path": "docs/guides/data-sources--site_upgrade_status--example--data-source.md", "provider_name": "site_upgrade_status", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_upgrade_status/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_site_upgrade_status.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_upgrade_status](../data-sources/site_upgrade_status.md)
- [Examples](data-sources--site_upgrade_status--examples.md)
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

- [Examples](data-sources--site_upgrade_status--examples.md)
- [xcsh_site_upgrade_status](../data-sources/site_upgrade_status.md)
