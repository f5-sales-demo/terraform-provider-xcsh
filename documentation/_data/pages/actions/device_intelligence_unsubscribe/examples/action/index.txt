---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_device_intelligence_unsubscribe."
xcsh_docs: {"aliases": ["action"], "body_bytes": 970, "body_sha256": "sha256:42c9c142ba2aeadb583252774aa034b02fdb5f18128dc45c18da0d1094d51cdf", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:device_intelligence_unsubscribe:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:391d15fe0fb026a460e8e78fbc03f16af6138839116ddca4fd7a4862c6cad9eb", "source_path": "examples/actions/xcsh_device_intelligence_unsubscribe/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:device_intelligence_unsubscribe:example:action", "parent_id": "xcsh-docs:actions:device_intelligence_unsubscribe:examples", "path": "documentation/actions/device_intelligence_unsubscribe/examples/action/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_unsubscribe", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "actions", "registry_anchor": "canonical-3311111003211210-3123212122310002-3232321300010031-3223330211231201-1201330003122032-3133031110321323-1300222223323000-3201321333320212", "registry_path": "docs/guides/actions--device_intelligence_unsubscribe--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/device_intelligence_unsubscribe/examples/action/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Action for xcsh_device_intelligence_unsubscribe.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": [], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Action

Breadcrumbs:

- [xcsh_device_intelligence_unsubscribe](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_unsubscribe/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_unsubscribe/examples/)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_device_intelligence_unsubscribe/action.tf`; digest `sha256:391d15fe0fb026a460e8e78fbc03f16af6138839116ddca4fd7a4862c6cad9eb`.

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
