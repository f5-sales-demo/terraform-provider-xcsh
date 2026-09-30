---
page_title: "xcsh_user_identification"
subcategory: ""
description: "xcsh_user_identification for xcsh_user_identification."
xcsh_docs: {"aliases": [], "body_bytes": 1253, "body_sha256": "sha256:7e8bda00de410cb0714c414f82b62a61a973f122094e9f359fe1b062468ed56c", "canonical_id": "xcsh-docs:data-sources:user_identification:fundamentals", "child_ids": ["xcsh-docs:data-sources:user_identification:reference", "xcsh-docs:data-sources:user_identification:examples"], "collection_id": "xcsh-docs:data-sources:user_identification:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:user_identification:fundamentals", "parent_id": null, "path": "docs/data-sources/user_identification.md", "provider_name": "user_identification", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/user_identification/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_user_identification for xcsh_user_identification.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["user_identificationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_user_identification

Breadcrumbs:

- xcsh_user_identification

Manages user\_identification creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UserIdentification Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing UserIdentification by name
data "xcsh_user_identification" "example" {
  name      = "example-user-identification"
  namespace = "staging"
}

output "user_identification_id" {
  value = data.xcsh_user_identification.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--user_identification--reference.md)
- [Examples](../guides/data-sources--user_identification--examples.md)
