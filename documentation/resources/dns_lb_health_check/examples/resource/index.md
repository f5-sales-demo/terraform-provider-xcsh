---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_dns_lb_health_check."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1309, "body_sha256": "sha256:7f5a048428f22deb786c13afc90e9029c053df2b2f1c86213b6249ea3c2aa9a2", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_health_check:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:219de26c8a2c69aaf219ac20d2e05053708abed48480ac00f40c1cb16da5893e", "source_path": "examples/resources/xcsh_dns_lb_health_check/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:dns_lb_health_check:example:resource", "parent_id": "xcsh-docs:resources:dns_lb_health_check:examples", "path": "documentation/resources/dns_lb_health_check/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-1212210331222111-1213312020222310-0202113023013312-0002130200321233-1212003200221131-3131131133211100-3013031102332123-2121000232231201", "registry_path": "docs/guides/resources--dns_lb_health_check--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_health_check/examples/resource/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Resource for xcsh_dns_lb_health_check.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/examples/)
- [xcsh_dns_lb_health_check](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/)
