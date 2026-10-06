---
page_title: "xcsh_device_intelligence_subscribe"
subcategory: ""
description: "Subscribes to Device Intelligence updates."
xcsh_docs: {"aliases": ["device intelligence subscribe"], "body_bytes": 1334, "body_sha256": "sha256:f9d2280555876840cf4c48954a91c4c4e0d9ee4c26611da5912c68ba74b6dade", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:device_intelligence_subscribe:reference", "xcsh-docs:actions:device_intelligence_subscribe:examples", "xcsh-docs:actions:device_intelligence_subscribe:lifecycle"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:device_intelligence_subscribe:collection", "completeness": "complete", "id": "xcsh-docs:actions:device_intelligence_subscribe:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/device_intelligence_subscribe/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_subscribe", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "actions", "registry_anchor": "canonical-2110103333002230-1333113312300003-3031011200033223-1012111303330101-3123021221010332-0102310002201001-2232032323131011-1213333131130320", "registry_path": "docs/actions/device_intelligence_subscribe.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/device_intelligence_subscribe/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Subscribes to Device Intelligence updates.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
