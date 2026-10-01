---
page_title: "Data source"
subcategory: "Security"
description: "Data source for xcsh_sensitive_data_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1394, "body_sha256": "sha256:b3f9bdb20857b30744f463cf2952f41fcbf52417f94d9547c2e9b3edcde80290", "child_ids": [], "collection_id": "xcsh-docs:data-sources:sensitive_data_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:cfa788d1ad2d015969f47a23be03161a048878a194ca7cc1ede99b43ed21acb8", "source_path": "examples/data-sources/xcsh_sensitive_data_policy/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:sensitive_data_policy:example:data-source", "parent_id": "xcsh-docs:data-sources:sensitive_data_policy:examples", "path": "documentation/data-sources/sensitive_data_policy/examples/data-source/index.md", "provider_name": "sensitive_data_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/sensitive_data_policy/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_sensitive_data_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["sensitive_data_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
