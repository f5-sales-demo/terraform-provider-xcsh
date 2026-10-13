---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_tmm_session_metrics."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1034, "body_sha256": "sha256:b475279546661ae8809c1fe74bcd1f457279692e19aa47c779e5c92625019df6", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:68c10b51b9e48b65d8a384096eef45f5c9fac00e135d661c65b0101b25c350c0", "source_path": "examples/data-sources/xcsh_tmm_session_metrics/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:tmm_session_metrics:example:data-source", "parent_id": "xcsh-docs:data-sources:tmm_session_metrics:examples", "path": "documentation/data-sources/tmm_session_metrics/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:1b9ea59a790d92083efb68ff1abebe57ebbb3f7b716b91cdb7b0416134676174", "provider_type": "data-sources", "registry_anchor": "canonical-2021203033233330-1201012121303000-3220311211300332-0310213011212013-3320303202032222-3320010012031321-1300330000321020-2332013121101232", "registry_path": "docs/guides/data-sources--tmm_session_metrics--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/examples/data-source/index.txt", "spec_pin_digest": "sha256:cf249826c2873b41e4c4f1a2341c0bebf16f66cf6d95d5c5b67ae7259e61767b", "summary": "Data source for xcsh_tmm_session_metrics.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.6", "schema_components": [], "target_commit": "06989705b3a435748dea253b9277f572cadb19c4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_tmm_session_metrics](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_tmm_session_metrics/data-source.tf`; digest `sha256:68c10b51b9e48b65d8a384096eef45f5c9fac00e135d661c65b0101b25c350c0`.

```terraform
# TmmSessionMetrics DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_tmm_session_metrics" "example" {
  namespace = "example-value"
}

output "tmm_session_metrics_result" {
  value = data.xcsh_tmm_session_metrics.example
}
```
