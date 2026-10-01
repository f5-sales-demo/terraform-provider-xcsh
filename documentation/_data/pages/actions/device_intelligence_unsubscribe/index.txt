---
page_title: "xcsh_device_intelligence_unsubscribe"
subcategory: ""
description: "xcsh_device_intelligence_unsubscribe for xcsh_device_intelligence_unsubscribe."
xcsh_docs: {"aliases": [], "body_bytes": 1321, "body_sha256": "sha256:722af403224e5bf6386ea46e3b02b9f2cc198a38b844227267c4fac6ee6fc4bf", "child_ids": ["xcsh-docs:actions:device_intelligence_unsubscribe:reference", "xcsh-docs:actions:device_intelligence_unsubscribe:examples", "xcsh-docs:actions:device_intelligence_unsubscribe:lifecycle"], "collection_id": "xcsh-docs:actions:device_intelligence_unsubscribe:collection", "completeness": "complete", "id": "xcsh-docs:actions:device_intelligence_unsubscribe:fundamentals", "parent_id": null, "path": "documentation/actions/device_intelligence_unsubscribe/index.md", "provider_name": "device_intelligence_unsubscribe", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/device_intelligence_unsubscribe/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_device_intelligence_unsubscribe for xcsh_device_intelligence_unsubscribe.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_unsubscribe/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_unsubscribe/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_unsubscribe/lifecycle/)
