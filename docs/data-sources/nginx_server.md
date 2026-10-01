---
page_title: "xcsh_nginx_server"
subcategory: ""
description: "xcsh_nginx_server for xcsh_nginx_server."
xcsh_docs: {"aliases": [], "body_bytes": 1298, "body_sha256": "sha256:4326d65ce92d013b480c8583aadd11c45e63957545f8d16bff19a47b1abb907e", "canonical_id": "xcsh-docs:data-sources:nginx_server:fundamentals", "child_ids": ["xcsh-docs:data-sources:nginx_server:reference", "xcsh-docs:data-sources:nginx_server:examples"], "collection_id": "xcsh-docs:data-sources:nginx_server:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_server:fundamentals", "parent_id": null, "path": "docs/data-sources/nginx_server.md", "provider_name": "nginx_server", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_server/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_nginx_server for xcsh_nginx_server.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_nginx_server

Breadcrumbs:

- xcsh_nginx_server

Manages a Nginx Server resource in F5 Distributed Cloud for get nginx server block configuration.
configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--nginx_server--reference.md)
- [Examples](../guides/data-sources--nginx_server--examples.md)
