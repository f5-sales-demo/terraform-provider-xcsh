---
page_title: "xcsh_dns_load_balancer"
subcategory: "DNS"
description: "xcsh_dns_load_balancer for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 1403, "body_sha256": "sha256:da2f461a4ca3984d82f56319b84f93f5135bc009045dffb9270fb546527c27f1", "canonical_id": "xcsh-docs:resources:dns_load_balancer:fundamentals", "child_ids": ["xcsh-docs:resources:dns_load_balancer:reference", "xcsh-docs:resources:dns_load_balancer:examples", "xcsh-docs:resources:dns_load_balancer:import", "xcsh-docs:resources:dns_load_balancer:timeouts"], "collection_id": "xcsh-docs:resources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_load_balancer:fundamentals", "parent_id": null, "path": "docs/resources/dns_load_balancer.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_load_balancer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_dns_load_balancer for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_dns_load_balancer

Breadcrumbs:

- xcsh_dns_load_balancer

Manages DNS Load Balancer in a given namespace. If one already exist it will give a error in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Required dependencies: `dns_zone`.

- dns_zone: Parent zone for DNS records

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--dns_load_balancer--reference.md)
- [Examples](../guides/resources--dns_load_balancer--examples.md)
- [Import](../guides/resources--dns_load_balancer--import.md)
- [Timeouts](../guides/resources--dns_load_balancer--timeouts.md)
