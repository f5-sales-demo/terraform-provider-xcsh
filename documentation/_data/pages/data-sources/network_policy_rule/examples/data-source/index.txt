---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_policy_rule."
xcsh_docs: {"aliases": [], "body_bytes": 1368, "body_sha256": "sha256:e6c1e58a55aaaaea7980a3168504c5945ef41d52b369767852343e342a0149e9", "child_ids": [], "collection_id": "xcsh-docs:data-sources:network_policy_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e28844bbbc2540f4dac23720e7a2c315eba810ea0e139291452a348547d903cb", "source_path": "examples/data-sources/xcsh_network_policy_rule/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_policy_rule:example:data-source", "parent_id": "xcsh-docs:data-sources:network_policy_rule:examples", "path": "documentation/data-sources/network_policy_rule/examples/data-source/index.md", "provider_name": "network_policy_rule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_rule/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_network_policy_rule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_policy_ruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_policy_rule/data-source.tf`; digest `sha256:e28844bbbc2540f4dac23720e7a2c315eba810ea0e139291452a348547d903cb`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/examples/)
- [xcsh_network_policy_rule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_rule/)
