---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_authorization_server."
xcsh_docs: {"aliases": [], "body_bytes": 1078, "body_sha256": "sha256:3262f6fa5b534cc746b4dbc1ef7bbca3e12f882e942087f2dd6039b8a53c9d39", "canonical_id": "xcsh-docs:data-sources:authorization_server:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:authorization_server:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:268fe1444519cbb845d8fcb9ea236d26293e2aeafb324fad41e08e7bc571ffcb", "source_path": "examples/data-sources/xcsh_authorization_server/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:authorization_server:example:data-source", "parent_id": "xcsh-docs:data-sources:authorization_server:examples", "path": "docs/guides/data-sources--authorization_server--example--data-source.md", "provider_name": "authorization_server", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/authorization_server/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_authorization_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["authorization_serverCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_authorization_server](../data-sources/authorization_server.md)
- [Examples](data-sources--authorization_server--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_authorization_server/data-source.tf`; digest `sha256:268fe1444519cbb845d8fcb9ea236d26293e2aeafb324fad41e08e7bc571ffcb`.

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

## Next pages

- [Examples](data-sources--authorization_server--examples.md)
- [xcsh_authorization_server](../data-sources/authorization_server.md)
