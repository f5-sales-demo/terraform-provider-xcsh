---
page_title: "xcsh_policer"
subcategory: ""
description: "Manages new policer with traffic rate limits in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["policer"], "body_bytes": 1586, "body_sha256": "sha256:7a9d0967d1859ad81e0ff2b82cac123af14c34d1a1d81e407df691da6794e3d5", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:policer:reference", "xcsh-docs:resources:policer:examples", "xcsh-docs:resources:policer:import", "xcsh-docs:resources:policer:timeouts"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:policer:collection", "completeness": "complete", "id": "xcsh-docs:resources:policer:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/policer/index.md", "product": "distributed-cloud", "provider_name": "policer", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-2122311132312033-0121013211002130-0311121323112230-0010131112301322-1312303210110212-3032113221002232-1112102011022312-1111003020123113", "registry_path": "docs/resources/policer.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/policer/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages new policer with traffic rate limits in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["policerCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
