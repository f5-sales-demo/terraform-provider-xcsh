---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_trusted_ca_list."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1316, "body_sha256": "sha256:0059069f243ea5e00b099045f76142dd9d66a26d390aef34902eba29be9f8ac1", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:trusted_ca_list:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:444d0b10e4a27614d6854b1c286e2b90a9468950261cdb881f0b8b693ceeb1e7", "source_path": "examples/data-sources/xcsh_trusted_ca_list/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:trusted_ca_list:example:data-source", "parent_id": "xcsh-docs:data-sources:trusted_ca_list:examples", "path": "documentation/data-sources/trusted_ca_list/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "trusted_ca_list", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3212322202330102-2120312121312322-1122023321111303-2030002312232112-1123333303312102-3010220332113302-1000302221200323-1130000003121232", "registry_path": "docs/guides/data-sources--trusted_ca_list--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/trusted_ca_list/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_trusted_ca_list.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["trusted_ca_listCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/trusted_ca_list/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/trusted_ca_list/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_trusted_ca_list/data-source.tf`; digest `sha256:444d0b10e4a27614d6854b1c286e2b90a9468950261cdb881f0b8b693ceeb1e7`.

```terraform
# TrustedCAList Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TrustedCAList by name
data "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}

output "trusted_ca_list_id" {
  value = data.xcsh_trusted_ca_list.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/trusted_ca_list/examples/)
- [xcsh_trusted_ca_list](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/trusted_ca_list/)
