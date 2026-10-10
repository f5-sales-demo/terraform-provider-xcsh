---
page_title: "xcsh_dns_load_balancer"
subcategory: "DNS"
description: "Reads an existing DNS Load Balancer in its namespace."
xcsh_docs: {"aliases": ["dns load balancer"], "body_bytes": 1458, "body_sha256": "sha256:d4e8cc5ff74c407f9989d7cfe4138532a2aa2ed2f81b6a1cd2c9bf2ba88eacba", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_load_balancer:reference", "xcsh-docs:data-sources:dns_load_balancer:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_load_balancer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_load_balancer:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/dns_load_balancer/index.md", "product": "distributed-cloud", "provider_name": "dns_load_balancer", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1203213300333030-1201131231332213-0311011110313002-2212212103302030-0333010310130302-0012333213231220-3200033212123102-0230103000222103", "registry_path": "docs/data-sources/dns_load_balancer.md", "relationships": [{"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:required", "target_id": "xcsh-docs:resources:dns_zone:fundamentals", "type": "advisory"}], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_load_balancer/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Reads an existing DNS Load Balancer in its namespace.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["dns_load_balancerCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_dns_load_balancer

Breadcrumbs:

- xcsh_dns_load_balancer

Reads an existing DNS Load Balancer in its namespace.

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_load_balancer/examples/)
