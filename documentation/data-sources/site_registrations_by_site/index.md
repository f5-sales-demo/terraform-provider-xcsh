---
page_title: "xcsh_site_registrations_by_site"
subcategory: ""
description: "xcsh_site_registrations_by_site for xcsh_site_registrations_by_site."
xcsh_docs: {"aliases": [], "body_bytes": 1222, "body_sha256": "sha256:15549177d3a7f96f3beb22486290e8226a897420f7218992f32622956b63c3e4", "child_ids": ["xcsh-docs:data-sources:site_registrations_by_site:reference", "xcsh-docs:data-sources:site_registrations_by_site:examples"], "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_site:fundamentals", "parent_id": null, "path": "documentation/data-sources/site_registrations_by_site/index.md", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_site_registrations_by_site for xcsh_site_registrations_by_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_site_registrations_by_site

Breadcrumbs:

- xcsh_site_registrations_by_site

List registrations for a Customer Edge site.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteRegistrationsBySite DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations_by_site" "example" {
  site_name = "example-value"
}

output "site_registrations_by_site_result" {
  value = data.xcsh_site_registrations_by_site.example
}
```

## Root configuration

Required root properties: `site_name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/examples/)
