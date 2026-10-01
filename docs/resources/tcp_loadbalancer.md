---
page_title: "xcsh_tcp_loadbalancer"
subcategory: "Load Balancing"
description: "xcsh_tcp_loadbalancer for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1669, "body_sha256": "sha256:856ee3da33c99b25bae5082019e84832cbeeb68495a149a77d3f9fde4463ad47", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:fundamentals", "child_ids": ["xcsh-docs:resources:tcp_loadbalancer:reference", "xcsh-docs:resources:tcp_loadbalancer:examples", "xcsh-docs:resources:tcp_loadbalancer:import", "xcsh-docs:resources:tcp_loadbalancer:timeouts"], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:tcp_loadbalancer:fundamentals", "parent_id": null, "path": "docs/resources/tcp_loadbalancer.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_tcp_loadbalancer for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_tcp_loadbalancer

Breadcrumbs:

- xcsh_tcp_loadbalancer

Manages a TCP Load Balancer resource in F5 Distributed Cloud for load balancing TCP traffic across
origin pools.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `origin_pool`.

Optional integrations: `healthcheck`.

- origin_pool: Backend servers for TCP/UDP traffic

- healthcheck: Monitor origin server health

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# TCPLoadBalancer Resource Example
# Manages a TCP Load Balancer resource in F5 Distributed Cloud for load balancing TCP traffic across origin pools.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic TCPLoadBalancer configuration
resource "xcsh_tcp_loadbalancer" "example" {
  name      = "example-tcp-loadbalancer"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--tcp_loadbalancer--reference.md)
- [Examples](../guides/resources--tcp_loadbalancer--examples.md)
- [Import](../guides/resources--tcp_loadbalancer--import.md)
- [Timeouts](../guides/resources--tcp_loadbalancer--timeouts.md)
