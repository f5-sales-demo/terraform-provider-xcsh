---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_protected_domain."
xcsh_docs: {"aliases": [], "body_bytes": 1101, "body_sha256": "sha256:e3fe5deb5eb8f259a11e65b273f543b4681c07235a0e9984793ec229b58f9379", "canonical_id": "xcsh-docs:resources:protected_domain:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:protected_domain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a41646b64f77c1c10c9bfb255c8aeb3c51bbd0f1f5060b667473b425512ade92", "source_path": "examples/resources/xcsh_protected_domain/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:protected_domain:example:resource", "parent_id": "xcsh-docs:resources:protected_domain:examples", "path": "docs/guides/resources--protected_domain--example--resource.md", "provider_name": "protected_domain", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_domain/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_protected_domain.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_domainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_protected_domain](../resources/protected_domain.md)
- [Examples](resources--protected_domain--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_protected_domain/resource.tf`; digest `sha256:a41646b64f77c1c10c9bfb255c8aeb3c51bbd0f1f5060b667473b425512ade92`.

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

## Next pages

- [Examples](resources--protected_domain--examples.md)
- [xcsh_protected_domain](../resources/protected_domain.md)
