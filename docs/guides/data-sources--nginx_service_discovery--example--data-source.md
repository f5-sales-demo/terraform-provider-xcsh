---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_nginx_service_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1214, "body_sha256": "sha256:ca74306cb90a3f42e81037b5697f8f7f61e4e32e3dc62038c9f1af55ff962bc5", "canonical_id": "xcsh-docs:data-sources:nginx_service_discovery:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:nginx_service_discovery:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c46b908211689c5888bfaaf70a487544236d409e407922fa907fe86ca0588fa6", "source_path": "examples/data-sources/xcsh_nginx_service_discovery/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:nginx_service_discovery:example:data-source", "parent_id": "xcsh-docs:data-sources:nginx_service_discovery:examples", "path": "docs/guides/data-sources--nginx_service_discovery--example--data-source.md", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/nginx_service_discovery/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_nginx_service_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md)
- [Examples](data-sources--nginx_service_discovery--examples.md)
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

- [Examples](data-sources--nginx_service_discovery--examples.md)
- [xcsh_nginx_service_discovery](../data-sources/nginx_service_discovery.md)
