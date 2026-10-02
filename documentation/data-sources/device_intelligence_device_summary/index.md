---
page_title: "xcsh_device_intelligence_device_summary"
subcategory: ""
description: "Resource creation operation."
xcsh_docs: {"aliases": ["device intelligence device summary"], "body_bytes": 1412, "body_sha256": "sha256:799aec5bc8912afcec4cd9d7ee50b15609c00d8316825b34db02c4016fb6fd96", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_device_summary:reference", "xcsh-docs:data-sources:device_intelligence_device_summary:examples"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_device_summary:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_summary:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/device_intelligence_device_summary/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_device_summary", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3320320012000030-1120213132323310-0033312123010313-0312323233302312-0102031103102222-2230101011132101-1321001213032302-3200123203212021", "registry_path": "docs/data-sources/device_intelligence_device_summary.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_summary/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource creation operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
