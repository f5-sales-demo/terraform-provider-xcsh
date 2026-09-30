---
page_title: "xcsh_fleet"
subcategory: ""
description: "xcsh_fleet for xcsh_fleet."
xcsh_docs: {"aliases": [], "body_bytes": 1275, "body_sha256": "sha256:7f5b74e920c8e39108d7e3ae0060211f3f9c1dd795bcd5406fe79900cbb4ac91", "canonical_id": "xcsh-docs:resources:fleet:fundamentals", "child_ids": ["xcsh-docs:resources:fleet:reference", "xcsh-docs:resources:fleet:examples", "xcsh-docs:resources:fleet:import", "xcsh-docs:resources:fleet:timeouts"], "collection_id": "xcsh-docs:resources:fleet:collection", "completeness": "complete", "id": "xcsh-docs:resources:fleet:fundamentals", "parent_id": null, "path": "docs/resources/fleet.md", "provider_name": "fleet", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/fleet/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_fleet for xcsh_fleet.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["fleetCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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

## Next pages

- [Property reference](../guides/resources--fleet--reference.md)
- [Examples](../guides/resources--fleet--examples.md)
- [Import](../guides/resources--fleet--import.md)
- [Timeouts](../guides/resources--fleet--timeouts.md)
