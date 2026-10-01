---
page_title: "xcsh_site_upgrade_sw"
subcategory: ""
description: "xcsh_site_upgrade_sw for xcsh_site_upgrade_sw."
xcsh_docs: {"aliases": [], "body_bytes": 1344, "body_sha256": "sha256:dfa994291a2eb096148706e8d3f99a902a5edadaf002293547470361fcb9df9d", "canonical_id": "xcsh-docs:actions:site_upgrade_sw:fundamentals", "child_ids": ["xcsh-docs:actions:site_upgrade_sw:reference", "xcsh-docs:actions:site_upgrade_sw:examples", "xcsh-docs:actions:site_upgrade_sw:lifecycle"], "collection_id": "xcsh-docs:actions:site_upgrade_sw:collection", "completeness": "complete", "id": "xcsh-docs:actions:site_upgrade_sw:fundamentals", "parent_id": null, "path": "docs/actions/site_upgrade_sw.md", "provider_name": "site_upgrade_sw", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "actions", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_upgrade_sw/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_site_upgrade_sw for xcsh_site_upgrade_sw.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site_upgrade_sw

Breadcrumbs:

- xcsh_site_upgrade_sw

Request an in-place site software upgrade.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `site`, `software_version`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/actions--site_upgrade_sw--reference.md)
- [Examples](../guides/actions--site_upgrade_sw--examples.md)
- [Lifecycle](../guides/actions--site_upgrade_sw--lifecycle.md)
