---
page_title: "xcsh_device_intelligence_subscribe"
subcategory: ""
description: "Subscribes to Device Intelligence updates."
xcsh_docs: {"aliases": ["device intelligence subscribe"], "body_bytes": 1334, "body_sha256": "sha256:f9d2280555876840cf4c48954a91c4c4e0d9ee4c26611da5912c68ba74b6dade", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:device_intelligence_subscribe:reference", "xcsh-docs:actions:device_intelligence_subscribe:examples", "xcsh-docs:actions:device_intelligence_subscribe:lifecycle"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:device_intelligence_subscribe:collection", "completeness": "complete", "id": "xcsh-docs:actions:device_intelligence_subscribe:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/device_intelligence_subscribe/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_subscribe", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "actions", "registry_anchor": "canonical-2110103333002230-1333113312300003-3031011200033223-1012111303330101-3123021221010332-0102310002201001-2232032323131011-1213333131130320", "registry_path": "docs/actions/device_intelligence_subscribe.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/device_intelligence_subscribe/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Subscribes to Device Intelligence updates.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": [], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_device_intelligence_subscribe

Breadcrumbs:

- xcsh_device_intelligence_subscribe

Subscribes to Device Intelligence updates.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceSubscribe Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_device_intelligence_subscribe" "example" {
  config {
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_subscribe/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_subscribe/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_subscribe/lifecycle/)
