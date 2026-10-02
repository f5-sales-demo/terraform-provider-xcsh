---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_registrations_by_site."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1359, "body_sha256": "sha256:e4b8aede543c8177bcfc0fac7fd5dade4d26c8691a66cccd2830ab5f094683f1", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_registrations_by_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:c5f91ec913e86f87149f588895f436b17703e7ae2ae9655f9226a6f3b281b896", "source_path": "examples/data-sources/xcsh_site_registrations_by_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_registrations_by_site:example:data-source", "parent_id": "xcsh-docs:data-sources:site_registrations_by_site:examples", "path": "documentation/data-sources/site_registrations_by_site/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "site_registrations_by_site", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3331332010233120-2013201132100112-3203020223211211-1111333330212103-0032231210120130-0033023333220123-3213303212230122-0133100331112311", "registry_path": "docs/guides/data-sources--site_registrations_by_site--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_registrations_by_site/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_site_registrations_by_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_registrations_by_site/data-source.tf`; digest `sha256:c5f91ec913e86f87149f588895f436b17703e7ae2ae9655f9226a6f3b281b896`.

```terraform
# SiteRegistrationsBySite DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_registrations_by_site" "example" {
  site_name = "example-value"
}

output "site_registrations_by_site_result" {
  value = data.xcsh_site_registrations_by_site.example
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/examples/)
- [xcsh_site_registrations_by_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_registrations_by_site/)
