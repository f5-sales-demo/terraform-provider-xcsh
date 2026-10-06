---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_site_upgrade_sw."
xcsh_docs: {"aliases": ["action"], "body_bytes": 1100, "body_sha256": "sha256:5be4178b8bdf02860e8e39958cb981def6983e59b016c416c275707323f67bb2", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:site_upgrade_sw:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f543d404a49e0f6ff11128bf33aab9d859e0913de2e4fcdcef829bf992c11087", "source_path": "examples/actions/xcsh_site_upgrade_sw/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:site_upgrade_sw:example:action", "parent_id": "xcsh-docs:actions:site_upgrade_sw:examples", "path": "documentation/actions/site_upgrade_sw/examples/action/index.md", "product": "distributed-cloud", "provider_name": "site_upgrade_sw", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "actions", "registry_anchor": "canonical-1323323103032321-2101300321023232-2032302101012302-2313102031131130-1311333332310113-3013030213333101-0223213103012110-0013302000322232", "registry_path": "docs/guides/actions--site_upgrade_sw--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_upgrade_sw/examples/action/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Action for xcsh_site_upgrade_sw.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_site_upgrade_sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_sw/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_sw/examples/)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_site_upgrade_sw/action.tf`; digest `sha256:f543d404a49e0f6ff11128bf33aab9d859e0913de2e4fcdcef829bf992c11087`.

```terraform
# SiteUpgradeSw Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# The API accepts the upgrade request immediately; convergence is asynchronous.
# This action does not reconcile a site's pinned software_settings.
action "xcsh_site_upgrade_sw" "example" {
  config {
    site             = "example-value"
    software_version = "example-value"
  }
}
```
