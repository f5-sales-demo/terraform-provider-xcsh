---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_secret_management_access."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1433, "body_sha256": "sha256:2a42f3360bdf912336624541dc196cfeb0e05b1bb0e2ca2ac82a72f30d6a085f", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ed6c81d9e9def7ced3f618a92fe6932b7593d44a1f95e72e70fe1cd4aab5ace2", "source_path": "examples/data-sources/xcsh_secret_management_access/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:secret_management_access:example:data-source", "parent_id": "xcsh-docs:data-sources:secret_management_access:examples", "path": "documentation/data-sources/secret_management_access/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3302210133323002-1000031112203132-2012120131121021-3210112320101210-3003101021323310-3212230000320000-2320123132012231-0101100201333230", "registry_path": "docs/guides/data-sources--secret_management_access--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_secret_management_access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
