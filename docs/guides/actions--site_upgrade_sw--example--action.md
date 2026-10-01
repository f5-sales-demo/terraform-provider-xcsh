---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_site_upgrade_sw."
xcsh_docs: {"aliases": [], "body_bytes": 1121, "body_sha256": "sha256:6f812429efe3649ecabc694fdcd6d29cb8612570b0facc20310a04846864e564", "canonical_id": "xcsh-docs:actions:site_upgrade_sw:example:action", "child_ids": [], "collection_id": "xcsh-docs:actions:site_upgrade_sw:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f543d404a49e0f6ff11128bf33aab9d859e0913de2e4fcdcef829bf992c11087", "source_path": "examples/actions/xcsh_site_upgrade_sw/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:site_upgrade_sw:example:action", "parent_id": "xcsh-docs:actions:site_upgrade_sw:examples", "path": "docs/guides/actions--site_upgrade_sw--example--action.md", "provider_name": "site_upgrade_sw", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "publishing_destination": "registry", "role": "example", "schema_path": ["action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_upgrade_sw/examples/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Action for xcsh_site_upgrade_sw.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_site_upgrade_sw](../actions/site_upgrade_sw.md)
- [Examples](actions--site_upgrade_sw--examples.md)
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

- [Examples](actions--site_upgrade_sw--examples.md)
- [xcsh_site_upgrade_sw](../actions/site_upgrade_sw.md)
