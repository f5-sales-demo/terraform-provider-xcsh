---
page_title: "xcsh_device_intelligence_multi_account_devices"
subcategory: ""
description: "Reads Device Intelligence Multi Account Devices information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["device intelligence multi account devices"], "body_bytes": 1495, "body_sha256": "sha256:d07dec2d740453de78a0f65f3ca7ef9facdd8ce804061d6936f9da81ed60c6a4", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_multi_account_devices:reference", "xcsh-docs:data-sources:device_intelligence_multi_account_devices:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_multi_account_devices:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/device_intelligence_multi_account_devices/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_multi_account_devices", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3113000302131120-2330303110103000-2313133302100232-0230010032333310-2102210302212230-1301202001211023-1123103101120300-3031201323311000", "registry_path": "docs/data-sources/device_intelligence_multi_account_devices.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_multi_account_devices/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reads Device Intelligence Multi Account Devices information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_device_intelligence_multi_account_devices

Breadcrumbs:

- xcsh_device_intelligence_multi_account_devices

Reads Device Intelligence Multi Account Devices information from F5 Distributed Cloud.

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_multi_account_devices/examples/)
