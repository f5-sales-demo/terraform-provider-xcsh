---
page_title: "xcsh_network_policy_rule"
subcategory: ""
description: "Manages network policy rule with configured parameters in specified namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["network policy rule"], "body_bytes": 1417, "body_sha256": "sha256:1b11aee176bc6099577da905ec94c77e2cc9cbcccae11953f8fb98a2add0be9a", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:network_policy_rule:reference", "xcsh-docs:data-sources:network_policy_rule:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_policy_rule:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_policy_rule/index.md", "product": "distributed-cloud", "provider_name": "network_policy_rule", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3010202012003101-0111222211322131-3011002003120212-1112301300332113-3123210312031031-1122031111001132-2013123210322220-2101202001331103", "registry_path": "docs/data-sources/network_policy_rule.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_rule/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Manages network policy rule with configured parameters in specified namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["network_policy_ruleCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_policy_rule

Breadcrumbs:

- xcsh_network_policy_rule

Manages network policy rule with configured parameters in specified namespace in F5 Distributed
Cloud.

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/examples/)
