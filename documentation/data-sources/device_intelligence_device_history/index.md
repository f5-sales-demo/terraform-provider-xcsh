---
page_title: "xcsh_device_intelligence_device_history"
subcategory: ""
description: "Reads Device Intelligence Device History information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["device intelligence device history"], "body_bytes": 1476, "body_sha256": "sha256:2d2084d86b1e93a05e373c3431dfbbd45ff3c5377fd5196f30f56b2867f1a0e9", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_device_history:reference", "xcsh-docs:data-sources:device_intelligence_device_history:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_device_history:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_history:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/device_intelligence_device_history/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_device_history", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1033022013313010-2230031200002233-1320022023130030-0022221001122220-0201311013330313-2011320120303230-1323223312110203-0123022120001321", "registry_path": "docs/data-sources/device_intelligence_device_history.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_history/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Reads Device Intelligence Device History information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": [], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_device_intelligence_device_history

Breadcrumbs:

- xcsh_device_intelligence_device_history

Reads Device Intelligence Device History information from F5 Distributed Cloud.

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_history/examples/)
