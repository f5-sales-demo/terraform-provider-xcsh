---
page_title: "xcsh_device_intelligence_device_summary"
subcategory: ""
description: "xcsh_device_intelligence_device_summary for xcsh_device_intelligence_device_summary."
xcsh_docs: {"aliases": [], "body_bytes": 1412, "body_sha256": "sha256:799aec5bc8912afcec4cd9d7ee50b15609c00d8316825b34db02c4016fb6fd96", "child_ids": ["xcsh-docs:data-sources:device_intelligence_device_summary:reference", "xcsh-docs:data-sources:device_intelligence_device_summary:examples"], "collection_id": "xcsh-docs:data-sources:device_intelligence_device_summary:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_summary:fundamentals", "parent_id": null, "path": "documentation/data-sources/device_intelligence_device_summary/index.md", "provider_name": "device_intelligence_device_summary", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_summary/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_device_intelligence_device_summary for xcsh_device_intelligence_device_summary.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_device_intelligence_device_summary

Breadcrumbs:

- xcsh_device_intelligence_device_summary

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceDeviceSummary DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_device_summary" "example" {
  device_id = "example-value"
  namespace = "example-value"
}

output "device_intelligence_device_summary_result" {
  value = data.xcsh_device_intelligence_device_summary.example
}
```

## Root configuration

Required root properties: `device_id`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/examples/)
