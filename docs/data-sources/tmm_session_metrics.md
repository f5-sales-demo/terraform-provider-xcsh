---
page_title: "xcsh_tmm_session_metrics"
subcategory: ""
description: "xcsh_tmm_session_metrics for xcsh_tmm_session_metrics."
xcsh_docs: {"aliases": [], "body_bytes": 1165, "body_sha256": "sha256:51c1fe005131e8f82fa8f99969f5d5e359c10cbb4a05ed5715f9e21fe68d9259", "canonical_id": "xcsh-docs:data-sources:tmm_session_metrics:fundamentals", "child_ids": ["xcsh-docs:data-sources:tmm_session_metrics:reference", "xcsh-docs:data-sources:tmm_session_metrics:examples"], "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tmm_session_metrics:fundamentals", "parent_id": null, "path": "docs/data-sources/tmm_session_metrics.md", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_tmm_session_metrics for xcsh_tmm_session_metrics.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_tmm_session_metrics

Breadcrumbs:

- xcsh_tmm_session_metrics

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--tmm_session_metrics--reference.md)
- [Examples](../guides/data-sources--tmm_session_metrics--examples.md)
