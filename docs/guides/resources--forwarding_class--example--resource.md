---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_forwarding_class."
xcsh_docs: {"aliases": [], "body_bytes": 1136, "body_sha256": "sha256:01b607f2ce9fa3624826a8a955e74809943299c2f96b28258550e49c7844a753", "canonical_id": "xcsh-docs:resources:forwarding_class:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:forwarding_class:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:eb45941e172f8b977e57cad7117bbdf4795101f471776f657bf6e2faa09a4be8", "source_path": "examples/resources/xcsh_forwarding_class/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:forwarding_class:example:resource", "parent_id": "xcsh-docs:resources:forwarding_class:examples", "path": "docs/guides/resources--forwarding_class--example--resource.md", "provider_name": "forwarding_class", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/forwarding_class/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_forwarding_class.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["forwarding_classCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_forwarding_class](../resources/forwarding_class.md)
- [Examples](resources--forwarding_class--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_forwarding_class/resource.tf`; digest `sha256:eb45941e172f8b977e57cad7117bbdf4795101f471776f657bf6e2faa09a4be8`.

```terraform
# ForwardingClass Resource Example
# Manages a Forwarding Class resource in F5 Distributed Cloud for forwarding class is created by users in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ForwardingClass configuration
resource "xcsh_forwarding_class" "example" {
  name      = "example-forwarding-class"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--forwarding_class--examples.md)
- [xcsh_forwarding_class](../resources/forwarding_class.md)
