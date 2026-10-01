---
page_title: "xcsh_filter_set"
subcategory: ""
description: "xcsh_filter_set for xcsh_filter_set."
xcsh_docs: {"aliases": [], "body_bytes": 1505, "body_sha256": "sha256:4d978cd60326483ac219df43891a5334d98b8661d3aa0a4c939c846cae8644fb", "child_ids": ["xcsh-docs:resources:filter_set:reference", "xcsh-docs:resources:filter_set:examples", "xcsh-docs:resources:filter_set:import", "xcsh-docs:resources:filter_set:timeouts"], "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:filter_set:fundamentals", "parent_id": null, "path": "documentation/resources/filter_set/index.md", "provider_name": "filter_set", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_filter_set for xcsh_filter_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["filter_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/filter_set/lifecycle/timeouts/)
