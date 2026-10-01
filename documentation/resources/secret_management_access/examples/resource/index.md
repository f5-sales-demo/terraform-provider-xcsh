---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 1462, "body_sha256": "sha256:a521379fdd97b52868b817462a5d019e7fffc9bd64be909d9a0eefa61639024e", "child_ids": [], "collection_id": "xcsh-docs:resources:secret_management_access:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:49f920527939852314c086a0e57abc992c317f684e7845daaf75a5df25aa0d81", "source_path": "examples/resources/xcsh_secret_management_access/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:secret_management_access:example:resource", "parent_id": "xcsh-docs:resources:secret_management_access:examples", "path": "documentation/resources/secret_management_access/examples/resource/index.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/secret_management_access/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_secret_management_access/resource.tf`; digest `sha256:49f920527939852314c086a0e57abc992c317f684e7845daaf75a5df25aa0d81`.

```terraform
# SecretManagementAccess Resource Example
# Manages secret_management_access creates a new object in storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecretManagementAccess configuration
resource "xcsh_secret_management_access" "example" {
  name      = "example-secret-management-access"
  namespace = "staging"

  provider_name = "example-value"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/examples/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/secret_management_access/)
