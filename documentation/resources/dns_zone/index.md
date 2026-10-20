---
page_title: "xcsh_dns_zone"
subcategory: "DNS"
description: "Manages DNS Zone in a given namespace. If one already exist it will give a error in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["dns zone"], "body_bytes": 1681, "body_sha256": "sha256:de1eb1f0fedca0a08b594ca783513a3791e930d166ea46a37497b22eb044e37c", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:resources:dns_zone:reference", "xcsh-docs:resources:dns_zone:examples", "xcsh-docs:resources:dns_zone:import", "xcsh-docs:resources:dns_zone:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:resources:dns_zone:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/dns_zone/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3231313303010030-0100221022033133-1333032030112013-0210022032301013-1120013101322111-1000210022113233-0321110131313221-2020312231302323", "registry_path": "docs/resources/dns_zone.md", "relationships": [{"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:optional", "target_id": "xcsh-docs:resources:dns_load_balancer:fundamentals", "type": "advisory"}], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Manages DNS Zone in a given namespace. If one already exist it will give a error in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/dns_zone/lifecycle/timeouts/)
