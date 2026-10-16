---
page_title: "Resource"
subcategory: "Load Balancing"
description: "Resource for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1099, "body_sha256": "sha256:1d5dce80a4f3503bb8f58cb17d6cee336861d796ad9aa4529db46e378825a47f", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:tcp_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:ff5551ee3b1058272666fb12f9288f29debebceead8ddd95379ba00b7d0384b2", "source_path": "examples/resources/xcsh_tcp_loadbalancer/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:tcp_loadbalancer:example:resource", "parent_id": "xcsh-docs:resources:tcp_loadbalancer:examples", "path": "documentation/resources/tcp_loadbalancer/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3030132121321003-2013130021300310-3112112030301121-3102020110020233-2021003331113332-3103233212023320-3122111323100032-3110120202132001", "registry_path": "docs/guides/resources--tcp_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tcp_loadbalancer/examples/resource/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Resource for xcsh_tcp_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/tcp_loadbalancer/examples/)
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
