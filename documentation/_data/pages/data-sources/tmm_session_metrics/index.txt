---
page_title: "xcsh_tmm_session_metrics"
subcategory: ""
description: "Resource creation operation."
xcsh_docs: {"aliases": ["tmm session metrics"], "body_bytes": 1250, "body_sha256": "sha256:bffcb5ca24df6cd26464ddacf61e22846fd63e33ef0fc176693abdb128ee2dd1", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:tmm_session_metrics:reference", "xcsh-docs:data-sources:tmm_session_metrics:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:tmm_session_metrics:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:tmm_session_metrics:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/tmm_session_metrics/index.md", "product": "distributed-cloud", "provider_name": "tmm_session_metrics", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-0001230110001322-3111312333111112-1132331032100110-2331112112213012-0233311131222231-0011030021132123-2021311322012101-1333301310111303", "registry_path": "docs/data-sources/tmm_session_metrics.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/tmm_session_metrics/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Resource creation operation.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/tmm_session_metrics/examples/)
