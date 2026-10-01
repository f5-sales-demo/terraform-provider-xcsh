---
page_title: "xcsh_authorization_server"
subcategory: ""
description: "xcsh_authorization_server for xcsh_authorization_server."
xcsh_docs: {"aliases": [], "body_bytes": 1448, "body_sha256": "sha256:c04a3b02fde7ff652af05b820c8b61fd7116b3c984cbb968e4e3cd6263960b59", "child_ids": ["xcsh-docs:data-sources:authorization_server:reference", "xcsh-docs:data-sources:authorization_server:examples"], "collection_id": "xcsh-docs:data-sources:authorization_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:authorization_server:fundamentals", "parent_id": null, "path": "documentation/data-sources/authorization_server/index.md", "provider_name": "authorization_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authorization_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_authorization_server for xcsh_authorization_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authorization_serverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_authorization_server

Breadcrumbs:

- xcsh_authorization_server

Manages authorization\_server creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AuthorizationServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AuthorizationServer by name
data "xcsh_authorization_server" "example" {
  name      = "example-authorization-server"
  namespace = "staging"
}

output "authorization_server_id" {
  value = data.xcsh_authorization_server.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authorization_server/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/authorization_server/examples/)
