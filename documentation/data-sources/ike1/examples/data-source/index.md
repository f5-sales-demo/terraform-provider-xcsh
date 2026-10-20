---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_ike1."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 973, "body_sha256": "sha256:b72c3b35fd129f07a8255ed3750117d2fb013eb8a6dfed6266da3c3164f0fc03", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:ike1:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1f5aa2a482784f4287f2518ad47fd88a6918d0ecdd01b6d56a7fd34d407e402f", "source_path": "examples/data-sources/xcsh_ike1/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:ike1:example:data-source", "parent_id": "xcsh-docs:data-sources:ike1:examples", "path": "documentation/data-sources/ike1/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "ike1", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "data-sources", "registry_anchor": "canonical-3022001323333201-2001223231333011-3210311031311210-3033232202212002-0021323023202230-2002021111130230-3113133000002103-0213310311101300", "registry_path": "docs/guides/data-sources--ike1--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike1/examples/data-source/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Data source for xcsh_ike1.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["ike1CreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
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
