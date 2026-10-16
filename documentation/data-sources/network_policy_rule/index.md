---
page_title: "xcsh_network_policy_rule"
subcategory: ""
description: "Reads Network Policy Rule information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["network policy rule"], "body_bytes": 1392, "body_sha256": "sha256:798b1bf2fe333c7de8d4a7af422b2088f30f3827a29202331bfbcf6e2ff4fb88", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:network_policy_rule:reference", "xcsh-docs:data-sources:network_policy_rule:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_rule:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_policy_rule/index.md", "product": "distributed-cloud", "provider_name": "network_policy_rule", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103", "registry_path": "docs/data-sources/network_policy_rule.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_rule/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Reads Network Policy Rule information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["network_policy_ruleCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_policy_rule

Breadcrumbs:

- xcsh_network_policy_rule

Reads Network Policy Rule information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicyRule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicyRule by name
data "xcsh_network_policy_rule" "example" {
  name      = "example-network-policy-rule"
  namespace = "staging"
}

output "network_policy_rule_id" {
  value = data.xcsh_network_policy_rule.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/examples/)
