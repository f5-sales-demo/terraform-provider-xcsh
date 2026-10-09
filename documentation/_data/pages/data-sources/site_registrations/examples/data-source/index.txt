---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_registrations."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1027, "body_sha256": "sha256:9b02d4c16220b3ed529de873c0d511944a213d8fca37544990de0ad6d885f8f8", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:382726257bf8ed644e4ba85a525a0636702e8ea87e1b309384adf95a431060c9", "source_path": "examples/data-sources/xcsh_site_registrations/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_registrations:example:data-source", "parent_id": "xcsh-docs:data-sources:site_registrations:examples", "path": "documentation/data-sources/site_registrations/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "site_registrations", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-1032132320201220-3321013310231312-0221133002123132-0200103001330002-1330032302230000-0313003301013012-3312011020312120-1021032233230303", "registry_path": "docs/guides/data-sources--site_registrations--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_site_registrations.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_registrations](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_registrations/data-source.tf`; digest `sha256:382726257bf8ed644e4ba85a525a0636702e8ea87e1b309384adf95a431060c9`.

```terraform
# SiteRegistrations DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations" "example" {
  namespace = "example-value"
}

output "site_registrations_result" {
  value = data.xcsh_site_registrations.example
}
```
