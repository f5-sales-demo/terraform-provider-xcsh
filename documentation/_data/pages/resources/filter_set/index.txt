---
page_title: "xcsh_filter_set"
subcategory: ""
description: "Manages specification in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["filter set"], "body_bytes": 1505, "body_sha256": "sha256:4d978cd60326483ac219df43891a5334d98b8661d3aa0a4c939c846cae8644fb", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:resources:filter_set:reference", "xcsh-docs:resources:filter_set:examples", "xcsh-docs:resources:filter_set:import", "xcsh-docs:resources:filter_set:timeouts"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:resources:filter_set:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/filter_set/index.md", "product": "distributed-cloud", "provider_name": "filter_set", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-3033322202313333-3021122111030012-3023212212330123-0300102111110132-3312003112301221-1222013003030021-1210202002210210-0310110313101001", "registry_path": "docs/resources/filter_set.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/filter_set/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages specification in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["filter_setCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
