---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_irule."
xcsh_docs: {"aliases": [], "body_bytes": 1091, "body_sha256": "sha256:fe2b5e69639991c3d1d68345a496aa7b50663c7178473184c6dd4e26621fda21", "child_ids": [], "collection_id": "xcsh-docs:data-sources:irule:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:567750397ff4df80dac106ded5aa33f7fcb698c11cea8ada22ebaa5fc6727f22", "source_path": "examples/data-sources/xcsh_irule/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:irule:example:data-source", "parent_id": "xcsh-docs:data-sources:irule:examples", "path": "documentation/data-sources/irule/examples/data-source/index.md", "provider_name": "irule", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/irule/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_irule.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["iruleCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_irule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/irule/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/irule/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_irule/data-source.tf`; digest `sha256:567750397ff4df80dac106ded5aa33f7fcb698c11cea8ada22ebaa5fc6727f22`.

```terraform
# Irule Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Irule by name
data "xcsh_irule" "example" {
  name      = "example-irule"
  namespace = "staging"
}

output "irule_id" {
  value = data.xcsh_irule.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/irule/examples/)
- [xcsh_irule](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/irule/)
