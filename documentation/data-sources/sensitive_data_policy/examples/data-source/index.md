---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_sensitive_data_policy."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1394, "body_sha256": "sha256:b3f9bdb20857b30744f463cf2952f41fcbf52417f94d9547c2e9b3edcde80290", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:sensitive_data_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:cfa788d1ad2d015969f47a23be03161a048878a194ca7cc1ede99b43ed21acb8", "source_path": "examples/data-sources/xcsh_sensitive_data_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:sensitive_data_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:sensitive_data_policy:examples", "path": "documentation/data-sources/sensitive_data_policy/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "sensitive_data_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3000113000010210-3110121213122023-0330003221011312-2213310130020331-2301323102013130-0031021032220203-2302310032030012-3023202303013101", "registry_path": "docs/guides/data-sources--sensitive_data_policy--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/sensitive_data_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_sensitive_data_policy.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["sensitive_data_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_sensitive_data_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/sensitive_data_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/sensitive_data_policy/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_sensitive_data_policy/data-source.tf`; digest `sha256:cfa788d1ad2d015969f47a23be03161a048878a194ca7cc1ede99b43ed21acb8`.

```terraform
# SensitiveDataPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SensitiveDataPolicy by name
data "xcsh_sensitive_data_policy" "example" {
  name      = "example-sensitive-data-policy"
  namespace = "staging"
}

output "sensitive_data_policy_id" {
  value = data.xcsh_sensitive_data_policy.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/sensitive_data_policy/examples/)
- [xcsh_sensitive_data_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/sensitive_data_policy/)
