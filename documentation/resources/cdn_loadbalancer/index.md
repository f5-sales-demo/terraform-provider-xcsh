---
page_title: "xcsh_cdn_loadbalancer"
subcategory: "Load Balancing"
description: "Manages a CDN Load Balancer resource in F5 Distributed Cloud for content delivery and edge caching with load balancing."
xcsh_docs: {"aliases": ["cdn loadbalancer"], "body_bytes": 1846, "body_sha256": "sha256:82436e125b1494d409d975c2d3b8077e87fcb26afbe040568ff5ac809a12438f", "capabilities": ["cdn"], "category": "cdn", "child_ids": ["xcsh-docs:resources:cdn_loadbalancer:reference", "xcsh-docs:resources:cdn_loadbalancer:examples", "xcsh-docs:resources:cdn_loadbalancer:import", "xcsh-docs:resources:cdn_loadbalancer:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:cdn_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_loadbalancer:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/cdn_loadbalancer/index.md", "product": "distributed-cloud", "provider_name": "cdn_loadbalancer", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0313012231313202-2332120333323111-0303222302313032-0220332322023120-0000220020202100-1331310201321110-1020312313103022-2113220113110323", "registry_path": "docs/resources/cdn_loadbalancer.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_loadbalancer/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Manages a CDN Load Balancer resource in F5 Distributed Cloud for content delivery and edge caching with load balancing.", "tasks": ["configuration"], "unresolved_relationships": [{"name": "cdn_origin_pool", "source": "receipt-pinned-dependency:required"}], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["cdn_loadbalancerCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/cdn_loadbalancer/lifecycle/timeouts/)
