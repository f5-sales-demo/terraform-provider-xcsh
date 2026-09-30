---
page_title: "xcsh_cdn_loadbalancer"
subcategory: "Load Balancing"
description: "xcsh_cdn_loadbalancer for xcsh_cdn_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1435, "body_sha256": "sha256:3ffa90b98abc7ec0ffab013930bb560267df5e668e6eae6b48abaae701cd2c22", "child_ids": ["xcsh-docs:data-sources:cdn_loadbalancer:reference", "xcsh-docs:data-sources:cdn_loadbalancer:examples"], "collection_id": "xcsh-docs:data-sources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_loadbalancer:fundamentals", "parent_id": null, "path": "documentation/data-sources/cdn_loadbalancer/index.md", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_loadbalancer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_cdn_loadbalancer for xcsh_cdn_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
# CDNLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNLoadBalancer by name
data "xcsh_cdn_loadbalancer" "example" {
  name      = "example-cdn-loadbalancer"
  namespace = "staging"
}

output "cdn_loadbalancer_id" {
  value = data.xcsh_cdn_loadbalancer.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_loadbalancer/examples/)
