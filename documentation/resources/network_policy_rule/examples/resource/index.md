---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_network_policy_rule."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1354, "body_sha256": "sha256:a9902c6363642c5e61fa7005b14d332ae8a6f1c796036883e986fbfc5b3b2ca5", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:network_policy_rule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0feaf66dfd5662ae652b573d32c61f50aa8eb63e0f1a54a1e0d029c6ad96ace9", "source_path": "examples/resources/xcsh_network_policy_rule/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:network_policy_rule:example:resource", "parent_id": "xcsh-docs:resources:network_policy_rule:examples", "path": "documentation/resources/network_policy_rule/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "network_policy_rule", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0230123322120131-3212132211311011-3332122203002302-1322201201212110-0110101331201331-1002301220103202-3221012201100012-0323020302322320", "registry_path": "docs/guides/resources--network_policy_rule--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy_rule/examples/resource/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Resource for xcsh_network_policy_rule.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["network_policy_ruleCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
