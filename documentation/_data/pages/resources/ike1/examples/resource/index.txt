---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_ike1."
xcsh_docs: {"aliases": [], "body_bytes": 1176, "body_sha256": "sha256:5c746ddec6f64fc94512f9517dedd40caa4ce903f933e6cfccea898a13caba91", "child_ids": [], "collection_id": "xcsh-docs:resources:ike1:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:3f9c75e223b14e98fb8be4f0bccd9b3e08736eec2e962a472124a5092f03950e", "source_path": "examples/resources/xcsh_ike1/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:ike1:example:resource", "parent_id": "xcsh-docs:resources:ike1:examples", "path": "documentation/resources/ike1/examples/resource/index.md", "provider_name": "ike1", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/ike1/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_ike1.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike1CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_ike1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_ike1/resource.tf`; digest `sha256:3f9c75e223b14e98fb8be4f0bccd9b3e08736eec2e962a472124a5092f03950e`.

```terraform
# Ike1 Resource Example
# Manages a Ike1 resource in F5 Distributed Cloud for ike phase1 profile specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Ike1 configuration
resource "xcsh_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/examples/)
- [xcsh_ike1](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/ike1/)
