---
page_title: "xcsh_nginx_service_discovery"
subcategory: ""
description: "xcsh_nginx_service_discovery for xcsh_nginx_service_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1645, "body_sha256": "sha256:544481d6e81718f3fecc717ac81d11c0232e97529328e2862a1afc4c2b76e2a3", "canonical_id": "xcsh-docs:resources:nginx_service_discovery:fundamentals", "child_ids": ["xcsh-docs:resources:nginx_service_discovery:reference", "xcsh-docs:resources:nginx_service_discovery:examples", "xcsh-docs:resources:nginx_service_discovery:import", "xcsh-docs:resources:nginx_service_discovery:timeouts"], "collection_id": "xcsh-docs:resources:nginx_service_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:nginx_service_discovery:fundamentals", "parent_id": null, "path": "docs/resources/nginx_service_discovery.md", "provider_name": "nginx_service_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/nginx_service_discovery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_nginx_service_discovery for xcsh_nginx_service_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["nginx_service_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_nginx_service_discovery

Breadcrumbs:

- xcsh_nginx_service_discovery

Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service
discovery object for a site or virtual site in system namespace. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NginxServiceDiscovery Resource Example
# Manages a Nginx Service Discovery resource in F5 Distributed Cloud for api to create nginx service discovery object for a site or virtual site in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NginxServiceDiscovery configuration
resource "xcsh_nginx_service_discovery" "example" {
  name      = "example-nginx-service-discovery"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--nginx_service_discovery--reference.md)
- [Examples](../guides/resources--nginx_service_discovery--examples.md)
- [Import](../guides/resources--nginx_service_discovery--import.md)
- [Timeouts](../guides/resources--nginx_service_discovery--timeouts.md)
