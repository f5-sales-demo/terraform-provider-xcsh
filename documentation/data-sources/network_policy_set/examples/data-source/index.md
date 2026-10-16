---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_policy_set."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1109, "body_sha256": "sha256:4c3c1e17d6c9aecec889a5ba876f904f44593fd052cc28fe565098e5959396cf", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_set:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:292e820cf31aaf1ca4db02cc9e7027c764e59f12b6fedea7ebfc71f60117c57e", "source_path": "examples/data-sources/xcsh_network_policy_set/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_policy_set:example:data-source", "parent_id": "xcsh-docs:data-sources:network_policy_set:examples", "path": "documentation/data-sources/network_policy_set/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_policy_set", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-1110122112330003-0202010302130112-1032221102200032-0032231131301032-2011230313113113-1132200310203222-1113103100301323-3200111212221102", "registry_path": "docs/guides/data-sources--network_policy_set--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_set/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Data source for xcsh_network_policy_set.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": [], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_policy_set](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_set/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_policy_set/data-source.tf`; digest `sha256:292e820cf31aaf1ca4db02cc9e7027c764e59f12b6fedea7ebfc71f60117c57e`.

```terraform
# NetworkPolicySet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicySet by name
data "xcsh_network_policy_set" "example" {
  name      = "example-network-policy-set"
  namespace = "staging"
}

output "network_policy_set_id" {
  value = data.xcsh_network_policy_set.example.id
}
```
