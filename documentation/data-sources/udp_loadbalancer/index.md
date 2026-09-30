---
page_title: "xcsh_udp_loadbalancer"
subcategory: ""
description: "xcsh_udp_loadbalancer for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1300, "body_sha256": "sha256:e794042580813154a8209e0b3e5af4f4d7809ad10de0198c1c67495e721da1ba", "child_ids": ["xcsh-docs:data-sources:udp_loadbalancer:reference", "xcsh-docs:data-sources:udp_loadbalancer:examples"], "collection_id": "xcsh-docs:data-sources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:udp_loadbalancer:fundamentals", "parent_id": null, "path": "documentation/data-sources/udp_loadbalancer/index.md", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/udp_loadbalancer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_udp_loadbalancer for xcsh_udp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
# UDPLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing UDPLoadBalancer by name
data "xcsh_udp_loadbalancer" "example" {
  name      = "example-udp-loadbalancer"
  namespace = "staging"
}

output "udp_loadbalancer_id" {
  value = data.xcsh_udp_loadbalancer.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/udp_loadbalancer/examples/)
