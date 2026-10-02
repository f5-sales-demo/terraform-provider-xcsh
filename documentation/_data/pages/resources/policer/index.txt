---
page_title: "xcsh_policer"
subcategory: ""
description: "Manages new policer with traffic rate limits in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["policer"], "body_bytes": 1586, "body_sha256": "sha256:7a9d0967d1859ad81e0ff2b82cac123af14c34d1a1d81e407df691da6794e3d5", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:policer:reference", "xcsh-docs:resources:policer:examples", "xcsh-docs:resources:policer:import", "xcsh-docs:resources:policer:timeouts"], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:policer:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/policer/index.md", "product": "distributed-cloud", "provider_name": "policer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2122311132312033-0121013211002130-0311121323112230-0010131112301322-1312303210110212-3032113221002232-1112102011022312-1111003020123113", "registry_path": "docs/resources/policer.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policer/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages new policer with traffic rate limits in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["policerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_policer

Breadcrumbs:

- xcsh_policer

Manages new policer with traffic rate limits in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Policer Resource Example
# Manages new policer with traffic rate limits in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Policer configuration
resource "xcsh_policer" "example" {
  name      = "example-policer"
  namespace = "staging"

  burst_size                 = 1
  committed_information_rate = 1
}
```

## Root configuration

Required root properties: `burst_size`, `committed_information_rate`, `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policer/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policer/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policer/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/policer/lifecycle/timeouts/)
