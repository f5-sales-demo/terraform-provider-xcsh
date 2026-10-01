---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_authorization_server."
xcsh_docs: {"aliases": [], "body_bytes": 1209, "body_sha256": "sha256:a630d198e4491db87bcd01743db2bbeb25e9fc484bbbca137cfc887e87054b32", "canonical_id": "xcsh-docs:resources:authorization_server:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:authorization_server:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:bbb05ccd40cd71dbeef4a9b6ed31c1cec2684fa158b0357f99c63ef6b31e6af5", "source_path": "examples/resources/xcsh_authorization_server/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:authorization_server:example:resource", "parent_id": "xcsh-docs:resources:authorization_server:examples", "path": "docs/guides/resources--authorization_server--example--resource.md", "provider_name": "authorization_server", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/authorization_server/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_authorization_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authorization_serverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_authorization_server](../resources/authorization_server.md)
- [Examples](resources--authorization_server--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_authorization_server/resource.tf`; digest `sha256:bbb05ccd40cd71dbeef4a9b6ed31c1cec2684fa158b0357f99c63ef6b31e6af5`.

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

## Next pages

- [Examples](resources--authorization_server--examples.md)
- [xcsh_authorization_server](../resources/authorization_server.md)
