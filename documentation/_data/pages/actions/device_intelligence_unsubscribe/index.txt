---
page_title: "xcsh_device_intelligence_unsubscribe"
subcategory: ""
description: "Unsubscribes from Device Intelligence updates."
xcsh_docs: {"aliases": ["device intelligence unsubscribe"], "body_bytes": 1352, "body_sha256": "sha256:ae2633928b23b52486d444e34c275dd6d95c4214827a9a094a3bd8a99d260976", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:actions:device_intelligence_unsubscribe:reference", "xcsh-docs:actions:device_intelligence_unsubscribe:examples", "xcsh-docs:actions:device_intelligence_unsubscribe:lifecycle"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:device_intelligence_unsubscribe:collection", "completeness": "complete", "id": "xcsh-docs:actions:device_intelligence_unsubscribe:fundamentals", "parent_id": "xcsh-docs:actions:xcsh:navigation", "path": "documentation/actions/device_intelligence_unsubscribe/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_unsubscribe", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "actions", "registry_anchor": "canonical-3110012200313320-0120221200221312-0230333311202000-0331023323320332-1101121132213210-3031113003001131-3201203222213110-2210102303020312", "registry_path": "docs/actions/device_intelligence_unsubscribe.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/device_intelligence_unsubscribe/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Unsubscribes from Device Intelligence updates.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_device_intelligence_unsubscribe

Breadcrumbs:

- xcsh_device_intelligence_unsubscribe

Unsubscribes from Device Intelligence updates.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DeviceIntelligenceUnsubscribe Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_device_intelligence_unsubscribe" "example" {
  config {
  }
}
```

## Root configuration

Required root properties: none. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_unsubscribe/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_unsubscribe/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_unsubscribe/lifecycle/)
