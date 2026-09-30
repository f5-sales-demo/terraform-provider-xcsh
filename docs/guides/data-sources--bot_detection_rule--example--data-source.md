---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_detection_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1050, "body_sha256": "sha256:6058f7010bd8e32eeb66351e83f1eeeec001b1704ede80de9a4df37c83495e0f", "canonical_id": "xcsh-docs:data-sources:bot_detection_rule:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:bot_detection_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fb0cc1608165dd7e57a5b8e5f57c447871327b816ef3b596ef84d5ce798774d2", "source_path": "examples/data-sources/xcsh_bot_detection_rule/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_detection_rule:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_detection_rule:examples", "path": "docs/guides/data-sources--bot_detection_rule--example--data-source.md", "provider_name": "bot_detection_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_detection_rule/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_bot_detection_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_bot_detection_rule](../data-sources/bot_detection_rule.md)
- [Examples](data-sources--bot_detection_rule--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_bot_detection_rule/data-source.tf`; digest `sha256:fb0cc1608165dd7e57a5b8e5f57c447871327b816ef3b596ef84d5ce798774d2`.

```terraform
# BotDetectionRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotDetectionRule by name
data "xcsh_bot_detection_rule" "example" {
  name      = "example-bot-detection-rule"
  namespace = "staging"
}

output "bot_detection_rule_id" {
  value = data.xcsh_bot_detection_rule.example.id
}
```

## Next pages

- [Examples](data-sources--bot_detection_rule--examples.md)
- [xcsh_bot_detection_rule](../data-sources/bot_detection_rule.md)
