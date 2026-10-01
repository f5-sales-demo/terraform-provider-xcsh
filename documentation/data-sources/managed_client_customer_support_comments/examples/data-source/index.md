---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_managed_client_customer_support_comments."
xcsh_docs: {"aliases": [], "body_bytes": 1508, "body_sha256": "sha256:e3e7d1a9804516a347b28b17ca58ee52cdae183ea001fb9fec10cf223addcac4", "child_ids": [], "collection_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:12767ce24f59f5ccb7ea0d90d9e1667a91d868bfd5e02b49dd1c582b7f041c1d", "source_path": "examples/data-sources/xcsh_managed_client_customer_support_comments/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:managed_client_customer_support_comments:example:data-source", "parent_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:examples", "path": "documentation/data-sources/managed_client_customer_support_comments/examples/data-source/index.md", "provider_name": "managed_client_customer_support_comments", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/managed_client_customer_support_comments/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_managed_client_customer_support_comments.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_managed_client_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_managed_client_customer_support_comments/data-source.tf`; digest `sha256:12767ce24f59f5ccb7ea0d90d9e1667a91d868bfd5e02b49dd1c582b7f041c1d`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/examples/)
- [xcsh_managed_client_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/managed_client_customer_support_comments/)
