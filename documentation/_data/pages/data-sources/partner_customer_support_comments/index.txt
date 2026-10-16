---
page_title: "xcsh_partner_customer_support_comments"
subcategory: ""
description: "Reads Partner Customer Support Comments information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["partner customer support comments"], "body_bytes": 1416, "body_sha256": "sha256:1b1513e86d98a504681b1e3bbea90cb19bf142c7fa7fa43bbfde194b7ba7cb79", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:partner_customer_support_comments:reference", "xcsh-docs:data-sources:partner_customer_support_comments:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:partner_customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:partner_customer_support_comments:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/partner_customer_support_comments/index.md", "product": "distributed-cloud", "provider_name": "partner_customer_support_comments", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-1102231002301002-3330320101030031-2203303313002132-2033021213111102-1030111323113201-3321222303201001-3333333120001110-2233122133002213", "registry_path": "docs/data-sources/partner_customer_support_comments.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/partner_customer_support_comments/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Reads Partner Customer Support Comments information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": [], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_partner_customer_support_comments

Breadcrumbs:

- xcsh_partner_customer_support_comments

Reads Partner Customer Support Comments information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# PartnerCustomerSupportComments DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_partner_customer_support_comments" "example" {
  tp_id = "example-value"
}

output "partner_customer_support_comments_result" {
  value = data.xcsh_partner_customer_support_comments.example
}
```

## Root configuration

Required root properties: `tp_id`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/examples/)
