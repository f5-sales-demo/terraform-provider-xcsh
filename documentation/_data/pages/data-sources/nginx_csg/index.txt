---
page_title: "xcsh_nginx_csg"
subcategory: ""
description: "xcsh_nginx_csg for xcsh_nginx_csg."
xcsh_docs: {"aliases": [], "body_bytes": 1341, "body_sha256": "sha256:52af73d2ea8ce980ba5432dace47aaad18015192c52cd9083ea6396d6556abb5", "child_ids": ["xcsh-docs:data-sources:nginx_csg:reference", "xcsh-docs:data-sources:nginx_csg:examples"], "collection_id": "xcsh-docs:data-sources:nginx_csg:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_csg:fundamentals", "parent_id": null, "path": "documentation/data-sources/nginx_csg/index.md", "provider_name": "nginx_csg", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_csg/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_nginx_csg for xcsh_nginx_csg.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_nginx_csg

Breadcrumbs:

- xcsh_nginx_csg

Manages a Nginx Csg resource in F5 Distributed Cloud for get nginx csg configuration. configuration.
(read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NginxCsg Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxCsg by name
data "xcsh_nginx_csg" "example" {
  name      = "example-nginx-csg"
  namespace = "staging"
}

output "nginx_csg_id" {
  value = data.xcsh_nginx_csg.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_csg/examples/)
