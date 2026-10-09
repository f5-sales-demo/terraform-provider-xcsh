---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_partner_customer_support_comments."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1141, "body_sha256": "sha256:68ba832e2ce5abd397ae5253a97e65e316ab15c78af3037ebdd0738c89af1498", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:partner_customer_support_comments:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1a19f6d9101375158d6964da30975931b7c863facf609615fe0cc1c3633c733c", "source_path": "examples/data-sources/xcsh_partner_customer_support_comments/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:partner_customer_support_comments:example:data-source", "parent_id": "xcsh-docs:data-sources:partner_customer_support_comments:examples", "path": "documentation/data-sources/partner_customer_support_comments/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "partner_customer_support_comments", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2112232030301011-0321211001202223-0133021320113300-1010110020032213-2000033130302233-2230303122323122-0110031303322312-1333010212210112", "registry_path": "docs/guides/data-sources--partner_customer_support_comments--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/partner_customer_support_comments/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_partner_customer_support_comments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": [], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
