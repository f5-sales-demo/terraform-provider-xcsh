---
page_title: "Data source"
subcategory: "Load Balancing"
description: "Data source for xcsh_tcp_loadbalancer."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1091, "body_sha256": "sha256:40e3f3a64c5dd9e528b291ea13f70346b4e3da79184d6976c696786addf3deec", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:tcp_loadbalancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:8aadc1f29b777e99791395cd3d6a0a6d69101be8986adffcec19d0c93ccfd1f9", "source_path": "examples/data-sources/xcsh_tcp_loadbalancer/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:tcp_loadbalancer:example:data-source", "parent_id": "xcsh-docs:data-sources:tcp_loadbalancer:examples", "path": "documentation/data-sources/tcp_loadbalancer/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "tcp_loadbalancer", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-0233201203120113-2030322203213203-0121222200220001-1100113313030002-1132311101133310-0203002030003103-3221013332201111-2231100001010010", "registry_path": "docs/guides/data-sources--tcp_loadbalancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tcp_loadbalancer/examples/data-source/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Data source for xcsh_tcp_loadbalancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["tcp_loadbalancerCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_tcp_loadbalancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tcp_loadbalancer/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_tcp_loadbalancer/data-source.tf`; digest `sha256:8aadc1f29b777e99791395cd3d6a0a6d69101be8986adffcec19d0c93ccfd1f9`.

```terraform
# TCPLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TCPLoadBalancer by name
data "xcsh_tcp_loadbalancer" "example" {
  name      = "example-tcp-loadbalancer"
  namespace = "staging"
}

output "tcp_loadbalancer_id" {
  value = data.xcsh_tcp_loadbalancer.example.id
}
```
