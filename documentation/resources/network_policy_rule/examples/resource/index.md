---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_network_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1354, "body_sha256": "sha256:a9902c6363642c5e61fa7005b14d332ae8a6f1c796036883e986fbfc5b3b2ca5", "child_ids": [], "collection_id": "xcsh-docs:resources:network_policy_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0feaf66dfd5662ae652b573d32c61f50aa8eb63e0f1a54a1e0d029c6ad96ace9", "source_path": "examples/resources/xcsh_network_policy_rule/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:network_policy_rule:example:resource", "parent_id": "xcsh-docs:resources:network_policy_rule:examples", "path": "documentation/resources/network_policy_rule/examples/resource/index.md", "provider_name": "network_policy_rule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_rule/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_network_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_network_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_rule/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_rule/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_policy_rule/resource.tf`; digest `sha256:0feaf66dfd5662ae652b573d32c61f50aa8eb63e0f1a54a1e0d029c6ad96ace9`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_rule/examples/)
- [xcsh_network_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy_rule/)
