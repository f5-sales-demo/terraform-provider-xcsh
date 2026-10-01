---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_route."
xcsh_docs: {"aliases": [], "body_bytes": 938, "body_sha256": "sha256:56030703a13ba28bf7a9c71783dadc39c9cb69e43c171cf53e37b94922c4ea09", "canonical_id": "xcsh-docs:resources:route:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:route:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:02c199de787def8c45388d6ef15d06881bc08c78cf9b74e0ee43c7ac9db177f3", "source_path": "examples/resources/xcsh_route/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:route:example:resource", "parent_id": "xcsh-docs:resources:route:examples", "path": "docs/guides/resources--route--example--resource.md", "provider_name": "route", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/route/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_route.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["routeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_route](../resources/route.md)
- [Examples](resources--route--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_route/resource.tf`; digest `sha256:02c199de787def8c45388d6ef15d06881bc08c78cf9b74e0ee43c7ac9db177f3`.

```terraform
# Route Resource Example
# Manages route object in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Route configuration
resource "xcsh_route" "example" {
  name      = "example-route"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--route--examples.md)
- [xcsh_route](../resources/route.md)
