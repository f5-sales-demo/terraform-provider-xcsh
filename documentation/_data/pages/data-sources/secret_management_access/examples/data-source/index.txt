---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_secret_management_access."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1169, "body_sha256": "sha256:b394d5353fd673b3e32010c2235ffd383f27fb4e7585bc01fced5519a159a702", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:secret_management_access:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ed6c81d9e9def7ced3f618a92fe6932b7593d44a1f95e72e70fe1cd4aab5ace2", "source_path": "examples/data-sources/xcsh_secret_management_access/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:secret_management_access:example:data-source", "parent_id": "xcsh-docs:data-sources:secret_management_access:examples", "path": "documentation/data-sources/secret_management_access/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "secret_management_access", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "data-sources", "registry_anchor": "canonical-3302210133323002-1000031112203132-2012120131121021-3210112320101210-3003101021323310-3212230000320000-2320123132012231-0101100201333230", "registry_path": "docs/guides/data-sources--secret_management_access--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/secret_management_access/examples/data-source/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Data source for xcsh_secret_management_access.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["secret_management_accessCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
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
