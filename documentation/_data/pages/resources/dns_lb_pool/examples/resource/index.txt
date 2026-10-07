---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_dns_lb_pool."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 996, "body_sha256": "sha256:7a07686c5703a6aaa13088c00d4bedb4ee2edea42e9e46ce5078ea884d66f3f1", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_pool:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a0f8bbd985a30a9aa4d49bcfaf3107e9c25c24df2c7857f8b4c1d06b0812d763", "source_path": "examples/resources/xcsh_dns_lb_pool/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:dns_lb_pool:example:resource", "parent_id": "xcsh-docs:resources:dns_lb_pool:examples", "path": "documentation/resources/dns_lb_pool/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-0131031232300121-2212232023211202-2122022301200203-0330133212020000-3310112012123003-2232233201322301-1220111133231000-0121033120100332", "registry_path": "docs/guides/resources--dns_lb_pool--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_pool/examples/resource/index.txt", "spec_pin_digest": "sha256:01157ff3cd6b7e1eaa3fb1bc73d0758e089e0e3bcf6e3d9957ada629b777809a", "summary": "Resource for xcsh_dns_lb_pool.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.2", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "4ee07ee75928a56fa7c8eb6603482f21b8472d1d"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_dns_lb_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_pool/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_lb_pool/resource.tf`; digest `sha256:a0f8bbd985a30a9aa4d49bcfaf3107e9c25c24df2c7857f8b4c1d06b0812d763`.

```terraform
# DNSLBPool Resource Example
# Manages DNS Load Balancer Pool in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSLBPool configuration
resource "xcsh_dns_lb_pool" "example" {
  name      = "example-dns-lb-pool"
  namespace = "system"
}
```
