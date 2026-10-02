---
page_title: "xcsh_dns_lb_health_check"
subcategory: ""
description: "Manages DNS Load Balancer Health Check in a given namespace. If one already exist it will give a error in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["dns lb health check"], "body_bytes": 1624, "body_sha256": "sha256:e28a5a26ea838875ec10f8bebd8782229e00cf9e5a7cd085e01bb0476ecbb7eb", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_lb_health_check:reference", "xcsh-docs:resources:dns_lb_health_check:examples", "xcsh-docs:resources:dns_lb_health_check:import", "xcsh-docs:resources:dns_lb_health_check:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_health_check:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/dns_lb_health_check/index.md", "product": "distributed-cloud", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0302000112101031-2320313333300101-1000303021303302-2311331100312232-2321211133331232-0230203302202230-0030030332120212-0001210320031333", "registry_path": "docs/resources/dns_lb_health_check.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages DNS Load Balancer Health Check in a given namespace. If one already exist it will give a error in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_dns_lb_health_check

Breadcrumbs:

- xcsh_dns_lb_health_check

Manages DNS Load Balancer Health Check in a given namespace. If one already exist it will give a
error in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_lb_health_check/lifecycle/timeouts/)
