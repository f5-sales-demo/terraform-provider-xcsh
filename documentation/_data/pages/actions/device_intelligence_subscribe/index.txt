---
page_title: "xcsh_device_intelligence_subscribe"
subcategory: ""
description: "Resource creation operation."
xcsh_docs: {"aliases": ["device intelligence subscribe"], "body_bytes": 1307, "body_sha256": "sha256:eecbdf89e79620889620b82f147d3ff28c543f8c4777c71c82c26932a6c41a28", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:device_intelligence_subscribe:reference", "xcsh-docs:actions:device_intelligence_subscribe:examples", "xcsh-docs:actions:device_intelligence_subscribe:lifecycle"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:device_intelligence_subscribe:collection", "completeness": "complete", "id": "xcsh-docs:actions:device_intelligence_subscribe:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/device_intelligence_subscribe/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_subscribe", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "actions", "registry_anchor": "canonical-2110103333002230-1333113312300003-3031011200033223-1012111303330101-3123021221010332-0102310002201001-2232032323131011-1213333131130320", "registry_path": "docs/actions/device_intelligence_subscribe.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/device_intelligence_subscribe/index.txt", "spec_pin_digest": "sha256:0e278b4afb598ac8628ac74dca6a1b2831cd8607de1d26f56e72e3461a7440c7", "summary": "Resource creation operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.0", "schema_components": [], "target_commit": "ac024ccbfb8b9f84412813f3e2ab9f5821937ef2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_device_intelligence_subscribe

Breadcrumbs:

- xcsh_device_intelligence_subscribe

Resource creation operation.

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_subscribe/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_subscribe/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_subscribe/lifecycle/)
