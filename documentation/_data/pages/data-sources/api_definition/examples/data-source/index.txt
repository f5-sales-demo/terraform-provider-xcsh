---
page_title: "Data source"
subcategory: "API Management"
description: "Data source for xcsh_api_definition."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1071, "body_sha256": "sha256:c6da081e6f27da7c45996856167b82c35f2267c3c2b1c2012ef1c3ded3069608", "capabilities": ["api-management"], "category": "api-management", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:api_definition:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:00117931ab83e7521384d3c585bd071aa24022def862d24108a8273097450108", "source_path": "examples/data-sources/xcsh_api_definition/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:api_definition:example:data-source", "parent_id": "xcsh-docs:data-sources:api_definition:examples", "path": "documentation/data-sources/api_definition/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "api_definition", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1011110003001222-0002133011213111-2121022013100123-1332320322213012-3013322122122210-0020023020120223-0230020032300213-1212100213312313", "registry_path": "docs/guides/data-sources--api_definition--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_definition/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_api_definition.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["api_definitionCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
