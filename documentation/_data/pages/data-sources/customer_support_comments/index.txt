---
page_title: "xcsh_customer_support_comments"
subcategory: ""
description: "Reads Customer Support Comments information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["customer support comments"], "body_bytes": 1343, "body_sha256": "sha256:d53ac3ec7a8cea4db60ef7514f5c5e4d075b519cb027f65ba2a818cb62cba41e", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:customer_support_comments:reference", "xcsh-docs:data-sources:customer_support_comments:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:customer_support_comments:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/customer_support_comments/index.md", "product": "distributed-cloud", "provider_name": "customer_support_comments", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3233331010300203-0002001013222330-0201130211303202-2232020123310031-2211320023331130-0033233023122313-1203101322331330-3321120232032311", "registry_path": "docs/data-sources/customer_support_comments.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/customer_support_comments/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reads Customer Support Comments information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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
