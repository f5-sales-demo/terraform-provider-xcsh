---
page_title: "xcsh_dns_zone"
subcategory: "DNS"
description: "Reads an existing DNS zone in its namespace."
xcsh_docs: {"aliases": ["dns zone"], "body_bytes": 1755, "body_sha256": "sha256:68a701e61249a01ab4d7be0692b919ada80a4db970346fc22154264100ba0ba9", "capabilities": ["dns"], "category": "dns", "child_ids": ["xcsh-docs:data-sources:dns_zone:reference", "xcsh-docs:data-sources:dns_zone:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:dns_zone:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/dns_zone/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "data-sources", "registry_anchor": "canonical-0002233020101332-0312300000031301-3011102330031321-2230013123231120-3001300022000230-1330132013110022-2030110121003300-0222000222331131", "registry_path": "docs/data-sources/dns_zone.md", "relationships": [{"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:optional", "target_id": "xcsh-docs:resources:dns_load_balancer:fundamentals", "type": "advisory"}], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Reads an existing DNS zone in its namespace.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_dns_zone

Breadcrumbs:

- xcsh_dns_zone

Reads an existing DNS zone in its namespace.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `dns_load_balancer`.

- dns_load_balancer: Geographic or weighted DNS routing

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSZone Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSZone by name
data "xcsh_dns_zone" "example" {
  name      = "example-dns-zone"
  namespace = "system"
}

# Fail closed when this stack depends on an externally owned zone.
resource "terraform_data" "require_managed_records" {
  lifecycle {
    precondition {
      condition = try(
        data.xcsh_dns_zone.example.primary.allow_http_lb_managed_records,
        false
      )
      error_message = "The selected DNS zone must enable HTTP LB managed records."
    }
  }
}

output "dns_zone_id" {
  value = data.xcsh_dns_zone.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/examples/)
