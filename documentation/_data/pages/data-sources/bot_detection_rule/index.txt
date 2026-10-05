---
page_title: "xcsh_bot_detection_rule"
subcategory: ""
description: "Manages a Bot Detection Rule resource in F5 Distributed Cloud for get bot detection rule. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["bot detection rule"], "body_bytes": 1433, "body_sha256": "sha256:3997f46a1e9fd5c36812f1f56cca7ea03b837ee508abba5bd83553ac8ef8ba94", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:bot_detection_rule:reference", "xcsh-docs:data-sources:bot_detection_rule:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:bot_detection_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:bot_detection_rule:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/bot_detection_rule/index.md", "product": "distributed-cloud", "provider_name": "bot_detection_rule", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2221001030010210-1213310213300133-3302112123210020-1030311021023312-0232211203112101-3232200001000312-1300200232311310-2121130011001102", "registry_path": "docs/data-sources/bot_detection_rule.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/bot_detection_rule/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages a Bot Detection Rule resource in F5 Distributed Cloud for get bot detection rule. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": [], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
