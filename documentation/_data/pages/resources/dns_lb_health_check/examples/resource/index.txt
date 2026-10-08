---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_dns_lb_health_check."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1066, "body_sha256": "sha256:46b0bb3c134fafae14120733832fc46d4cce4640cf8a1e4a0b6d65301b561e1e", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_health_check:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:219de26c8a2c69aaf219ac20d2e05053708abed48480ac00f40c1cb16da5893e", "source_path": "examples/resources/xcsh_dns_lb_health_check/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:dns_lb_health_check:example:resource", "parent_id": "xcsh-docs:resources:dns_lb_health_check:examples", "path": "documentation/resources/dns_lb_health_check/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1212210331222111-1213312020222310-0202113023013312-0002130200321233-1212003200221131-3131131133211100-3013031102332123-2121000232231201", "registry_path": "docs/guides/resources--dns_lb_health_check--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_health_check/examples/resource/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Resource for xcsh_dns_lb_health_check.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_lb_health_check/resource.tf`; digest `sha256:219de26c8a2c69aaf219ac20d2e05053708abed48480ac00f40c1cb16da5893e`.

```terraform
# DNSLBHealthCheck Resource Example
# Manages DNS Load Balancer Health Check in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSLBHealthCheck configuration
resource "xcsh_dns_lb_health_check" "example" {
  name      = "example-dns-lb-health-check"
  namespace = "system"
}
```
