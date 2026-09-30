---
page_title: "xcsh_dns_zone"
subcategory: "DNS"
description: "xcsh_dns_zone for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 1380, "body_sha256": "sha256:97136c6d78cc1a04d6345688e276fcdb57aea1555a71814b357d41ab95f27316", "canonical_id": "xcsh-docs:resources:dns_zone:fundamentals", "child_ids": ["xcsh-docs:resources:dns_zone:reference", "xcsh-docs:resources:dns_zone:examples", "xcsh-docs:resources:dns_zone:import", "xcsh-docs:resources:dns_zone:timeouts"], "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:fundamentals", "parent_id": null, "path": "docs/resources/dns_zone.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_dns_zone for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_dns_zone

Breadcrumbs:

- xcsh_dns_zone

Manages DNS Zone in a given namespace. If one already exist it will give a error in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `dns_load_balancer`.

- dns_load_balancer: Geographic or weighted DNS routing

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSZone Resource Example
# Manages DNS Zone in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DNSZone configuration
resource "xcsh_dns_zone" "example" {
  name      = "example-dns-zone"
  namespace = "system"

  primary {
    allow_http_lb_managed_records = true
  }
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--dns_zone--reference.md)
- [Examples](../guides/resources--dns_zone--examples.md)
- [Import](../guides/resources--dns_zone--import.md)
- [Timeouts](../guides/resources--dns_zone--timeouts.md)
