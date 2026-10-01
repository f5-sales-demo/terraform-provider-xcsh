---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_nginx_instance."
xcsh_docs: {"aliases": [], "body_bytes": 1305, "body_sha256": "sha256:b97c1a43051f886b029456501711576e9929d7ce406c853bda9b0025c0477f63", "child_ids": [], "collection_id": "xcsh-docs:data-sources:nginx_instance:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:772405375dce9fbbb8eabe6970ee7d9b3a944178189b317c587b8f2983eaf8fd", "source_path": "examples/data-sources/xcsh_nginx_instance/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:nginx_instance:example:data-source", "parent_id": "xcsh-docs:data-sources:nginx_instance:examples", "path": "documentation/data-sources/nginx_instance/examples/data-source/index.md", "provider_name": "nginx_instance", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_instance/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_nginx_instance.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nginx_instance/data-source.tf`; digest `sha256:772405375dce9fbbb8eabe6970ee7d9b3a944178189b317c587b8f2983eaf8fd`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/examples/)
- [xcsh_nginx_instance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_instance/)
