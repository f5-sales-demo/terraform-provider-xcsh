---
page_title: "xcsh_device_intelligence_devices"
subcategory: ""
description: "Resource creation operation."
xcsh_docs: {"aliases": ["device intelligence devices"], "body_bytes": 1314, "body_sha256": "sha256:cdcb35f3fc489cbb742d537d1f7e05847298fdfb2b42fc6fca348f8970dce89d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_devices:reference", "xcsh-docs:data-sources:device_intelligence_devices:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_devices:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_devices:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/device_intelligence_devices/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_devices", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1321010130030011-2203222211110110-0002220011211330-1110301000223122-2213011310020223-2303303332130221-2212020001131033-0330011210030333", "registry_path": "docs/data-sources/device_intelligence_devices.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_devices/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource creation operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_devices/examples/)
