---
page_title: "xcsh_site_image"
subcategory: ""
description: "xcsh_site_image for xcsh_site_image."
xcsh_docs: {"aliases": [], "body_bytes": 1214, "body_sha256": "sha256:9b2987cadb36e305f4572db9fc33abab13c932c877fcfb6ac48c8f65844d92bb", "canonical_id": "xcsh-docs:data-sources:site_image:fundamentals", "child_ids": ["xcsh-docs:data-sources:site_image:reference", "xcsh-docs:data-sources:site_image:examples"], "collection_id": "xcsh-docs:data-sources:site_image:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_image:fundamentals", "parent_id": null, "path": "docs/data-sources/site_image.md", "provider_name": "site_image", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_image/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_site_image for xcsh_site_image.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

- [Property reference](../guides/data-sources--site_image--reference.md)
- [Examples](../guides/data-sources--site_image--examples.md)
