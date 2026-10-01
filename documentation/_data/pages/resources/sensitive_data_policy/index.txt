---
page_title: "xcsh_sensitive_data_policy"
subcategory: "Security"
description: "xcsh_sensitive_data_policy for xcsh_sensitive_data_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1751, "body_sha256": "sha256:12e3ada6dc89c237711bebbf6fa1b9e73a21bc29e595b7c9040cc7e43bf7df32", "child_ids": ["xcsh-docs:resources:sensitive_data_policy:reference", "xcsh-docs:resources:sensitive_data_policy:examples", "xcsh-docs:resources:sensitive_data_policy:import", "xcsh-docs:resources:sensitive_data_policy:timeouts"], "collection_id": "xcsh-docs:resources:sensitive_data_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:sensitive_data_policy:fundamentals", "parent_id": null, "path": "documentation/resources/sensitive_data_policy/index.md", "provider_name": "sensitive_data_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/sensitive_data_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_sensitive_data_policy for xcsh_sensitive_data_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["sensitive_data_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/sensitive_data_policy/lifecycle/timeouts/)
