---
page_title: "xcsh_ike1"
subcategory: ""
description: "xcsh_ike1 for xcsh_ike1."
xcsh_docs: {"aliases": [], "body_bytes": 1303, "body_sha256": "sha256:ecfc2d3c0bc6830e29dc9856bbfa1c571d1a9b69b2f9a93898208e567b4cda2f", "canonical_id": "xcsh-docs:resources:ike1:fundamentals", "child_ids": ["xcsh-docs:resources:ike1:reference", "xcsh-docs:resources:ike1:examples", "xcsh-docs:resources:ike1:import", "xcsh-docs:resources:ike1:timeouts"], "collection_id": "xcsh-docs:resources:ike1:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike1:fundamentals", "parent_id": null, "path": "docs/resources/ike1.md", "provider_name": "ike1", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike1/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_ike1 for xcsh_ike1.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike1CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_ike1

Breadcrumbs:

- xcsh_ike1

Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile specification. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Ike1 Resource Example
# Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Ike1 configuration
resource "xcsh_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--ike1--reference.md)
- [Examples](../guides/resources--ike1--examples.md)
- [Import](../guides/resources--ike1--import.md)
- [Timeouts](../guides/resources--ike1--timeouts.md)
