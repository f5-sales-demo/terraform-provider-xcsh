---
page_title: "xcsh_authorization_server"
subcategory: ""
description: "xcsh_authorization_server for xcsh_authorization_server."
xcsh_docs: {"aliases": [], "body_bytes": 1559, "body_sha256": "sha256:66d0bea7932d715275e2964ff30e14f9d871204faed6ce7099c08a6eb58fcfb0", "canonical_id": "xcsh-docs:resources:authorization_server:fundamentals", "child_ids": ["xcsh-docs:resources:authorization_server:reference", "xcsh-docs:resources:authorization_server:examples", "xcsh-docs:resources:authorization_server:import", "xcsh-docs:resources:authorization_server:timeouts"], "collection_id": "xcsh-docs:resources:authorization_server:collection", "completeness": "complete", "id": "xcsh-docs:resources:authorization_server:fundamentals", "parent_id": null, "path": "docs/resources/authorization_server.md", "provider_name": "authorization_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authorization_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_authorization_server for xcsh_authorization_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authorization_serverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
# AuthorizationServer Resource Example
# Manages authorization_server creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AuthorizationServer configuration
resource "xcsh_authorization_server" "example" {
  name      = "example-authorization-server"
  namespace = "staging"

  jwks_uri = "example-value"
}
```

## Root configuration

Required root properties: `jwks_uri`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--authorization_server--reference.md)
- [Examples](../guides/resources--authorization_server--examples.md)
- [Import](../guides/resources--authorization_server--import.md)
- [Timeouts](../guides/resources--authorization_server--timeouts.md)
