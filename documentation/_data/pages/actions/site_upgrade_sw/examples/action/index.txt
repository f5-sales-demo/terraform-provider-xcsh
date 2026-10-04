---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_site_upgrade_sw."
xcsh_docs: {"aliases": ["action"], "body_bytes": 1327, "body_sha256": "sha256:6e91665a740e873b17347dde7a85026cb4f6526cdf8b9ee17c372aa1febe8df2", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:site_upgrade_sw:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f543d404a49e0f6ff11128bf33aab9d859e0913de2e4fcdcef829bf992c11087", "source_path": "examples/actions/xcsh_site_upgrade_sw/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:site_upgrade_sw:example:action", "parent_id": "xcsh-docs:actions:site_upgrade_sw:examples", "path": "documentation/actions/site_upgrade_sw/examples/action/index.md", "product": "distributed-cloud", "provider_name": "site_upgrade_sw", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "actions", "registry_anchor": "canonical-1323323103032321-2101300321023232-2032302101012302-2313102031131130-1311333332310113-3013030213333101-0223213103012110-0013302000322232", "registry_path": "docs/guides/actions--site_upgrade_sw--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_upgrade_sw/examples/action/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Action for xcsh_site_upgrade_sw.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_sw/examples/)
- [xcsh_site_upgrade_sw](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_sw/)
