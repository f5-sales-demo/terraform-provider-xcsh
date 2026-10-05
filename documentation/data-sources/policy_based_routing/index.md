---
page_title: "xcsh_policy_based_routing"
subcategory: ""
description: "Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing create specification. configuration."
xcsh_docs: {"aliases": ["policy based routing"], "body_bytes": 1458, "body_sha256": "sha256:325df147fd56e88437d40aa255c89d7e9958faa18e1a2f8174eb6a8b45149783", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:policy_based_routing:reference", "xcsh-docs:data-sources:policy_based_routing:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policy_based_routing:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/policy_based_routing/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2133102010301311-1122300333110212-0113100300221123-0232113133011013-1302230200202323-2120230033020310-2101300103030113-3212011122230202", "registry_path": "docs/data-sources/policy_based_routing.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policy_based_routing/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing create specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_policy_based_routing

Breadcrumbs:

- xcsh_policy_based_routing

Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing
create specification. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# PolicyBasedRouting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing PolicyBasedRouting by name
data "xcsh_policy_based_routing" "example" {
  name      = "example-policy-based-routing"
  namespace = "staging"
}

output "policy_based_routing_id" {
  value = data.xcsh_policy_based_routing.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policy_based_routing/examples/)
