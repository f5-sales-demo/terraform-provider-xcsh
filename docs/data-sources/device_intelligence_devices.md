---
page_title: "xcsh_device_intelligence_devices"
subcategory: ""
description: "xcsh_device_intelligence_devices for xcsh_device_intelligence_devices."
xcsh_docs: {"aliases": [], "body_bytes": 1130, "body_sha256": "sha256:eead40bbb567daac83091f311835569795200192f5b7b78f6f2adb8477a6bf71", "canonical_id": "xcsh-docs:data-sources:device_intelligence_devices:fundamentals", "child_ids": ["xcsh-docs:data-sources:device_intelligence_devices:reference", "xcsh-docs:data-sources:device_intelligence_devices:examples"], "collection_id": "xcsh-docs:data-sources:device_intelligence_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_devices:fundamentals", "parent_id": null, "path": "docs/data-sources/device_intelligence_devices.md", "provider_name": "device_intelligence_devices", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_devices/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_device_intelligence_devices for xcsh_device_intelligence_devices.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_device_intelligence_devices

Breadcrumbs:

- xcsh_device_intelligence_devices

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceDevices DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_devices" "example" {
  namespace = "example-value"
}

output "device_intelligence_devices_result" {
  value = data.xcsh_device_intelligence_devices.example
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--device_intelligence_devices--reference.md)
- [Examples](../guides/data-sources--device_intelligence_devices--examples.md)
