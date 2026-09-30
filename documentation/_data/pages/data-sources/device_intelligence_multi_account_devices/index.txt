---
page_title: "xcsh_device_intelligence_multi_account_devices"
subcategory: ""
description: "xcsh_device_intelligence_multi_account_devices for xcsh_device_intelligence_multi_account_devices."
xcsh_docs: {"aliases": [], "body_bytes": 1325, "body_sha256": "sha256:4a1a28c01273eaf3ba3b7fb11dd0fc6473aead922f3a3cccbb57fd99432c18d3", "child_ids": ["xcsh-docs:data-sources:device_intelligence_multi_account_devices:reference", "xcsh-docs:data-sources:device_intelligence_multi_account_devices:examples"], "collection_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:fundamentals", "parent_id": null, "path": "documentation/data-sources/device_intelligence_multi_account_devices/index.md", "provider_name": "device_intelligence_multi_account_devices", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_multi_account_devices/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_device_intelligence_multi_account_devices for xcsh_device_intelligence_multi_account_devices.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_device_intelligence_multi_account_devices

Breadcrumbs:

- xcsh_device_intelligence_multi_account_devices

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceMultiAccountDevices DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_multi_account_devices" "example" {
  namespace = "example-value"
}

output "device_intelligence_multi_account_devices_result" {
  value = data.xcsh_device_intelligence_multi_account_devices.example
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/examples/)
