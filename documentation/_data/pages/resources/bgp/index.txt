---
page_title: "xcsh_bgp"
subcategory: ""
description: "Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with external bgp servers. it is created by users in system namespace. configuration."
xcsh_docs: {"aliases": ["bgp"], "body_bytes": 1611, "body_sha256": "sha256:3ce16d243fb1f34f71964ace4ab624888586a6cc68a05e23655659eb23dd6d81", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:bgp:reference", "xcsh-docs:resources:bgp:examples", "xcsh-docs:resources:bgp:import", "xcsh-docs:resources:bgp:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:bgp:collection", "completeness": "complete", "id": "xcsh-docs:resources:bgp:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/bgp/index.md", "product": "distributed-cloud", "provider_name": "bgp", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3310110230203031-0323100203221200-2312202022333201-2213220303201313-0321312110122102-3032233203123111-2211311331302223-2201000110103200", "registry_path": "docs/resources/bgp.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/bgp/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with external bgp servers. it is created by users in system namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["bgpCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_bgp

Breadcrumbs:

- xcsh_bgp

Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with
external bgp servers. it is created by users in system namespace. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BGP Resource Example
# Manages a BGP resource in F5 Distributed Cloud for bgp object is the configuration for peering with external bgp servers.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic BGP configuration
resource "xcsh_bgp" "example" {
  name      = "example-bgp"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/bgp/lifecycle/timeouts/)
