---
page_title: "xcsh_site_cloud_init"
subcategory: ""
description: "xcsh_site_cloud_init for xcsh_site_cloud_init."
xcsh_docs: {"aliases": [], "body_bytes": 1223, "body_sha256": "sha256:e9c83ad55bc6732dc554d4a1ea4b7ada722bb71a384c932d53eac713ce8af23d", "canonical_id": "xcsh-docs:data-sources:site_cloud_init:fundamentals", "child_ids": ["xcsh-docs:data-sources:site_cloud_init:reference", "xcsh-docs:data-sources:site_cloud_init:examples"], "collection_id": "xcsh-docs:data-sources:site_cloud_init:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_cloud_init:fundamentals", "parent_id": null, "path": "docs/data-sources/site_cloud_init.md", "provider_name": "site_cloud_init", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_cloud_init/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_site_cloud_init for xcsh_site_cloud_init.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site_cloud_init

Breadcrumbs:

- xcsh_site_cloud_init

Retrieve Customer Edge cloud-init template.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteCloudInit DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_cloud_init" "example" {
  provider_ref = "example-value"
  site_name    = "example-value"
}

output "site_cloud_init_result" {
  value     = data.xcsh_site_cloud_init.example
  sensitive = true
}
```

## Root configuration

Required root properties: `provider_ref`, `site_name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--site_cloud_init--reference.md)
- [Examples](../guides/data-sources--site_cloud_init--examples.md)
