---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_irule."
xcsh_docs: {"aliases": [], "body_bytes": 1212, "body_sha256": "sha256:1139d326e46cd4f509a283f40fd24b3e2f7bf4830df65ae9b36b00258bee300a", "child_ids": [], "collection_id": "xcsh-docs:resources:irule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:42cb09305e24144e92e6cea8fc9300e3823e0b03a09b749cb4b4a6120a2dca4b", "source_path": "examples/resources/xcsh_irule/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:irule:example:resource", "parent_id": "xcsh-docs:resources:irule:examples", "path": "documentation/resources/irule/examples/resource/index.md", "provider_name": "irule", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/irule/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_irule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["iruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_irule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_irule/resource.tf`; digest `sha256:42cb09305e24144e92e6cea8fc9300e3823e0b03a09b749cb4b4a6120a2dca4b`.

```terraform
# Irule Resource Example
# Manages iRule in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Irule configuration
resource "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"

  description_spec = "example-value"
  irule            = "example-value"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/examples/)
- [xcsh_irule](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/irule/)
