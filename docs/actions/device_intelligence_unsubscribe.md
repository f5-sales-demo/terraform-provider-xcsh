---
page_title: "xcsh_device_intelligence_unsubscribe"
subcategory: ""
description: "xcsh_device_intelligence_unsubscribe for xcsh_device_intelligence_unsubscribe."
xcsh_docs: {"aliases": [], "body_bytes": 1095, "body_sha256": "sha256:7496cdbd0158c7df5e66aef57bf0cc12c80621892b8ee5315847ab3d2a9d6f7e", "canonical_id": "xcsh-docs:actions:device_intelligence_unsubscribe:fundamentals", "child_ids": ["xcsh-docs:actions:device_intelligence_unsubscribe:reference", "xcsh-docs:actions:device_intelligence_unsubscribe:examples", "xcsh-docs:actions:device_intelligence_unsubscribe:lifecycle"], "collection_id": "xcsh-docs:actions:device_intelligence_unsubscribe:collection", "completeness": "complete", "id": "xcsh-docs:actions:device_intelligence_unsubscribe:fundamentals", "parent_id": null, "path": "docs/actions/device_intelligence_unsubscribe.md", "provider_name": "device_intelligence_unsubscribe", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/device_intelligence_unsubscribe/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_device_intelligence_unsubscribe for xcsh_device_intelligence_unsubscribe.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_device_intelligence_unsubscribe

Breadcrumbs:

- xcsh_device_intelligence_unsubscribe

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceUnsubscribe Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_device_intelligence_unsubscribe" "example" {
  config {
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/actions--device_intelligence_unsubscribe--reference.md)
- [Examples](../guides/actions--device_intelligence_unsubscribe--examples.md)
- [Lifecycle](../guides/actions--device_intelligence_unsubscribe--lifecycle.md)
