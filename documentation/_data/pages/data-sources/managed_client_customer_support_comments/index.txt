---
page_title: "xcsh_managed_client_customer_support_comments"
subcategory: ""
description: "xcsh_managed_client_customer_support_comments for xcsh_managed_client_customer_support_comments."
xcsh_docs: {"aliases": [], "body_bytes": 1409, "body_sha256": "sha256:f438962c37977ff708c3c3da99e5429659a5569cd52d2ba74bc011a7b88ed5f2", "child_ids": ["xcsh-docs:data-sources:managed_client_customer_support_comments:reference", "xcsh-docs:data-sources:managed_client_customer_support_comments:examples"], "collection_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:managed_client_customer_support_comments:fundamentals", "parent_id": null, "path": "documentation/data-sources/managed_client_customer_support_comments/index.md", "provider_name": "managed_client_customer_support_comments", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/managed_client_customer_support_comments/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_managed_client_customer_support_comments for xcsh_managed_client_customer_support_comments.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_managed_client_customer_support_comments

Breadcrumbs:

- xcsh_managed_client_customer_support_comments

Resource retrieval operation.

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/examples/)
