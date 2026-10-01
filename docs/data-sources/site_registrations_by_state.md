---
page_title: "xcsh_site_registrations_by_state"
subcategory: ""
description: "xcsh_site_registrations_by_state for xcsh_site_registrations_by_state."
xcsh_docs: {"aliases": [], "body_bytes": 1227, "body_sha256": "sha256:35602d22599417c170d1f5d87a458a1a84a0a027bfcce04d60ffbd268444344c", "canonical_id": "xcsh-docs:data-sources:site_registrations_by_state:fundamentals", "child_ids": ["xcsh-docs:data-sources:site_registrations_by_state:reference", "xcsh-docs:data-sources:site_registrations_by_state:examples"], "collection_id": "xcsh-docs:data-sources:site_registrations_by_state:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:site_registrations_by_state:fundamentals", "parent_id": null, "path": "docs/data-sources/site_registrations_by_state.md", "provider_name": "site_registrations_by_state", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_state/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_site_registrations_by_state for xcsh_site_registrations_by_state.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_site_registrations_by_state

Breadcrumbs:

- xcsh_site_registrations_by_state

List Customer Edge registrations by state.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteRegistrationsByState DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations_by_state" "example" {
  state = "NOTSET"
}

output "site_registrations_by_state_result" {
  value = data.xcsh_site_registrations_by_state.example
}
```

## Root configuration

Required root properties: `state`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--site_registrations_by_state--reference.md)
- [Examples](../guides/data-sources--site_registrations_by_state--examples.md)
