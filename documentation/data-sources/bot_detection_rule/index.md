---
page_title: "xcsh_bot_detection_rule"
subcategory: ""
description: "Manages a Bot Detection Rule resource in F5 Distributed Cloud for get bot detection rule. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["bot detection rule"], "body_bytes": 1433, "body_sha256": "sha256:3997f46a1e9fd5c36812f1f56cca7ea03b837ee508abba5bd83553ac8ef8ba94", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_detection_rule:reference", "xcsh-docs:data-sources:bot_detection_rule:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_detection_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_detection_rule:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/bot_detection_rule/index.md", "product": "distributed-cloud", "provider_name": "bot_detection_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2221001030010210-1213310213300133-3302112123210020-1030311021023312-0232211203112101-3232200001000312-1300200232311310-2121130011001102", "registry_path": "docs/data-sources/bot_detection_rule.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_detection_rule/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Bot Detection Rule resource in F5 Distributed Cloud for get bot detection rule. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bot_detection_rule

Breadcrumbs:

- xcsh_bot_detection_rule

Manages a Bot Detection Rule resource in F5 Distributed Cloud for get bot detection rule.
configuration. (read-only data source)

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/bot_detection_rule/examples/)
