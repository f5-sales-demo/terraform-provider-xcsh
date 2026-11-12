---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_device_intelligence_subscribe."
xcsh_docs: {"aliases": ["action"], "body_bytes": 958, "body_sha256": "sha256:49b0aecc0eca125facaa353eb8e022c1aaf9746f3f17a97829e2343e3e0838c0", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:device_intelligence_subscribe:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f3174d51768147466c8e2454428347852090bb677186e716ca6cee429d892fbc", "source_path": "examples/actions/xcsh_device_intelligence_subscribe/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:device_intelligence_subscribe:example:action", "parent_id": "xcsh-docs:actions:device_intelligence_subscribe:examples", "path": "documentation/actions/device_intelligence_subscribe/examples/action/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_subscribe", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "actions", "registry_anchor": "canonical-1121302300110102-3021032313203020-3332332311011113-1011031011003200-2221123022001303-1323323030121002-3221012303212200-3013111111113110", "registry_path": "docs/guides/actions--device_intelligence_subscribe--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/device_intelligence_subscribe/examples/action/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Action for xcsh_device_intelligence_subscribe.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": [], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_device_intelligence_subscribe](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_subscribe/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_subscribe/examples/)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_device_intelligence_subscribe/action.tf`; digest `sha256:f3174d51768147466c8e2454428347852090bb677186e716ca6cee429d892fbc`.

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
