---
page_title: "xcsh_dns_lb_health_check"
subcategory: ""
description: "xcsh_dns_lb_health_check for xcsh_dns_lb_health_check."
xcsh_docs: {"aliases": [], "body_bytes": 1435, "body_sha256": "sha256:56e5c9f5064f4223835ad59601aa9de3cc541152f571b02f1873fe8925c3916d", "canonical_id": "xcsh-docs:resources:dns_lb_health_check:fundamentals", "child_ids": ["xcsh-docs:resources:dns_lb_health_check:reference", "xcsh-docs:resources:dns_lb_health_check:examples", "xcsh-docs:resources:dns_lb_health_check:import", "xcsh-docs:resources:dns_lb_health_check:timeouts"], "collection_id": "xcsh-docs:resources:dns_lb_health_check:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_health_check:fundamentals", "parent_id": null, "path": "docs/resources/dns_lb_health_check.md", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_health_check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_dns_lb_health_check for xcsh_dns_lb_health_check.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

- [Property reference](../guides/resources--dns_lb_health_check--reference.md)
- [Examples](../guides/resources--dns_lb_health_check--examples.md)
- [Import](../guides/resources--dns_lb_health_check--import.md)
- [Timeouts](../guides/resources--dns_lb_health_check--timeouts.md)
