---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_policy_view."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1118, "body_sha256": "sha256:b8ebccd8d58e00fdedb408113648f2815f5e6d9fe650cd1ac74695c7a137ef6d", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:961fb42feeb0126081bc0ae8691b145c93619816626daf020f96266ec095e50a", "source_path": "examples/data-sources/xcsh_network_policy_view/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_policy_view:example:data-source", "parent_id": "xcsh-docs:data-sources:network_policy_view:examples", "path": "documentation/data-sources/network_policy_view/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1121223333211212-3310103032120301-0011323123031111-1330000112112300-0130211020023132-3101102212301313-3001322302212313-3300202112323030", "registry_path": "docs/guides/data-sources--network_policy_view--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/examples/data-source/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Data source for xcsh_network_policy_view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_policy_view/data-source.tf`; digest `sha256:961fb42feeb0126081bc0ae8691b145c93619816626daf020f96266ec095e50a`.

```terraform
# NetworkPolicyView Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkPolicyView by name
data "xcsh_network_policy_view" "example" {
  name      = "example-network-policy-view"
  namespace = "system"
}

output "network_policy_view_id" {
  value = data.xcsh_network_policy_view.example.id
}
```
