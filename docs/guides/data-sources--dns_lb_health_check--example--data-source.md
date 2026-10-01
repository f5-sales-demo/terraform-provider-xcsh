---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_dns_lb_health_check."
xcsh_docs: {"aliases": [], "body_bytes": 1159, "body_sha256": "sha256:c84717f9ebd6ad8cc309c3a05c701b751d53befe5298ce580b1066aa5a1239a0", "canonical_id": "xcsh-docs:data-sources:dns_lb_health_check:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:dns_lb_health_check:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:687cdf9541bb7302cd31e410ec561a4c9e1516b1f24c1acc1a51bcdfad32aecc", "source_path": "examples/data-sources/xcsh_dns_lb_health_check/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:dns_lb_health_check:example:data-source", "parent_id": "xcsh-docs:data-sources:dns_lb_health_check:examples", "path": "docs/guides/data-sources--dns_lb_health_check--example--data-source.md", "provider_name": "dns_lb_health_check", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/dns_lb_health_check/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_dns_lb_health_check.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["dns_lb_health_checkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md)
- [Examples](data-sources--dns_lb_health_check--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_lb_health_check/data-source.tf`; digest `sha256:687cdf9541bb7302cd31e410ec561a4c9e1516b1f24c1acc1a51bcdfad32aecc`.

```terraform
# DNSLBHealthCheck Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DNSLBHealthCheck by name
data "xcsh_dns_lb_health_check" "example" {
  name      = "example-dns-lb-health-check"
  namespace = "system"
}

output "dns_lb_health_check_id" {
  value = data.xcsh_dns_lb_health_check.example.id
}
```

## Next pages

- [Examples](data-sources--dns_lb_health_check--examples.md)
- [xcsh_dns_lb_health_check](../data-sources/dns_lb_health_check.md)
