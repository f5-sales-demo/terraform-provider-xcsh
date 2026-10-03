---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_network_policy_view."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1367, "body_sha256": "sha256:2289c9974e8b0116683b48c41b514213f474b8bfe35532b210e58826e7178725", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_policy_view:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:961fb42feeb0126081bc0ae8691b145c93619816626daf020f96266ec095e50a", "source_path": "examples/data-sources/xcsh_network_policy_view/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:network_policy_view:example:data-source", "parent_id": "xcsh-docs:data-sources:network_policy_view:examples", "path": "documentation/data-sources/network_policy_view/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "network_policy_view", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1121223333211212-3310103032120301-0011323123031111-1330000112112300-0130211020023132-3101102212301313-3001322302212313-3300202112323030", "registry_path": "docs/guides/data-sources--network_policy_view--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_policy_view/examples/data-source/index.txt", "spec_pin_digest": "sha256:62f71ec22260bc99f65753ef4581eb9e0dec1b65c506bb0d53099db73a05e19e", "summary": "Data source for xcsh_network_policy_view.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.2", "schema_components": ["network_policy_viewCreateRequest"], "target_commit": "1a0b5141f4589ffaf7bb696a4369a16ee74ae2ff"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/examples/)
- [xcsh_network_policy_view](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_policy_view/)
