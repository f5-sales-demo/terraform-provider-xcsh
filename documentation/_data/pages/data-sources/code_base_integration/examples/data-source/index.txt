---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_code_base_integration."
xcsh_docs: {"aliases": [], "body_bytes": 1394, "body_sha256": "sha256:caf135b29512245bb0fc8f93bb8238e68d1a00c50015dfe871608ecad7db9e10", "child_ids": [], "collection_id": "xcsh-docs:data-sources:code_base_integration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:170c39b82f55f5ee2a4de8f66b829b7a3cbefe8bbffc6b09199912a373cb3b88", "source_path": "examples/data-sources/xcsh_code_base_integration/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:code_base_integration:example:data-source", "parent_id": "xcsh-docs:data-sources:code_base_integration:examples", "path": "documentation/data-sources/code_base_integration/examples/data-source/index.md", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/code_base_integration/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_code_base_integration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_code_base_integration/data-source.tf`; digest `sha256:170c39b82f55f5ee2a4de8f66b829b7a3cbefe8bbffc6b09199912a373cb3b88`.

```terraform
# CodeBaseIntegration Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CodeBaseIntegration by name
data "xcsh_code_base_integration" "example" {
  name      = "example-code-base-integration"
  namespace = "staging"
}

output "code_base_integration_id" {
  value = data.xcsh_code_base_integration.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/examples/)
- [xcsh_code_base_integration](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/code_base_integration/)
