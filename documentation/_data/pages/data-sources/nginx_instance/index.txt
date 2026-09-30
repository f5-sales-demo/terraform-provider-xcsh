---
page_title: "xcsh_nginx_instance"
subcategory: ""
description: "xcsh_nginx_instance for xcsh_nginx_instance."
xcsh_docs: {"aliases": [], "body_bytes": 1302, "body_sha256": "sha256:8475b77b7634ee804a5b4ad53f009a6be81f9ed7ed01b665f25a84e3895d4ebc", "child_ids": ["xcsh-docs:data-sources:nginx_instance:reference", "xcsh-docs:data-sources:nginx_instance:examples"], "collection_id": "xcsh-docs:data-sources:nginx_instance:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:nginx_instance:fundamentals", "parent_id": null, "path": "documentation/data-sources/nginx_instance/index.md", "provider_name": "nginx_instance", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_instance/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_nginx_instance for xcsh_nginx_instance.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_nginx_instance

Breadcrumbs:

- xcsh_nginx_instance

Manages a Nginx Instance resource in F5 Distributed Cloud for get nginx instance configuration.
configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NginxInstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxInstance by name
data "xcsh_nginx_instance" "example" {
  name      = "example-nginx-instance"
  namespace = "staging"
}

output "nginx_instance_id" {
  value = data.xcsh_nginx_instance.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/examples/)
