---
page_title: "xcsh_policy_based_routing"
subcategory: ""
description: "Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing create specification. configuration."
xcsh_docs: {"aliases": ["policy based routing"], "body_bytes": 1714, "body_sha256": "sha256:05f8a7873c58494e82f52872842dc136906b754e4d4c17e52c5a12fbc14e5b66", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:policy_based_routing:reference", "xcsh-docs:resources:policy_based_routing:examples", "xcsh-docs:resources:policy_based_routing:import", "xcsh-docs:resources:policy_based_routing:timeouts"], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/policy_based_routing/index.md", "product": "distributed-cloud", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-1233013013303222-0100132003201320-0122321122231030-3111323001332220-3310313213102020-3302101031100230-3012131203102032-1111300301011102", "registry_path": "docs/resources/policy_based_routing.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing create specification. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
# PolicyBasedRouting Resource Example
# Manages a Policy Based Routing resource in F5 Distributed Cloud for network policy based routing create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic PolicyBasedRouting configuration
resource "xcsh_policy_based_routing" "example" {
  name      = "example-policy-based-routing"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policy_based_routing/lifecycle/timeouts/)
