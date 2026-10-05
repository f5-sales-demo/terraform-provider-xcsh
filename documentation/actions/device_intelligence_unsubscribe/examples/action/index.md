---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_device_intelligence_unsubscribe."
xcsh_docs: {"aliases": ["action"], "body_bytes": 1245, "body_sha256": "sha256:1d1e8745c2d18580e65393e4a05d097f184535aecf4029000d6d6a8951bfcd83", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:actions:device_intelligence_unsubscribe:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:391d15fe0fb026a460e8e78fbc03f16af6138839116ddca4fd7a4862c6cad9eb", "source_path": "examples/actions/xcsh_device_intelligence_unsubscribe/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:device_intelligence_unsubscribe:example:action", "parent_id": "xcsh-docs:actions:device_intelligence_unsubscribe:examples", "path": "documentation/actions/device_intelligence_unsubscribe/examples/action/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_unsubscribe", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "actions", "registry_anchor": "canonical-3311111003211210-3123212122310002-3232321300010031-3223330211231201-1201330003122032-3133031110321323-1300222223323000-3201321333320212", "registry_path": "docs/guides/actions--device_intelligence_unsubscribe--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["action"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/device_intelligence_unsubscribe/examples/action/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Action for xcsh_device_intelligence_unsubscribe.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_unsubscribe/examples/)
- [xcsh_device_intelligence_unsubscribe](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_unsubscribe/)
