---
page_title: "xcsh_device_intelligence_device_history"
subcategory: ""
description: "xcsh_device_intelligence_device_history for xcsh_device_intelligence_device_history."
xcsh_docs: {"aliases": [], "body_bytes": 1327, "body_sha256": "sha256:e35c20cedd7c607ff6b2b614d8e127e0d1fa0f1f227260e78f849c4282f94cf2", "canonical_id": "xcsh-docs:data-sources:device_intelligence_device_history:fundamentals", "child_ids": ["xcsh-docs:data-sources:device_intelligence_device_history:reference", "xcsh-docs:data-sources:device_intelligence_device_history:examples"], "collection_id": "xcsh-docs:data-sources:device_intelligence_device_history:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_history:fundamentals", "parent_id": null, "path": "docs/data-sources/device_intelligence_device_history.md", "provider_name": "device_intelligence_device_history", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_history/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_device_intelligence_device_history for xcsh_device_intelligence_device_history.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_device_intelligence_device_history

Breadcrumbs:

- xcsh_device_intelligence_device_history

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceDeviceHistory DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_device_history" "example" {
  device_id = "example-value"
  namespace = "example-value"
}

output "device_intelligence_device_history_result" {
  value = data.xcsh_device_intelligence_device_history.example
}
```

## Root configuration

Required root properties: `device_id`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--device_intelligence_device_history--reference.md)
- [Examples](../guides/data-sources--device_intelligence_device_history--examples.md)
