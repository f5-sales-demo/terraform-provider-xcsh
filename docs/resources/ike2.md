---
page_title: "xcsh_ike2"
subcategory: ""
description: "xcsh_ike2 for xcsh_ike2."
xcsh_docs: {"aliases": [], "body_bytes": 1303, "body_sha256": "sha256:340f208a62ecaf1f118cf71cd7c7ead0240fb7234f65aa1d40cc17d13db2509f", "canonical_id": "xcsh-docs:resources:ike2:fundamentals", "child_ids": ["xcsh-docs:resources:ike2:reference", "xcsh-docs:resources:ike2:examples", "xcsh-docs:resources:ike2:import", "xcsh-docs:resources:ike2:timeouts"], "collection_id": "xcsh-docs:resources:ike2:collection", "completeness": "complete", "id": "xcsh-docs:resources:ike2:fundamentals", "parent_id": null, "path": "docs/resources/ike2.md", "provider_name": "ike2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike2/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_ike2 for xcsh_ike2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_ike2

Breadcrumbs:

- xcsh_ike2

Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Ike2 Resource Example
# Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Ike2 configuration
resource "xcsh_ike2" "example" {
  name      = "example-ike2"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--ike2--reference.md)
- [Examples](../guides/resources--ike2--examples.md)
- [Import](../guides/resources--ike2--import.md)
- [Timeouts](../guides/resources--ike2--timeouts.md)
