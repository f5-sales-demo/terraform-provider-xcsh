---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_code_base_integration."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1394, "body_sha256": "sha256:caf135b29512245bb0fc8f93bb8238e68d1a00c50015dfe871608ecad7db9e10", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:code_base_integration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:170c39b82f55f5ee2a4de8f66b829b7a3cbefe8bbffc6b09199912a373cb3b88", "source_path": "examples/data-sources/xcsh_code_base_integration/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:code_base_integration:example:data-source", "parent_id": "xcsh-docs:data-sources:code_base_integration:examples", "path": "documentation/data-sources/code_base_integration/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "code_base_integration", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2313002030023133-0033012021330030-1333322022112320-2122213022302232-2333110212101030-2103111100231103-0302312010003330-2103332012021212", "registry_path": "docs/guides/data-sources--code_base_integration--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/code_base_integration/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_code_base_integration.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["code_base_integrationCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
