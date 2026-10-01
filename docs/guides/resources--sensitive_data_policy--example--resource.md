---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_sensitive_data_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1189, "body_sha256": "sha256:61851dc99ae81d3ad7ed4c0af7e286e2218eb79f4665cb7767541b25dd9cf511", "canonical_id": "xcsh-docs:resources:sensitive_data_policy:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:sensitive_data_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1472df2a4b2a381ee1201907e79b19112e2352a74e50b30a503511f47b7c7e1a", "source_path": "examples/resources/xcsh_sensitive_data_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:sensitive_data_policy:example:resource", "parent_id": "xcsh-docs:resources:sensitive_data_policy:examples", "path": "docs/guides/resources--sensitive_data_policy--example--resource.md", "provider_name": "sensitive_data_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/sensitive_data_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_sensitive_data_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["sensitive_data_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md)
- [Examples](resources--sensitive_data_policy--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_sensitive_data_policy/resource.tf`; digest `sha256:1472df2a4b2a381ee1201907e79b19112e2352a74e50b30a503511f47b7c7e1a`.

```terraform
# SensitiveDataPolicy Resource Example
# Manages sensitive_data_policy creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SensitiveDataPolicy configuration
resource "xcsh_sensitive_data_policy" "example" {
  name      = "example-sensitive-data-policy"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--sensitive_data_policy--examples.md)
- [xcsh_sensitive_data_policy](../resources/sensitive_data_policy.md)
