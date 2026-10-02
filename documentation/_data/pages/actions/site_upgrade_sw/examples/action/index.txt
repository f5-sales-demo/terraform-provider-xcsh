---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_site_upgrade_sw."
xcsh_docs: {"aliases": ["action"], "body_bytes": 1327, "body_sha256": "sha256:6e91665a740e873b17347dde7a85026cb4f6526cdf8b9ee17c372aa1febe8df2", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:site_upgrade_sw:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f543d404a49e0f6ff11128bf33aab9d859e0913de2e4fcdcef829bf992c11087", "source_path": "examples/actions/xcsh_site_upgrade_sw/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:site_upgrade_sw:example:action", "parent_id": "xcsh-docs:actions:site_upgrade_sw:examples", "path": "documentation/actions/site_upgrade_sw/examples/action/index.md", "product": "distributed-cloud", "provider_name": "site_upgrade_sw", "provider_schema_digest": "sha256:d20ce271369d414e9d5661659e153a5441a9bde0053a466518eae9b8b5ab32fd", "provider_type": "actions", "registry_anchor": "canonical-1323323103032321-2101300321023232-2032302101012302-2313102031131130-1311333332310113-3013030213333101-0223213103012110-0013302000322232", "registry_path": "docs/guides/actions--site_upgrade_sw--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_upgrade_sw/examples/action/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Action for xcsh_site_upgrade_sw.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
