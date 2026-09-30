---
page_title: "xcsh_site_upgrade_os"
subcategory: ""
description: "xcsh_site_upgrade_os for xcsh_site_upgrade_os."
xcsh_docs: {"aliases": [], "body_bytes": 1214, "body_sha256": "sha256:7fa7cddd412baeeb52f8d03bb526e00708acde9bf03b3587f57329ca8a4ccf64", "child_ids": ["xcsh-docs:actions:site_upgrade_os:reference", "xcsh-docs:actions:site_upgrade_os:examples", "xcsh-docs:actions:site_upgrade_os:lifecycle"], "collection_id": "xcsh-docs:actions:site_upgrade_os:collection", "completeness": "complete", "id": "xcsh-docs:actions:site_upgrade_os:fundamentals", "parent_id": null, "path": "documentation/actions/site_upgrade_os/index.md", "provider_name": "site_upgrade_os", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/site_upgrade_os/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_site_upgrade_os for xcsh_site_upgrade_os.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
