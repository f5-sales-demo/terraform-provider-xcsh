---
page_title: "xcsh_site_upgrade_os"
subcategory: ""
description: "Request an in-place site operating-system upgrade."
xcsh_docs: {"aliases": ["site upgrade os"], "body_bytes": 1313, "body_sha256": "sha256:e3fc88ad588a4ab46850d09835bf107ef4e1dad153cce2fac9ffbf4c935e3eac", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:site_upgrade_os:reference", "xcsh-docs:actions:site_upgrade_os:examples", "xcsh-docs:actions:site_upgrade_os:lifecycle"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:site_upgrade_os:collection", "completeness": "complete", "id": "xcsh-docs:actions:site_upgrade_os:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/site_upgrade_os/index.md", "product": "distributed-cloud", "provider_name": "site_upgrade_os", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "actions", "registry_anchor": "canonical-0021103023333201-3221302120300023-0003110201010333-0330313301123130-0121011331210220-0002322000123231-1001203221200011-1100310302120313", "registry_path": "docs/actions/site_upgrade_os.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_upgrade_os/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Request an in-place site operating-system upgrade.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": [], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site_upgrade_os

Breadcrumbs:

- xcsh_site_upgrade_os

Request an in-place site operating-system upgrade.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteUpgradeOS Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_site_upgrade_os" "example" {
  config {
    site       = "example-value"
    os_version = "example-value"
  }
}
```

## Root configuration

Required root properties: `os_version`, `site`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_os/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_os/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/site_upgrade_os/lifecycle/)
