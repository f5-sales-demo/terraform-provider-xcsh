---
page_title: "Action"
subcategory: ""
description: "Action for xcsh_device_intelligence_subscribe."
xcsh_docs: {"aliases": [], "body_bytes": 1227, "body_sha256": "sha256:a10839251a1de8e372b5cc8499e46650cf5870f137fb5d74ecf2c2f0f0ee412c", "child_ids": [], "collection_id": "xcsh-docs:actions:device_intelligence_subscribe:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f3174d51768147466c8e2454428347852090bb677186e716ca6cee429d892fbc", "source_path": "examples/actions/xcsh_device_intelligence_subscribe/action.tf", "validation": "terraform validate"}, "id": "xcsh-docs:actions:device_intelligence_subscribe:example:action", "parent_id": "xcsh-docs:actions:device_intelligence_subscribe:examples", "path": "documentation/actions/device_intelligence_subscribe/examples/action/index.md", "provider_name": "device_intelligence_subscribe", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "actions", "role": "example", "schema_path": ["action"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/device_intelligence_subscribe/examples/action/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Action for xcsh_device_intelligence_subscribe.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_subscribe/examples/)
- [xcsh_device_intelligence_subscribe](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/device_intelligence_subscribe/)
