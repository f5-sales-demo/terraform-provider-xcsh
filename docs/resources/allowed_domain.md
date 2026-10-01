---
page_title: "xcsh_allowed_domain"
subcategory: ""
description: "xcsh_allowed_domain for xcsh_allowed_domain."
xcsh_docs: {"aliases": [], "body_bytes": 1364, "body_sha256": "sha256:807f0ccf995260ca7adf5d8ad2640103e8428ff40aad27b8bdf980f438d9dfd9", "canonical_id": "xcsh-docs:resources:allowed_domain:fundamentals", "child_ids": ["xcsh-docs:resources:allowed_domain:reference", "xcsh-docs:resources:allowed_domain:examples", "xcsh-docs:resources:allowed_domain:import", "xcsh-docs:resources:allowed_domain:timeouts"], "collection_id": "xcsh-docs:resources:allowed_domain:collection", "completeness": "complete", "id": "xcsh-docs:resources:allowed_domain:fundamentals", "parent_id": null, "path": "docs/resources/allowed_domain.md", "provider_name": "allowed_domain", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/allowed_domain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_allowed_domain for xcsh_allowed_domain.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["allowed_domainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_allowed_domain

Breadcrumbs:

- xcsh_allowed_domain

Manages allowed domain in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AllowedDomain Resource Example
# Manages allowed domain in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AllowedDomain configuration
resource "xcsh_allowed_domain" "example" {
  name      = "example-allowed-domain"
  namespace = "staging"

  allowed_domain = "example-value"
}
```

## Root configuration

Required root properties: `allowed_domain`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--allowed_domain--reference.md)
- [Examples](../guides/resources--allowed_domain--examples.md)
- [Import](../guides/resources--allowed_domain--import.md)
- [Timeouts](../guides/resources--allowed_domain--timeouts.md)
