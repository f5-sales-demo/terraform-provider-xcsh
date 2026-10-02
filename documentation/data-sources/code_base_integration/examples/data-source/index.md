---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_code_base_integration."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1394, "body_sha256": "sha256:caf135b29512245bb0fc8f93bb8238e68d1a00c50015dfe871608ecad7db9e10", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:code_base_integration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:170c39b82f55f5ee2a4de8f66b829b7a3cbefe8bbffc6b09199912a373cb3b88", "source_path": "examples/data-sources/xcsh_code_base_integration/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:code_base_integration:example:data-source", "parent_id": "xcsh-docs:data-sources:code_base_integration:examples", "path": "documentation/data-sources/code_base_integration/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2313002030023133-0033012021330030-1333322022112320-2122213022302232-2333110212101030-2103111100231103-0302312010003330-2103332012021212", "registry_path": "docs/guides/data-sources--code_base_integration--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/code_base_integration/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_code_base_integration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
