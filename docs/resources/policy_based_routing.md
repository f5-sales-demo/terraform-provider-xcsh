---
page_title: "xcsh_policy_based_routing"
subcategory: ""
description: "xcsh_policy_based_routing for xcsh_policy_based_routing."
xcsh_docs: {"aliases": [], "body_bytes": 1426, "body_sha256": "sha256:5e3749ce9bcc838e047aa305f19cc539e67bfdc579caf163a8ffd54bd3ed1a99", "canonical_id": "xcsh-docs:resources:policy_based_routing:fundamentals", "child_ids": ["xcsh-docs:resources:policy_based_routing:reference", "xcsh-docs:resources:policy_based_routing:examples", "xcsh-docs:resources:policy_based_routing:import", "xcsh-docs:resources:policy_based_routing:timeouts"], "collection_id": "xcsh-docs:resources:policy_based_routing:collection", "completeness": "complete", "id": "xcsh-docs:resources:policy_based_routing:fundamentals", "parent_id": null, "path": "docs/resources/policy_based_routing.md", "provider_name": "policy_based_routing", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policy_based_routing/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_policy_based_routing for xcsh_policy_based_routing.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policy_based_routingCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

- [Property reference](../guides/resources--policy_based_routing--reference.md)
- [Examples](../guides/resources--policy_based_routing--examples.md)
- [Import](../guides/resources--policy_based_routing--import.md)
- [Timeouts](../guides/resources--policy_based_routing--timeouts.md)
