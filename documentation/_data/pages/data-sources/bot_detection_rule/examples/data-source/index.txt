---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_bot_detection_rule."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1109, "body_sha256": "sha256:d055116fa328ddb3540286600b6ab1cca90cc1baaef68fe75c26771d64e29793", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_detection_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fb0cc1608165dd7e57a5b8e5f57c447871327b816ef3b596ef84d5ce798774d2", "source_path": "examples/data-sources/xcsh_bot_detection_rule/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:bot_detection_rule:example:data-source", "parent_id": "xcsh-docs:data-sources:bot_detection_rule:examples", "path": "documentation/data-sources/bot_detection_rule/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "bot_detection_rule", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-1111333301020220-2003002210000223-2332322201122313-1231303101202201-1202013020013331-1120313321311220-0120012013330230-3230022331003020", "registry_path": "docs/guides/data-sources--bot_detection_rule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_detection_rule/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_bot_detection_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": [], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
