---
page_title: "xcsh_sensitive_data_policy"
subcategory: "Security"
description: "xcsh_sensitive_data_policy for xcsh_sensitive_data_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1407, "body_sha256": "sha256:a8a92708bf976ea10e1abe186c75a7bf0652e4b0217d1b7e252ecb9633d62c48", "canonical_id": "xcsh-docs:data-sources:sensitive_data_policy:fundamentals", "child_ids": ["xcsh-docs:data-sources:sensitive_data_policy:reference", "xcsh-docs:data-sources:sensitive_data_policy:examples"], "collection_id": "xcsh-docs:data-sources:sensitive_data_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:sensitive_data_policy:fundamentals", "parent_id": null, "path": "docs/data-sources/sensitive_data_policy.md", "provider_name": "sensitive_data_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/sensitive_data_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_sensitive_data_policy for xcsh_sensitive_data_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["sensitive_data_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_sensitive_data_policy

Breadcrumbs:

- xcsh_sensitive_data_policy

Manages sensitive\_data\_policy creates a new object in the storage backend for metadata.namespace
in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--sensitive_data_policy--reference.md)
- [Examples](../guides/data-sources--sensitive_data_policy--examples.md)
