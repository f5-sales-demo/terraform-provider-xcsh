---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_infraprotect_mitigation_ips."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1409, "body_sha256": "sha256:303911c507bf833902df7b37c046b9ec9be160b54522191073d9bd0bc6a3bf73", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0666bca16af2fe6a216b181ad7ffa5f07262937f90b186cb1b67917f3286333d", "source_path": "examples/data-sources/xcsh_infraprotect_mitigation_ips/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:example:data-source", "parent_id": "xcsh-docs:data-sources:infraprotect_mitigation_ips:examples", "path": "documentation/data-sources/infraprotect_mitigation_ips/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "infraprotect_mitigation_ips", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2323222212111121-1020330310021221-2100022220100331-2103120000311113-3220321301201120-0113320021300110-2202332100312021-0113102330213302", "registry_path": "docs/guides/data-sources--infraprotect_mitigation_ips--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/infraprotect_mitigation_ips/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_infraprotect_mitigation_ips.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_infraprotect_mitigation_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_infraprotect_mitigation_ips/data-source.tf`; digest `sha256:0666bca16af2fe6a216b181ad7ffa5f07262937f90b186cb1b67917f3286333d`.

```terraform
# InfraprotectMitigationIps DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_infraprotect_mitigation_ips" "example" {
  mitigation_id = "example-value"
  namespace     = "example-value"
}

output "infraprotect_mitigation_ips_result" {
  value = data.xcsh_infraprotect_mitigation_ips.example
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/examples/)
- [xcsh_infraprotect_mitigation_ips](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/infraprotect_mitigation_ips/)
