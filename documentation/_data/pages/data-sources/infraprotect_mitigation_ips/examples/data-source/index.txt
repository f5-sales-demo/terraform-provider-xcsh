---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_infraprotect_mitigation_ips."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1136, "body_sha256": "sha256:cfffda448d825692eb4013a4e1a591d2dabdf4f0d9a4cde74bb3bf61b596d9f3", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0666bca16af2fe6a216b181ad7ffa5f07262937f90b186cb1b67917f3286333d", "source_path": "examples/data-sources/xcsh_infraprotect_mitigation_ips/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:example:data-source", "parent_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:examples", "path": "documentation/data-sources/infraprotect_mitigation_ips/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_mitigation_ips", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2323222212111121-1020330310021221-2100022220100331-2103120000311113-3220321301201120-0113320021300110-2202332100312021-0113102330213302", "registry_path": "docs/guides/data-sources--infraprotect_mitigation_ips--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/infraprotect_mitigation_ips/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_infraprotect_mitigation_ips.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_infraprotect_mitigation_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_infraprotect_mitigation_ips/data-source.tf`; digest `sha256:0666bca16af2fe6a216b181ad7ffa5f07262937f90b186cb1b67917f3286333d`.

```terraform
# InfraprotectMitigationIps DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_infraprotect_mitigation_ips" "example" {
  mitigation_id = "example-value"
  namespace     = "example-value"
}

output "infraprotect_mitigation_ips_result" {
  value = data.xcsh_infraprotect_mitigation_ips.example
}
```
