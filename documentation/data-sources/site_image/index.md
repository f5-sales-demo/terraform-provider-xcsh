---
page_title: "xcsh_site_image"
subcategory: ""
description: "xcsh_site_image for xcsh_site_image."
xcsh_docs: {"aliases": [], "body_bytes": 1398, "body_sha256": "sha256:1b207e85df2e6c0df3f366a20436a0fff462b3fd447330fc9712bb1aa4ba22a8", "child_ids": ["xcsh-docs:data-sources:site_image:reference", "xcsh-docs:data-sources:site_image:examples"], "collection_id": "xcsh-docs:data-sources:site_image:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_image:fundamentals", "parent_id": null, "path": "documentation/data-sources/site_image/index.md", "provider_name": "site_image", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_image/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_site_image for xcsh_site_image.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site_image

Breadcrumbs:

- xcsh_site_image

Resolve the current KVM image using exactly one Site owned by the named SMSv2 configuration. No
caller UID or static fallback is supported. Verify the artifact MD5 before use; image resolution
does not imply successful boot.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteImage DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_image" "example" {
  site_name = "example-value"
}

output "site_image_result" {
  value     = data.xcsh_site_image.example
  sensitive = true
}
```

## Root configuration

Required root properties: `site_name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/examples/)
