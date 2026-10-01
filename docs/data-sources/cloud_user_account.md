---
page_title: "xcsh_cloud_user_account"
subcategory: ""
description: "xcsh_cloud_user_account for xcsh_cloud_user_account."
xcsh_docs: {"aliases": [], "body_bytes": 1349, "body_sha256": "sha256:9d22a985b0120ebefe92667723b9e1ce29531613c147f2b357593ac388a7a340", "canonical_id": "xcsh-docs:data-sources:cloud_user_account:fundamentals", "child_ids": ["xcsh-docs:data-sources:cloud_user_account:reference", "xcsh-docs:data-sources:cloud_user_account:examples"], "collection_id": "xcsh-docs:data-sources:cloud_user_account:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_user_account:fundamentals", "parent_id": null, "path": "docs/data-sources/cloud_user_account.md", "provider_name": "cloud_user_account", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_user_account/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_cloud_user_account for xcsh_cloud_user_account.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_user_accountCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_cloud_user_account

Breadcrumbs:

- xcsh_cloud_user_account

Manages a Cloud User Account resource in F5 Distributed Cloud for cloud user account object create
specifications. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--cloud_user_account--reference.md)
- [Examples](../guides/data-sources--cloud_user_account--examples.md)
