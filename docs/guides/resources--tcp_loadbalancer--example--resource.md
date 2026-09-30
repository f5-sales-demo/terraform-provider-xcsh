---
page_title: "Resource"
subcategory: "Load Balancing"
description: "Resource for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": [], "body_bytes": 1028, "body_sha256": "sha256:da8625fd1e9a69e4ef58f6c85968ed802d45d70039e98972f930e6dba6929b2a", "canonical_id": "xcsh-docs:resources:tcp_loadbalancer:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ff5551ee3b1058272666fb12f9288f29debebceead8ddd95379ba00b7d0384b2", "source_path": "examples/resources/xcsh_tcp_loadbalancer/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:tcp_loadbalancer:example:resource", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:examples", "path": "docs/guides/resources--tcp_loadbalancer--example--resource.md", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_tcp_loadbalancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
- [Examples](resources--tcp_loadbalancer--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_tcp_loadbalancer/resource.tf`; digest `sha256:ff5551ee3b1058272666fb12f9288f29debebceead8ddd95379ba00b7d0384b2`.

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

## Next pages

- [Examples](resources--tcp_loadbalancer--examples.md)
- [xcsh_tcp_loadbalancer](../resources/tcp_loadbalancer.md)
