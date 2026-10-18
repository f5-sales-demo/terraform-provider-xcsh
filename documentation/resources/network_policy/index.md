---
page_title: "xcsh_network_policy"
subcategory: "Security"
description: "Manages new network policy with configured parameters in specified namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["network policy"], "body_bytes": 1654, "body_sha256": "sha256:39308e77bf0d15f1dcf77d69e1e28774eca20e440b3c7eed9ec8906837fa4639", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:resources:network_policy:reference", "xcsh-docs:resources:network_policy:examples", "xcsh-docs:resources:network_policy:import", "xcsh-docs:resources:network_policy:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_policy:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/network_policy/index.md", "product": "distributed-cloud", "provider_name": "network_policy", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3303202210130023-0022113313122313-1112200011302303-3013332022223332-1112311323201031-3001211201221121-2232321123000331-0000002213000112", "registry_path": "docs/resources/network_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_policy/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Manages new network policy with configured parameters in specified namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["network_policyCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_policy/lifecycle/timeouts/)
