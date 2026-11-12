---
page_title: "xcsh_customer_support_comments"
subcategory: ""
description: "Reads Customer Support Comments information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["customer support comments"], "body_bytes": 1343, "body_sha256": "sha256:d53ac3ec7a8cea4db60ef7514f5c5e4d075b519cb027f65ba2a818cb62cba41e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:customer_support_comments:reference", "xcsh-docs:data-sources:customer_support_comments:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:customer_support_comments:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/customer_support_comments/index.md", "product": "distributed-cloud", "provider_name": "customer_support_comments", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "data-sources", "registry_anchor": "canonical-3233331010300203-0002001013222330-0201130211303202-2232020123310031-2211320023331130-0033233023122313-1203101322331330-3321120232032311", "registry_path": "docs/data-sources/customer_support_comments.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/customer_support_comments/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Reads Customer Support Comments information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": [], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_customer_support_comments

Breadcrumbs:

- xcsh_customer_support_comments

Reads Customer Support Comments information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CustomerSupportComments DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_customer_support_comments" "example" {
  name = "example-value"
}

output "customer_support_comments_result" {
  value = data.xcsh_customer_support_comments.example
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/examples/)
