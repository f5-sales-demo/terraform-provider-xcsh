---
page_title: "xcsh_network_policy"
subcategory: "Security"
description: "Manages new network policy with configured parameters in specified namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["network policy"], "body_bytes": 1641, "body_sha256": "sha256:fe5d28f34327d8501db828910bd15b274fd901f67788c0218c74ce43dccb7e04", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_policy:reference", "xcsh-docs:resources:network_policy:examples", "xcsh-docs:resources:network_policy:import", "xcsh-docs:resources:network_policy:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/network_policy/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112", "registry_path": "docs/resources/network_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages new network policy with configured parameters in specified namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["network_policyCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_policy

Breadcrumbs:

- xcsh_network_policy

Manages new network policy with configured parameters in specified namespace in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkPolicy Resource Example
# Manages new network policy with configured parameters in specified namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkPolicy configuration
resource "xcsh_network_policy" "example" {
  name      = "example-network-policy"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/lifecycle/timeouts/)
