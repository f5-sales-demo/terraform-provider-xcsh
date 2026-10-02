---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_managed_client_customer_support_comments."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1508, "body_sha256": "sha256:e3e7d1a9804516a347b28b17ca58ee52cdae183ea001fb9fec10cf223addcac4", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:12767ce24f59f5ccb7ea0d90d9e1667a91d868bfd5e02b49dd1c582b7f041c1d", "source_path": "examples/data-sources/xcsh_managed_client_customer_support_comments/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:managed_client_customer_support_comments:example:data-source", "parent_id": "xcsh-docs:data-sources:managed_client_customer_support_comments:examples", "path": "documentation/data-sources/managed_client_customer_support_comments/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "managed_client_customer_support_comments", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-3100013300203120-1302101310131031-0112230210132103-3013322000120302-0231112013020300-2012331310222323-1333201110213001-1103010231112222", "registry_path": "docs/guides/data-sources--managed_client_customer_support_comments--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/managed_client_customer_support_comments/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_managed_client_customer_support_comments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
