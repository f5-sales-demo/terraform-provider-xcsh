---
page_title: "Resource"
subcategory: "DNS"
description: "Resource for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1039, "body_sha256": "sha256:25edcb9746192062014591edd0dc425e7acc47790509cb1a6f17a9a8bf51d5fe", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:103fabb6495b3c7befdb807ff9ebc6d11ed6c79f3f4a4083e4f2815709c3a01c", "source_path": "examples/resources/xcsh_dns_load_balancer/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:dns_load_balancer:example:resource", "parent_id": "xcsh-docs:resources:dns_load_balancer:examples", "path": "documentation/resources/dns_load_balancer/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3122312210120211-2301233313110112-1223001302033003-3321333322000211-0233313321130303-3302100022220202-2101330001313100-0300131302133001", "registry_path": "docs/guides/resources--dns_load_balancer--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/examples/resource/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Resource for xcsh_dns_load_balancer.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_dns_load_balancer](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_load_balancer/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_load_balancer/resource.tf`; digest `sha256:103fabb6495b3c7befdb807ff9ebc6d11ed6c79f3f4a4083e4f2815709c3a01c`.

```terraform
# DNSLoadBalancer Resource Example
# Manages DNS Load Balancer in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSLoadBalancer configuration
resource "xcsh_dns_load_balancer" "example" {
  name      = "example-dns-load-balancer"
  namespace = "system"
}
```
