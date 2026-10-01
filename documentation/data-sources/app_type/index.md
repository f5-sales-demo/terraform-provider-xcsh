---
page_title: "xcsh_app_type"
subcategory: ""
description: "xcsh_app_type for xcsh_app_type."
xcsh_docs: {"aliases": [], "body_bytes": 1310, "body_sha256": "sha256:d256fad9cde768e8acff56c9d3ba72e82fe0e7f8b923016b452a98f8a9aa7c49", "child_ids": ["xcsh-docs:data-sources:app_type:reference", "xcsh-docs:data-sources:app_type:examples"], "collection_id": "xcsh-docs:data-sources:app_type:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:app_type:fundamentals", "parent_id": null, "path": "documentation/data-sources/app_type/index.md", "provider_name": "app_type", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/app_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_app_type for xcsh_app_type.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_app_type

Breadcrumbs:

- xcsh_app_type

Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppType Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppType by name
data "xcsh_app_type" "example" {
  name      = "example-app-type"
  namespace = "staging"
}

output "app_type_id" {
  value = data.xcsh_app_type.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/app_type/examples/)
