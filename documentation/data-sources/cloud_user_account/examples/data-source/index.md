---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cloud_user_account."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1355, "body_sha256": "sha256:f8b6ad0728b36527031e32de64eee76d0f62445f7cee4768d3305de228bb7eed", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:cloud_user_account:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4bd66f34c8654fcf2e99224154cac4393ee6bb933699a10ecdb50c1d8d402e99", "source_path": "examples/data-sources/xcsh_cloud_user_account/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cloud_user_account:example:data-source", "parent_id": "xcsh-docs:data-sources:cloud_user_account:examples", "path": "documentation/data-sources/cloud_user_account/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-0320012310310023-2232332310013003-0010130032332200-0303302101222303-2002231232020011-1311312300232133-2201330011113332-3321030110103222", "registry_path": "docs/guides/data-sources--cloud_user_account--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_user_account/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_cloud_user_account.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_user_account/data-source.tf`; digest `sha256:4bd66f34c8654fcf2e99224154cac4393ee6bb933699a10ecdb50c1d8d402e99`.

```terraform
# CloudUserAccount Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudUserAccount by name
data "xcsh_cloud_user_account" "example" {
  name      = "example-cloud-user-account"
  namespace = "staging"
}

output "cloud_user_account_id" {
  value = data.xcsh_cloud_user_account.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/examples/)
- [xcsh_cloud_user_account](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_user_account/)
