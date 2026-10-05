---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_network_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1305, "body_sha256": "sha256:6abfc950aaad7ea890f34a620445df26ce7074372cdd67810257b6a7a4cc3671", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6c10eea9b1a5304339370438b5c012d3317ee2ecb95c539f946182ee126b8c01", "source_path": "examples/data-sources/xcsh_network_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:network_policy:examples", "path": "documentation/data-sources/network_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2132333112011321-3110031010120001-1012210013132100-1300112110232020-0212011303021021-1331012033020110-3220200233013013-1121211120210330", "registry_path": "docs/guides/data-sources--network_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_network_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_policy/data-source.tf`; digest `sha256:6c10eea9b1a5304339370438b5c012d3317ee2ecb95c539f946182ee126b8c01`.

```terraform
# NetworkPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicy by name
data "xcsh_network_policy" "example" {
  name      = "example-network-policy"
  namespace = "staging"
}

output "network_policy_id" {
  value = data.xcsh_network_policy.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/examples/)
- [xcsh_network_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy/)
