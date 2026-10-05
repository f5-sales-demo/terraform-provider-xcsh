---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_authentication."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1307, "body_sha256": "sha256:9e3aa75ad6a908f486c8323a188bbe3b786f74fd5c8f3da8d40e0270ad83d326", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:authentication:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:fd8b227ee8d771e5d805c51da258f79cea8ff5031e710c318c14f6217d6efd78", "source_path": "examples/data-sources/xcsh_authentication/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:authentication:example:data-source", "parent_id": "xcsh-docs:data-sources:authentication:examples", "path": "documentation/data-sources/authentication/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "authentication", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0123220332103303-1322311321030321-3231001223333000-0323201230010231-0332330330200220-0233121310120303-3123121322311302-3301120023202033", "registry_path": "docs/guides/data-sources--authentication--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authentication/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_authentication.", "tasks": ["authentication", "configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["authenticationCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_authentication/data-source.tf`; digest `sha256:fd8b227ee8d771e5d805c51da258f79cea8ff5031e710c318c14f6217d6efd78`.

```terraform
# Authentication Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Authentication by name
data "xcsh_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"
}

output "authentication_id" {
  value = data.xcsh_authentication.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/examples/)
- [xcsh_authentication](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authentication/)
