---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_udp_loadbalancer."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1333, "body_sha256": "sha256:e31f06a391fdcab63f64a533c48fa063ab8567cda315198f5da94c04f195e7ec", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:udp_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f54b33fab2f01b75ab0f74a179601bfe368063a44f1d292cd4b49a5f0936b681", "source_path": "examples/resources/xcsh_udp_loadbalancer/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:udp_loadbalancer:example:resource", "parent_id": "xcsh-docs:resources:udp_loadbalancer:examples", "path": "documentation/resources/udp_loadbalancer/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "udp_loadbalancer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1301313203231213-3220311132100110-0330103211030102-0333331101201213-3012011110303223-3023132303012223-3010301333232332-0310230313101110", "registry_path": "docs/guides/resources--udp_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/udp_loadbalancer/examples/resource/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource for xcsh_udp_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["udp_loadbalancerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/examples/)
- [xcsh_udp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/udp_loadbalancer/)
