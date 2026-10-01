---
page_title: "xcsh_site_signatures_update"
subcategory: ""
description: "xcsh_site_signatures_update for xcsh_site_signatures_update."
xcsh_docs: {"aliases": [], "body_bytes": 1170, "body_sha256": "sha256:2c86a80f4331df7e96614fccccf4d5ba242440cd82fb05544613dd736b5c358f", "canonical_id": "xcsh-docs:actions:site_signatures_update:fundamentals", "child_ids": ["xcsh-docs:actions:site_signatures_update:reference", "xcsh-docs:actions:site_signatures_update:examples", "xcsh-docs:actions:site_signatures_update:lifecycle"], "collection_id": "xcsh-docs:actions:site_signatures_update:collection", "completeness": "complete", "id": "xcsh-docs:actions:site_signatures_update:fundamentals", "parent_id": null, "path": "docs/actions/site_signatures_update.md", "provider_name": "site_signatures_update", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "actions", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_signatures_update/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_site_signatures_update for xcsh_site_signatures_update.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site_signatures_update

Breadcrumbs:

- xcsh_site_signatures_update

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteSignaturesUpdate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_site_signatures_update" "example" {
  config {
    namespace = "example-value"
  }
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/actions--site_signatures_update--reference.md)
- [Examples](../guides/actions--site_signatures_update--examples.md)
- [Lifecycle](../guides/actions--site_signatures_update--lifecycle.md)
