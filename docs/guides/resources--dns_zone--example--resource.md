---
page_title: "Resource"
subcategory: "DNS"
description: "Resource for xcsh_dns_zone."
xcsh_docs: {"aliases": [], "body_bytes": 1022, "body_sha256": "sha256:effd6e3c281aa3f95e9111d5c98a40800c9652769682a01006c5a9c3ea03e8aa", "canonical_id": "xcsh-docs:resources:dns_zone:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:dns_zone:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a65c60c4a2a466885fcf90c9e747d8f8ba0052be1693d774ae212f433fd43631", "source_path": "examples/resources/xcsh_dns_zone/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:dns_zone:example:resource", "parent_id": "xcsh-docs:resources:dns_zone:examples", "path": "docs/guides/resources--dns_zone--example--resource.md", "provider_name": "dns_zone", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/dns_zone/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_dns_zone.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_zoneCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_dns_zone](../resources/dns_zone.md)
- [Examples](resources--dns_zone--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_dns_zone/resource.tf`; digest `sha256:a65c60c4a2a466885fcf90c9e747d8f8ba0052be1693d774ae212f433fd43631`.

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

## Next pages

- [Examples](resources--dns_zone--examples.md)
- [xcsh_dns_zone](../resources/dns_zone.md)
