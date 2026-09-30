---
page_title: "xcsh_protected_domain"
subcategory: ""
description: "xcsh_protected_domain for xcsh_protected_domain."
xcsh_docs: {"aliases": [], "body_bytes": 1293, "body_sha256": "sha256:610275969a1e3ca2bdeb5400349c08a7e6238d690dd7ccce23e9acb785f69d65", "canonical_id": "xcsh-docs:resources:protected_domain:fundamentals", "child_ids": ["xcsh-docs:resources:protected_domain:reference", "xcsh-docs:resources:protected_domain:examples", "xcsh-docs:resources:protected_domain:import", "xcsh-docs:resources:protected_domain:timeouts"], "collection_id": "xcsh-docs:resources:protected_domain:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_domain:fundamentals", "parent_id": null, "path": "docs/resources/protected_domain.md", "provider_name": "protected_domain", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_domain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_protected_domain for xcsh_protected_domain.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_domainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_protected_domain

Breadcrumbs:

- xcsh_protected_domain

Manages Domain to protect in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtectedDomain Resource Example
# Manages Domain to protect in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtectedDomain configuration
resource "xcsh_protected_domain" "example" {
  name      = "example-protected-domain"
  namespace = "staging"

  protected_domain = "example.com"
}
```

## Root configuration

Required root properties: `name`, `namespace`, `protected_domain`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--protected_domain--reference.md)
- [Examples](../guides/resources--protected_domain--examples.md)
- [Import](../guides/resources--protected_domain--import.md)
- [Timeouts](../guides/resources--protected_domain--timeouts.md)
