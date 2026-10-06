---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_policy_rule."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1119, "body_sha256": "sha256:eba496f6e26fa1574bb3347f7533b7b2f5ea9252916c776de3f6ae5acff5c982", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e28844bbbc2540f4dac23720e7a2c315eba810ea0e139291452a348547d903cb", "source_path": "examples/data-sources/xcsh_network_policy_rule/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_policy_rule:example:data-source", "parent_id": "xcsh-docs:data-sources:network_policy_rule:examples", "path": "documentation/data-sources/network_policy_rule/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_policy_rule", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-1331321020003333-3322022322032012-3220312212110232-3221333312323020-0321220020211031-0311033031013131-0022220031323013-2020330210222203", "registry_path": "docs/guides/data-sources--network_policy_rule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_rule/examples/data-source/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Data source for xcsh_network_policy_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["network_policy_ruleCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
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
