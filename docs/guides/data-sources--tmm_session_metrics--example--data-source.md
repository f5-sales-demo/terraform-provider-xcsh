---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_tmm_session_metrics."
xcsh_docs: {"aliases": [], "body_bytes": 1077, "body_sha256": "sha256:c105a76ded86786dcc7730dd6007f5f2e167e06f930d415d7568a9306e8eab6a", "canonical_id": "xcsh-docs:data-sources:tmm_session_metrics:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:68c10b51b9e48b65d8a384096eef45f5c9fac00e135d661c65b0101b25c350c0", "source_path": "examples/data-sources/xcsh_tmm_session_metrics/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:tmm_session_metrics:example:data-source", "parent_id": "xcsh-docs:data-sources:tmm_session_metrics:examples", "path": "docs/guides/data-sources--tmm_session_metrics--example--data-source.md", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_tmm_session_metrics.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md)
- [Examples](data-sources--tmm_session_metrics--examples.md)
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

## Next pages

- [Examples](data-sources--tmm_session_metrics--examples.md)
- [xcsh_tmm_session_metrics](../data-sources/tmm_session_metrics.md)
