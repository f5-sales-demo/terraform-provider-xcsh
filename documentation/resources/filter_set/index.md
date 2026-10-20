---
page_title: "xcsh_filter_set"
subcategory: ""
description: "Manages specification in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["filter set"], "body_bytes": 1518, "body_sha256": "sha256:92e0e06f705409ddd48a1d0e5dac8eec6daeebaf747e9d98301961b917b0319e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:filter_set:reference", "xcsh-docs:resources:filter_set:examples", "xcsh-docs:resources:filter_set:import", "xcsh-docs:resources:filter_set:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:filter_set:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/filter_set/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-3033322202313333-3021122111030012-3023212212330123-0300102111110132-3312003112301221-1222013003030021-1210202002210210-0310110313101001", "registry_path": "docs/resources/filter_set.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/index.txt", "spec_pin_digest": "sha256:2a27992e2c450f8ea5ba537ebc7f5c26b0dfbc946db7ef0e7e86625dba9fb492", "summary": "Manages specification in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.1.4", "schema_components": ["filter_setCreateRequest"], "target_commit": "a9149265e7737f6841c289490cb1a3ca946de83e"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_filter_set

Breadcrumbs:

- xcsh_filter_set

Manages specification in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FilterSet Resource Example
# Manages specification in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FilterSet configuration
resource "xcsh_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "staging"

  context_key = "example-value"
}
```

## Root configuration

Required root properties: `context_key`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/lifecycle/timeouts/)
