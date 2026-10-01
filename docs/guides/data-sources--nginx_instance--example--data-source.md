---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_nginx_instance."
xcsh_docs: {"aliases": [], "body_bytes": 1099, "body_sha256": "sha256:5624f1c7b3c054efc1350ad8645d003eae3021c6ab8ccd56fe05514be2a9862f", "canonical_id": "xcsh-docs:data-sources:nginx_instance:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:nginx_instance:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:772405375dce9fbbb8eabe6970ee7d9b3a944178189b317c587b8f2983eaf8fd", "source_path": "examples/data-sources/xcsh_nginx_instance/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:nginx_instance:example:data-source", "parent_id": "xcsh-docs:data-sources:nginx_instance:examples", "path": "docs/guides/data-sources--nginx_instance--example--data-source.md", "provider_name": "nginx_instance", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_instance/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_nginx_instance.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_nginx_instance](../data-sources/nginx_instance.md)
- [Examples](data-sources--nginx_instance--examples.md)
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

- [Examples](data-sources--nginx_instance--examples.md)
- [xcsh_nginx_instance](../data-sources/nginx_instance.md)
