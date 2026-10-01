---
page_title: "xcsh_protected_domain"
subcategory: ""
description: "xcsh_protected_domain for xcsh_protected_domain."
xcsh_docs: {"aliases": [], "body_bytes": 1252, "body_sha256": "sha256:bc3cd2d0a73c40e8c24453202fc7bda469c7ad2f65c7b43de77fef04906dc61e", "canonical_id": "xcsh-docs:data-sources:protected_domain:fundamentals", "child_ids": ["xcsh-docs:data-sources:protected_domain:reference", "xcsh-docs:data-sources:protected_domain:examples"], "collection_id": "xcsh-docs:data-sources:protected_domain:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:protected_domain:fundamentals", "parent_id": null, "path": "docs/data-sources/protected_domain.md", "provider_name": "protected_domain", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/protected_domain/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_protected_domain for xcsh_protected_domain.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_domainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_protected_domain

Breadcrumbs:

- xcsh_protected_domain

Manages Domain to protect in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtectedDomain Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtectedDomain by name
data "xcsh_protected_domain" "example" {
  name      = "example-protected-domain"
  namespace = "staging"
}

output "protected_domain_id" {
  value = data.xcsh_protected_domain.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--protected_domain--reference.md)
- [Examples](../guides/data-sources--protected_domain--examples.md)
