---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_secret_management_access."
xcsh_docs: {"aliases": [], "body_bytes": 1433, "body_sha256": "sha256:2a42f3360bdf912336624541dc196cfeb0e05b1bb0e2ca2ac82a72f30d6a085f", "child_ids": [], "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ed6c81d9e9def7ced3f618a92fe6932b7593d44a1f95e72e70fe1cd4aab5ace2", "source_path": "examples/data-sources/xcsh_secret_management_access/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:secret_management_access:example:data-source", "parent_id": "xcsh-docs:data-sources:secret_management_access:examples", "path": "documentation/data-sources/secret_management_access/examples/data-source/index.md", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_secret_management_access.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_secret_management_access/data-source.tf`; digest `sha256:ed6c81d9e9def7ced3f618a92fe6932b7593d44a1f95e72e70fe1cd4aab5ace2`.

```terraform
# SecretManagementAccess Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecretManagementAccess by name
data "xcsh_secret_management_access" "example" {
  name      = "example-secret-management-access"
  namespace = "staging"
}

output "secret_management_access_id" {
  value = data.xcsh_secret_management_access.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/examples/)
- [xcsh_secret_management_access](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/secret_management_access/)
