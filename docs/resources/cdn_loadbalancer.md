---
page_title: "xcsh_cdn_loadbalancer"
subcategory: "Load Balancing"
description: "xcsh_cdn_loadbalancer for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1545, "body_sha256": "sha256:bb3c2a28309a24ea9890c57c6f9106ffde0e28dfcabb066f25e3e68195135b05", "canonical_id": "xcsh-docs:resources:cdn_loadbalancer:fundamentals", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:reference", "xcsh-docs:resources:cdn_loadbalancer:examples", "xcsh-docs:resources:cdn_loadbalancer:import", "xcsh-docs:resources:cdn_loadbalancer:timeouts"], "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:fundamentals", "parent_id": null, "path": "docs/resources/cdn_loadbalancer.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_cdn_loadbalancer for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_cdn_loadbalancer

Breadcrumbs:

- xcsh_cdn_loadbalancer

Manages a CDN Load Balancer resource in F5 Distributed Cloud for content delivery and edge caching
with load balancing.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `cdn_origin_pool`.

- cdn_origin_pool: Origin servers for CDN content

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CDNLoadBalancer Resource Example
# Manages a CDN Load Balancer resource in F5 Distributed Cloud for content delivery and edge caching with load balancing.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNLoadBalancer configuration
resource "xcsh_cdn_loadbalancer" "example" {
  name      = "example-cdn-loadbalancer"
  namespace = "staging"

  domains = ["example-value"]
}
```

## Root configuration

Required root properties: `domains`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--cdn_loadbalancer--reference.md)
- [Examples](../guides/resources--cdn_loadbalancer--examples.md)
- [Import](../guides/resources--cdn_loadbalancer--import.md)
- [Timeouts](../guides/resources--cdn_loadbalancer--timeouts.md)
