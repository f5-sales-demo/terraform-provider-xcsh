---
page_title: "xcsh_fleet"
subcategory: ""
description: "Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["fleet"], "body_bytes": 1576, "body_sha256": "sha256:a0188f2002564fa03add09828f26a8368fb4bef58b2573c21e647c38df1e552d", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:fleet:reference", "xcsh-docs:resources:fleet:examples", "xcsh-docs:resources:fleet:import", "xcsh-docs:resources:fleet:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/fleet/index.md", "product": "distributed-cloud", "provider_name": "fleet", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-1321121000121112-1223302232121013-1011232202230112-2103020112013103-1311332313321220-3321211333313013-1320031110133332-2012010122222131", "registry_path": "docs/resources/fleet.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/index.txt", "spec_pin_digest": "sha256:66dcd7f1752e908b5a58a7638e63662f0daa121544d954167e4fab269122ff4a", "summary": "Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.1", "schema_components": ["fleetCreateRequest"], "target_commit": "bafefa99cbad3ad2e3c2c4fd9c0611a0cbcc36b3"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_fleet

Breadcrumbs:

- xcsh_fleet

Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Fleet Resource Example
# Manages fleet will create a fleet object in 'system' namespace of the user in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Fleet configuration
resource "xcsh_fleet" "example" {
  name      = "example-fleet"
  namespace = "staging"

  fleet_label = "example-value"
}
```

## Root configuration

Required root properties: `fleet_label`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/fleet/lifecycle/timeouts/)
