---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_device_intelligence_devices."
xcsh_docs: {"aliases": [], "body_bytes": 1272, "body_sha256": "sha256:6c7f25db7d2d33130f6637ca4568a841e450670c5d36fe4e1cb4be28d6de0ce9", "child_ids": [], "collection_id": "xcsh-docs:data-sources:device_intelligence_devices:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:59c1a13d7be54a7427ce8e310798e9737dbdf20c08bef5ed916ec4310ce79e92", "source_path": "examples/data-sources/xcsh_device_intelligence_devices/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:device_intelligence_devices:example:data-source", "parent_id": "xcsh-docs:data-sources:device_intelligence_devices:examples", "path": "documentation/data-sources/device_intelligence_devices/examples/data-source/index.md", "provider_name": "device_intelligence_devices", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_devices/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_device_intelligence_devices.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_device_intelligence_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_devices/data-source.tf`; digest `sha256:59c1a13d7be54a7427ce8e310798e9737dbdf20c08bef5ed916ec4310ce79e92`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/examples/)
- [xcsh_device_intelligence_devices](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/)
