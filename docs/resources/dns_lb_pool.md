---
page_title: "xcsh_dns_lb_pool"
subcategory: ""
description: "xcsh_dns_lb_pool for xcsh_dns_lb_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1341, "body_sha256": "sha256:2fff8c304b8d1ffd88cabc1ecf82cc14e064464bb598f626a3827ebbf0ae9c12", "canonical_id": "xcsh-docs:resources:dns_lb_pool:fundamentals", "child_ids": ["xcsh-docs:resources:dns_lb_pool:reference", "xcsh-docs:resources:dns_lb_pool:examples", "xcsh-docs:resources:dns_lb_pool:import", "xcsh-docs:resources:dns_lb_pool:timeouts"], "collection_id": "xcsh-docs:resources:dns_lb_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_lb_pool:fundamentals", "parent_id": null, "path": "docs/resources/dns_lb_pool.md", "provider_name": "dns_lb_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_lb_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_dns_lb_pool for xcsh_dns_lb_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_dns_lb_pool

Breadcrumbs:

- xcsh_dns_lb_pool

Manages DNS Load Balancer Pool in a given namespace. If one already exist it will give a error in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--dns_lb_pool--reference.md)
- [Examples](../guides/resources--dns_lb_pool--examples.md)
- [Import](../guides/resources--dns_lb_pool--import.md)
- [Timeouts](../guides/resources--dns_lb_pool--timeouts.md)
