---
page_title: "xcsh_bot_detection_rule"
subcategory: ""
description: "Reads Bot Detection Rule information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["bot detection rule"], "body_bytes": 1381, "body_sha256": "sha256:176047767eef06afffc76d6b5a6cf9c8d6aab5748c93d6b1b4207729167353da", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_detection_rule:reference", "xcsh-docs:data-sources:bot_detection_rule:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_detection_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_detection_rule:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/bot_detection_rule/index.md", "product": "distributed-cloud", "provider_name": "bot_detection_rule", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-2221001030010210-1213310213300133-3302112123210020-1030311021023312-0232211203112101-3232200001000312-1300200232311310-2121130011001102", "registry_path": "docs/data-sources/bot_detection_rule.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_detection_rule/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Reads Bot Detection Rule information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": [], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bot_detection_rule

Breadcrumbs:

- xcsh_bot_detection_rule

Reads Bot Detection Rule information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/examples/)
