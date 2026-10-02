---
page_title: "xcsh_udp_loadbalancer"
subcategory: ""
description: "Manages a UDP Load Balancer resource in F5 Distributed Cloud for load balancing UDP traffic across origin pools."
xcsh_docs: {"aliases": ["backend servers", "origin servers", "udp loadbalancer", "upstream servers"], "body_bytes": 1649, "body_sha256": "sha256:7834c1db1fd1857842978bc3e981778c2ae464246b3fbcb2cf144771d0a14dfb", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:resources:udp_loadbalancer:reference", "xcsh-docs:resources:udp_loadbalancer:examples", "xcsh-docs:resources:udp_loadbalancer:import", "xcsh-docs:resources:udp_loadbalancer:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:udp_loadbalancer:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/udp_loadbalancer/index.md", "product": "distributed-cloud", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-2031002110012331-0223121021210130-3122111031231033-0102213100033031-3300120332233213-0233003311303232-0311110113213220-1211333320223212", "registry_path": "docs/resources/udp_loadbalancer.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/udp_loadbalancer/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a UDP Load Balancer resource in F5 Distributed Cloud for load balancing UDP traffic across origin pools.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_udp_loadbalancer

Breadcrumbs:

- xcsh_udp_loadbalancer

Manages a UDP Load Balancer resource in F5 Distributed Cloud for load balancing UDP traffic across
origin pools.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# UDPLoadBalancer Resource Example
# Manages a UDP Load Balancer resource in F5 Distributed Cloud for load balancing UDP traffic across origin pools.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UDPLoadBalancer configuration
resource "xcsh_udp_loadbalancer" "example" {
  name      = "example-udp-loadbalancer"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/lifecycle/timeouts/)
