---
page_title: "xcsh_usb_policy"
subcategory: ""
description: "xcsh_usb_policy for xcsh_usb_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1196, "body_sha256": "sha256:4e4711d6fa22f1de28fbff8bf37a8bd7bcc8b41306fe373f56770174f21bb18e", "canonical_id": "xcsh-docs:data-sources:usb_policy:fundamentals", "child_ids": ["xcsh-docs:data-sources:usb_policy:reference", "xcsh-docs:data-sources:usb_policy:examples"], "collection_id": "xcsh-docs:data-sources:usb_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:usb_policy:fundamentals", "parent_id": null, "path": "docs/data-sources/usb_policy.md", "provider_name": "usb_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/usb_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_usb_policy for xcsh_usb_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["usb_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_usb_policy

Breadcrumbs:

- xcsh_usb_policy

Manages new USB policy object in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UsbPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing UsbPolicy by name
data "xcsh_usb_policy" "example" {
  name      = "example-usb-policy"
  namespace = "staging"
}

output "usb_policy_id" {
  value = data.xcsh_usb_policy.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--usb_policy--reference.md)
- [Examples](../guides/data-sources--usb_policy--examples.md)
