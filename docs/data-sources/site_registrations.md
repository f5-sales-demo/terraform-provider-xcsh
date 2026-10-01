---
page_title: "xcsh_site_registrations"
subcategory: ""
description: "xcsh_site_registrations for xcsh_site_registrations."
xcsh_docs: {"aliases": [], "body_bytes": 1163, "body_sha256": "sha256:65949612252bf6b386f1ae52c4a6847c8e496d09ecf5dbf11a8f454c98c19756", "canonical_id": "xcsh-docs:data-sources:site_registrations:fundamentals", "child_ids": ["xcsh-docs:data-sources:site_registrations:reference", "xcsh-docs:data-sources:site_registrations:examples"], "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations:fundamentals", "parent_id": null, "path": "docs/data-sources/site_registrations.md", "provider_name": "site_registrations", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_site_registrations for xcsh_site_registrations.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site_registrations

Breadcrumbs:

- xcsh_site_registrations

List Customer Edge registrations.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteRegistrations DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations" "example" {
  namespace = "example-value"
}

output "site_registrations_result" {
  value = data.xcsh_site_registrations.example
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--site_registrations--reference.md)
- [Examples](../guides/data-sources--site_registrations--examples.md)
