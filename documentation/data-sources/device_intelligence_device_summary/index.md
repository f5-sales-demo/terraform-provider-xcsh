---
page_title: "xcsh_device_intelligence_device_summary"
subcategory: ""
description: "Reads Device Intelligence Device Summary information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["device intelligence device summary"], "body_bytes": 1476, "body_sha256": "sha256:b17eb0ae2c4051037d6d183e8de2399416d7c7d2ca547d61715b2642a4735741", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:device_intelligence_device_summary:reference", "xcsh-docs:data-sources:device_intelligence_device_summary:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_device_summary:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:device_intelligence_device_summary:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/device_intelligence_device_summary/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_device_summary", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3320320012000030-1120213132323310-0033312123010313-0312323233302312-0102031103102222-2230101011132101-1321001213032302-3200123203212021", "registry_path": "docs/data-sources/device_intelligence_device_summary.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_device_summary/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Reads Device Intelligence Device Summary information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": [], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_device_intelligence_device_summary

Breadcrumbs:

- xcsh_device_intelligence_device_summary

Reads Device Intelligence Device Summary information from F5 Distributed Cloud.

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_device_summary/examples/)
