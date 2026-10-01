---
page_title: "xcsh_dns_load_balancer"
subcategory: "DNS"
description: "xcsh_dns_load_balancer for xcsh_dns_load_balancer."
xcsh_docs: {"aliases": [], "body_bytes": 1506, "body_sha256": "sha256:6346bb69c476f31015cf34ebb0c346086b6e4305d39aac5ba8c3cb8dc99eace4", "child_ids": ["xcsh-docs:data-sources:dns_load_balancer:reference", "xcsh-docs:data-sources:dns_load_balancer:examples"], "collection_id": "xcsh-docs:data-sources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_load_balancer:fundamentals", "parent_id": null, "path": "documentation/data-sources/dns_load_balancer/index.md", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_load_balancer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_dns_load_balancer for xcsh_dns_load_balancer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
# DNSLoadBalancer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSLoadBalancer by name
data "xcsh_dns_load_balancer" "example" {
  name      = "example-dns-load-balancer"
  namespace = "system"
}

output "dns_load_balancer_id" {
  value = data.xcsh_dns_load_balancer.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/examples/)
