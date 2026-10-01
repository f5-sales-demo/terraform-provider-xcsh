---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1127, "body_sha256": "sha256:b380c0a593d5b73d451f9a60f813e6bf1d8b71c97aa21bef0288dc88ad7345f9", "canonical_id": "xcsh-docs:resources:udp_loadbalancer:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:udp_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f54b33fab2f01b75ab0f74a179601bfe368063a44f1d292cd4b49a5f0936b681", "source_path": "examples/resources/xcsh_udp_loadbalancer/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:udp_loadbalancer:example:resource", "parent_id": "xcsh-docs:resources:udp_loadbalancer:examples", "path": "docs/guides/resources--udp_loadbalancer--example--resource.md", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/udp_loadbalancer/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_udp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
- [Examples](resources--udp_loadbalancer--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_udp_loadbalancer/resource.tf`; digest `sha256:f54b33fab2f01b75ab0f74a179601bfe368063a44f1d292cd4b49a5f0936b681`.

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

## Next pages

- [Examples](resources--udp_loadbalancer--examples.md)
- [xcsh_udp_loadbalancer](../resources/udp_loadbalancer.md)
