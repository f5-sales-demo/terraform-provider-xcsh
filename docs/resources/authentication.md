---
page_title: "xcsh_authentication"
subcategory: ""
description: "xcsh_authentication for xcsh_authentication."
xcsh_docs: {"aliases": [], "body_bytes": 1334, "body_sha256": "sha256:d4419712d4a208463a921ad8b20cee40e1c28d7b34604dfad6bff713f51ae0ba", "canonical_id": "xcsh-docs:resources:authentication:fundamentals", "child_ids": ["xcsh-docs:resources:authentication:reference", "xcsh-docs:resources:authentication:examples", "xcsh-docs:resources:authentication:import", "xcsh-docs:resources:authentication:timeouts"], "collection_id": "xcsh-docs:resources:authentication:collection", "completeness": "complete", "id": "xcsh-docs:resources:authentication:fundamentals", "parent_id": null, "path": "docs/resources/authentication.md", "provider_name": "authentication", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authentication/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_authentication for xcsh_authentication.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authenticationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_authentication

Breadcrumbs:

- xcsh_authentication

Manages a Authentication resource in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Authentication Resource Example
# Manages a Authentication resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Authentication configuration
resource "xcsh_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--authentication--reference.md)
- [Examples](../guides/resources--authentication--examples.md)
- [Import](../guides/resources--authentication--import.md)
- [Timeouts](../guides/resources--authentication--timeouts.md)
