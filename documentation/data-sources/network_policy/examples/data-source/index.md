---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_network_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1071, "body_sha256": "sha256:686e7451d8e80c431adeed8ddd2f2d5720c3b00e686d3d5f7716dfc2ad425303", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:6c10eea9b1a5304339370438b5c012d3317ee2ecb95c539f946182ee126b8c01", "source_path": "examples/data-sources/xcsh_network_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:network_policy:examples", "path": "documentation/data-sources/network_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2132333112011321-3110031010120001-1012210013132100-1300112110232020-0212011303021021-1331012033020110-3220200233013013-1121211120210330", "registry_path": "docs/guides/data-sources--network_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_network_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["network_policyCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
