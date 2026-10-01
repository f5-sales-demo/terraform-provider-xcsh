---
page_title: "Data source"
subcategory: "API Management"
description: "Data source for xcsh_api_definition."
xcsh_docs: {"aliases": [], "body_bytes": 1305, "body_sha256": "sha256:a8f4e5d4f905b5489279e46b28138979e9d0bdcb101adee35a06f7ab266643fe", "child_ids": [], "collection_id": "xcsh-docs:data-sources:api_definition:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:00117931ab83e7521384d3c585bd071aa24022def862d24108a8273097450108", "source_path": "examples/data-sources/xcsh_api_definition/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:api_definition:example:data-source", "parent_id": "xcsh-docs:data-sources:api_definition:examples", "path": "documentation/data-sources/api_definition/examples/data-source/index.md", "provider_name": "api_definition", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_definition/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_api_definition.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_definitionCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_api_definition/data-source.tf`; digest `sha256:00117931ab83e7521384d3c585bd071aa24022def862d24108a8273097450108`.

```terraform
# APIDefinition Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APIDefinition by name
data "xcsh_api_definition" "example" {
  name      = "example-api-definition"
  namespace = "staging"
}

output "api_definition_id" {
  value = data.xcsh_api_definition.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/examples/)
- [xcsh_api_definition](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/api_definition/)
