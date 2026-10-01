---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_partner_customer_support_comments."
xcsh_docs: {"aliases": [], "body_bytes": 1226, "body_sha256": "sha256:64248d14c42e6939f65f109f91162d1bbd4566c315e009bfaf54af6065705c5e", "canonical_id": "xcsh-docs:data-sources:partner_customer_support_comments:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:partner_customer_support_comments:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1a19f6d9101375158d6964da30975931b7c863facf609615fe0cc1c3633c733c", "source_path": "examples/data-sources/xcsh_partner_customer_support_comments/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:partner_customer_support_comments:example:data-source", "parent_id": "xcsh-docs:data-sources:partner_customer_support_comments:examples", "path": "docs/guides/data-sources--partner_customer_support_comments--example--data-source.md", "provider_name": "partner_customer_support_comments", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/partner_customer_support_comments/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_partner_customer_support_comments.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_partner_customer_support_comments](../data-sources/partner_customer_support_comments.md)
- [Examples](data-sources--partner_customer_support_comments--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_partner_customer_support_comments/data-source.tf`; digest `sha256:1a19f6d9101375158d6964da30975931b7c863facf609615fe0cc1c3633c733c`.

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

## Next pages

- [Examples](data-sources--partner_customer_support_comments--examples.md)
- [xcsh_partner_customer_support_comments](../data-sources/partner_customer_support_comments.md)
