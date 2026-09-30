---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_device_intelligence_subscribe."
xcsh_docs: {"aliases": [], "body_bytes": 1128, "body_sha256": "sha256:72e76095f537c971ad193c76d3bc48b660492f2af6c0c27c753e6e0f3ee64753", "child_ids": [], "collection_id": "xcsh-docs:actions:device_intelligence_subscribe:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f3174d51768147466c8e2454428347852090bb677186e716ca6cee429d892fbc", "source_path": "examples/actions/xcsh_device_intelligence_subscribe/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:device_intelligence_subscribe:example:action", "parent_id": "xcsh-docs:actions:device_intelligence_subscribe:examples", "path": "documentation/actions/device_intelligence_subscribe/examples/action/index.md", "provider_name": "device_intelligence_subscribe", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "role": "example", "schema_path": ["action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/device_intelligence_subscribe/examples/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Action for xcsh_device_intelligence_subscribe.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_subscribe/examples/)
- [xcsh_device_intelligence_subscribe](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_subscribe/)
