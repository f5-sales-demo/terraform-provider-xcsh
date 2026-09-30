---
page_title: "Data source"
subcategory: "DNS"
description: "Data source for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 1290, "body_sha256": "sha256:9e8c1f7ffaf1ab39d690cf4a743ca0077d5e7f44512450a1770c59684f3654d9", "canonical_id": "xcsh-docs:data-sources:dns_zone:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:dns_zone:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:45c5f9a8c56f2ed91ab404e89f23a1d58d7ad35539ecaaccb097a0e0066d41b9", "source_path": "examples/data-sources/xcsh_dns_zone/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:dns_zone:example:data-source", "parent_id": "xcsh-docs:data-sources:dns_zone:examples", "path": "docs/guides/data-sources--dns_zone--example--data-source.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_zone/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_dns_zone](../data-sources/dns_zone.md)
- [Examples](data-sources--dns_zone--examples.md)
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

## Next pages

- [Examples](data-sources--dns_zone--examples.md)
- [xcsh_dns_zone](../data-sources/dns_zone.md)
