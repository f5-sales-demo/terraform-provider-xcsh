---
page_title: "Data source"
subcategory: "DNS"
description: "Data source for xcsh_dns_zone."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1379, "body_sha256": "sha256:8a86097308f6d55cc05c4e07183c10cd531e82960b48bd593da6e4b18c6f5091", "capabilities": ["dns"], "category": "dns", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:45c5f9a8c56f2ed91ab404e89f23a1d58d7ad35539ecaaccb097a0e0066d41b9", "source_path": "examples/data-sources/xcsh_dns_zone/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:dns_zone:example:data-source", "parent_id": "xcsh-docs:data-sources:dns_zone:examples", "path": "documentation/data-sources/dns_zone/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "dns_zone", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1022231110303301-2111303011121232-1120312000231330-3303002003013002-2303310130300221-1313130113032001-2020202123303123-0122221332001130", "registry_path": "docs/guides/data-sources--dns_zone--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_dns_zone.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_dns_zone](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/dns_zone/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_zone/data-source.tf`; digest `sha256:45c5f9a8c56f2ed91ab404e89f23a1d58d7ad35539ecaaccb097a0e0066d41b9`.

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
