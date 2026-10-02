---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_detection_rule."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1355, "body_sha256": "sha256:58611b78171c755d0fdfce18eace6962d18b84303aef60bbf8839ac0be2a769d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_detection_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fb0cc1608165dd7e57a5b8e5f57c447871327b816ef3b596ef84d5ce798774d2", "source_path": "examples/data-sources/xcsh_bot_detection_rule/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_detection_rule:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_detection_rule:examples", "path": "documentation/data-sources/bot_detection_rule/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_detection_rule", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1111333301020220-2003002210000223-2332322201122313-1231303101202201-1202013020013331-1120313321311220-0120012013330230-3230022331003020", "registry_path": "docs/guides/data-sources--bot_detection_rule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_detection_rule/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_bot_detection_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_bot_detection_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/examples/)
- [xcsh_bot_detection_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/)
