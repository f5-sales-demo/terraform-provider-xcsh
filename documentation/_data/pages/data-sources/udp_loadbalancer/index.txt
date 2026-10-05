---
page_title: "xcsh_udp_loadbalancer"
subcategory: ""
description: "Manages a UDP Load Balancer resource in F5 Distributed Cloud for load balancing UDP traffic across origin pools."
xcsh_docs: {"aliases": ["udp loadbalancer"], "body_bytes": 1399, "body_sha256": "sha256:f05b0d0e2dca0a2c50c85894d29d8b4e6911bcd10900f4cdfd8d4bafc63853e3", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": ["xcsh-docs:data-sources:udp_loadbalancer:reference", "xcsh-docs:data-sources:udp_loadbalancer:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:udp_loadbalancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:udp_loadbalancer:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/udp_loadbalancer/index.md", "product": "distributed-cloud", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-0111123322321202-3332321330232132-0032130102310300-0213133321331211-2320232221123003-0303120113132321-1130032130200121-1121203011232203", "registry_path": "docs/data-sources/udp_loadbalancer.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/udp_loadbalancer/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages a UDP Load Balancer resource in F5 Distributed Cloud for load balancing UDP traffic across origin pools.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
