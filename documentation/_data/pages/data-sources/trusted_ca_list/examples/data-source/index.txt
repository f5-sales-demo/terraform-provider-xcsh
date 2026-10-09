---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_trusted_ca_list."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1079, "body_sha256": "sha256:47636b3db1602577027e3b979c879696995ae587b0db6f5b2f3eb086675f3929", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:trusted_ca_list:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:444d0b10e4a27614d6854b1c286e2b90a9468950261cdb881f0b8b693ceeb1e7", "source_path": "examples/data-sources/xcsh_trusted_ca_list/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:trusted_ca_list:example:data-source", "parent_id": "xcsh-docs:data-sources:trusted_ca_list:examples", "path": "documentation/data-sources/trusted_ca_list/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "trusted_ca_list", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3212322202330102-2120312121312322-1122023321111303-2030002312232112-1123333303312102-3010220332113302-1000302221200323-1130000003121232", "registry_path": "docs/guides/data-sources--trusted_ca_list--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/trusted_ca_list/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_trusted_ca_list.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["trusted_ca_listCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/trusted_ca_list/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/trusted_ca_list/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_trusted_ca_list/data-source.tf`; digest `sha256:444d0b10e4a27614d6854b1c286e2b90a9468950261cdb881f0b8b693ceeb1e7`.

```terraform
# TrustedCAList Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TrustedCAList by name
data "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}

output "trusted_ca_list_id" {
  value = data.xcsh_trusted_ca_list.example.id
}
```
