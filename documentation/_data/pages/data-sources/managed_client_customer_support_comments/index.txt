---
page_title: "xcsh_managed_client_customer_support_comments"
subcategory: ""
description: "Reads Managed Client Customer Support Comments information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["managed client customer support comments"], "body_bytes": 1478, "body_sha256": "sha256:0ce4c82b1ce43550df955564d4d19e9f9cf7fa83b58cc6f42322e5fa108650cb", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:managed_client_customer_support_comments:reference", "xcsh-docs:data-sources:managed_client_customer_support_comments:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:managed_client_customer_support_comments:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/managed_client_customer_support_comments/index.md", "product": "distributed-cloud", "provider_name": "managed_client_customer_support_comments", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "data-sources", "registry_anchor": "canonical-0121323102230033-0320023001323221-1211130012203100-2333331223231001-1203112221121320-3132121331212101-3013023323130212-2212123012212223", "registry_path": "docs/data-sources/managed_client_customer_support_comments.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/managed_client_customer_support_comments/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Reads Managed Client Customer Support Comments information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": [], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_managed_client_customer_support_comments

Breadcrumbs:

- xcsh_managed_client_customer_support_comments

Reads Managed Client Customer Support Comments information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ManagedClientCustomerSupportComments DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_managed_client_customer_support_comments" "example" {
  tp_id = "example-value"
}

output "managed_client_customer_support_comments_result" {
  value = data.xcsh_managed_client_customer_support_comments.example
}
```

## Root configuration

Required root properties: `tp_id`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/examples/)
