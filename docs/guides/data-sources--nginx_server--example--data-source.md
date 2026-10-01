---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_nginx_server."
xcsh_docs: {"aliases": [], "body_bytes": 1073, "body_sha256": "sha256:693fe584fe5cb8212b8a1223fd837bb3d318c78c3e64d8a417afc5a8134a1e50", "canonical_id": "xcsh-docs:data-sources:nginx_server:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:255e0e3089bba8c6998465d72659b209b4c83e33fec880e96e0bccd476d39e6e", "source_path": "examples/data-sources/xcsh_nginx_server/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:nginx_server:example:data-source", "parent_id": "xcsh-docs:data-sources:nginx_server:examples", "path": "docs/guides/data-sources--nginx_server--example--data-source.md", "provider_name": "nginx_server", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_nginx_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_nginx_server](../data-sources/nginx_server.md)
- [Examples](data-sources--nginx_server--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nginx_server/data-source.tf`; digest `sha256:255e0e3089bba8c6998465d72659b209b4c83e33fec880e96e0bccd476d39e6e`.

```terraform
# NginxServer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxServer by name
data "xcsh_nginx_server" "example" {
  name      = "example-nginx-server"
  namespace = "staging"
}

output "nginx_server_id" {
  value = data.xcsh_nginx_server.example.id
}
```

## Next pages

- [Examples](data-sources--nginx_server--examples.md)
- [xcsh_nginx_server](../data-sources/nginx_server.md)
