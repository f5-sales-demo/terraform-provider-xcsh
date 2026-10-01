---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cloud_user_account."
xcsh_docs: {"aliases": [], "body_bytes": 1149, "body_sha256": "sha256:2975d16951932df09420559ff1305802c13d94db5082f58aabd9e06bd4bb2a01", "canonical_id": "xcsh-docs:data-sources:cloud_user_account:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cloud_user_account:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:4bd66f34c8654fcf2e99224154cac4393ee6bb933699a10ecdb50c1d8d402e99", "source_path": "examples/data-sources/xcsh_cloud_user_account/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cloud_user_account:example:data-source", "parent_id": "xcsh-docs:data-sources:cloud_user_account:examples", "path": "docs/guides/data-sources--cloud_user_account--example--data-source.md", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_user_account/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_cloud_user_account.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md)
- [Examples](data-sources--cloud_user_account--examples.md)
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

- [Examples](data-sources--cloud_user_account--examples.md)
- [xcsh_cloud_user_account](../data-sources/cloud_user_account.md)
