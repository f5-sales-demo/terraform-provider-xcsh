---
page_title: "xcsh_network_policy_rule"
subcategory: ""
description: "xcsh_network_policy_rule for xcsh_network_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1468, "body_sha256": "sha256:244f78a0ac888d945f4468a8581a98c8f9d467a54057d64c3369292d7e5ab5a5", "canonical_id": "xcsh-docs:resources:network_policy_rule:fundamentals", "child_ids": ["xcsh-docs:resources:network_policy_rule:reference", "xcsh-docs:resources:network_policy_rule:examples", "xcsh-docs:resources:network_policy_rule:import", "xcsh-docs:resources:network_policy_rule:timeouts"], "collection_id": "xcsh-docs:resources:network_policy_rule:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy_rule:fundamentals", "parent_id": null, "path": "docs/resources/network_policy_rule.md", "provider_name": "network_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_rule/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_network_policy_rule for xcsh_network_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
# NetworkPolicyRule Resource Example
# Manages network policy rule with configured parameters in specified namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicyRule configuration
resource "xcsh_network_policy_rule" "example" {
  name      = "example-network-policy-rule"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--network_policy_rule--reference.md)
- [Examples](../guides/resources--network_policy_rule--examples.md)
- [Import](../guides/resources--network_policy_rule--import.md)
- [Timeouts](../guides/resources--network_policy_rule--timeouts.md)
