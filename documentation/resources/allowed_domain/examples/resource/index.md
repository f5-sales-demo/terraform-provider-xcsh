---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_allowed_domain."
xcsh_docs: {"aliases": [], "body_bytes": 1282, "body_sha256": "sha256:8ea0ffb0a7c9f6e691f9739ecaa09b2a5b3b8cd759bd1a190abdd18a5392786d", "child_ids": [], "collection_id": "xcsh-docs:resources:allowed_domain:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:311173eaa23e9e9afb5856fa5a592866eb29afd044fe81b255b02bc483f4948e", "source_path": "examples/resources/xcsh_allowed_domain/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:allowed_domain:example:resource", "parent_id": "xcsh-docs:resources:allowed_domain:examples", "path": "documentation/resources/allowed_domain/examples/resource/index.md", "provider_name": "allowed_domain", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/allowed_domain/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_allowed_domain.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["allowed_domainCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_allowed_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_allowed_domain/resource.tf`; digest `sha256:311173eaa23e9e9afb5856fa5a592866eb29afd044fe81b255b02bc483f4948e`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/examples/)
- [xcsh_allowed_domain](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/allowed_domain/)
