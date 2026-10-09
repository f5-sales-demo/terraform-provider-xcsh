---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_ike1."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 973, "body_sha256": "sha256:b72c3b35fd129f07a8255ed3750117d2fb013eb8a6dfed6266da3c3164f0fc03", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike1:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1f5aa2a482784f4287f2518ad47fd88a6918d0ecdd01b6d56a7fd34d407e402f", "source_path": "examples/data-sources/xcsh_ike1/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:ike1:example:data-source", "parent_id": "xcsh-docs:data-sources:ike1:examples", "path": "documentation/data-sources/ike1/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "ike1", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-3022001323333201-2001223231333011-3210311031311210-3033232202212002-0021323023202230-2002021111130230-3113133000002103-0213310311101300", "registry_path": "docs/guides/data-sources--ike1--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike1/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_ike1.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["ike1CreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_ike1](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/ike1/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ike1/data-source.tf`; digest `sha256:1f5aa2a482784f4287f2518ad47fd88a6918d0ecdd01b6d56a7fd34d407e402f`.

```terraform
# Ike1 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Ike1 by name
data "xcsh_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"
}

output "ike1_id" {
  value = data.xcsh_ike1.example.id
}
```
