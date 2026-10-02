---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_partner_customer_support_comments."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1432, "body_sha256": "sha256:49ee7a4ac46ee30045984fb27c8a30522264710050db67187bd391117972e184", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:partner_customer_support_comments:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1a19f6d9101375158d6964da30975931b7c863facf609615fe0cc1c3633c733c", "source_path": "examples/data-sources/xcsh_partner_customer_support_comments/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:partner_customer_support_comments:example:data-source", "parent_id": "xcsh-docs:data-sources:partner_customer_support_comments:examples", "path": "documentation/data-sources/partner_customer_support_comments/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "partner_customer_support_comments", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2112232030301011-0321211001202223-0133021320113300-1010110020032213-2000033130302233-2230303122323122-0110031303322312-1333010212210112", "registry_path": "docs/guides/data-sources--partner_customer_support_comments--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/partner_customer_support_comments/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_partner_customer_support_comments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_partner_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/examples/)
- [xcsh_partner_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/partner_customer_support_comments/)
