---
page_title: "xcsh_device_intelligence_subscribe"
subcategory: ""
description: "xcsh_device_intelligence_subscribe for xcsh_device_intelligence_subscribe."
xcsh_docs: {"aliases": [], "body_bytes": 1307, "body_sha256": "sha256:eecbdf89e79620889620b82f147d3ff28c543f8c4777c71c82c26932a6c41a28", "child_ids": ["xcsh-docs:actions:device_intelligence_subscribe:reference", "xcsh-docs:actions:device_intelligence_subscribe:examples", "xcsh-docs:actions:device_intelligence_subscribe:lifecycle"], "collection_id": "xcsh-docs:actions:device_intelligence_subscribe:collection", "completeness": "complete", "id": "xcsh-docs:actions:device_intelligence_subscribe:fundamentals", "parent_id": null, "path": "documentation/actions/device_intelligence_subscribe/index.md", "provider_name": "device_intelligence_subscribe", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/device_intelligence_subscribe/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_device_intelligence_subscribe for xcsh_device_intelligence_subscribe.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
