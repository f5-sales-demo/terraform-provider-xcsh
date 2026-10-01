---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_nginx_service_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1420, "body_sha256": "sha256:79fb22be44d37158b2715bd07b0999ef241d3b7a69442b6c60830c2d51cec200", "child_ids": [], "collection_id": "xcsh-docs:data-sources:nginx_service_discovery:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c46b908211689c5888bfaaf70a487544236d409e407922fa907fe86ca0588fa6", "source_path": "examples/data-sources/xcsh_nginx_service_discovery/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:nginx_service_discovery:example:data-source", "parent_id": "xcsh-docs:data-sources:nginx_service_discovery:examples", "path": "documentation/data-sources/nginx_service_discovery/examples/data-source/index.md", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_service_discovery/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_nginx_service_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nginx_service_discovery/data-source.tf`; digest `sha256:c46b908211689c5888bfaaf70a487544236d409e407922fa907fe86ca0588fa6`.

```terraform
# NginxServiceDiscovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NginxServiceDiscovery by name
data "xcsh_nginx_service_discovery" "example" {
  name      = "example-nginx-service-discovery"
  namespace = "staging"
}

output "nginx_service_discovery_id" {
  value = data.xcsh_nginx_service_discovery.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/examples/)
- [xcsh_nginx_service_discovery](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/nginx_service_discovery/)
