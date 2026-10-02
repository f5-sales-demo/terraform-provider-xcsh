---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_device_intelligence_summary."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1371, "body_sha256": "sha256:eef77a40265907fbdc39a98b4ceb29a74a8ca7dbc3a7e8a3cadbc0472fab3cfd", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:device_intelligence_summary:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a3b5af58871245c4f3794bc2ed886a4421185806d4cb6b7c6eb5328aa94877da", "source_path": "examples/data-sources/xcsh_device_intelligence_summary/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:device_intelligence_summary:example:data-source", "parent_id": "xcsh-docs:data-sources:device_intelligence_summary:examples", "path": "documentation/data-sources/device_intelligence_summary/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "device_intelligence_summary", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3202102222333320-1312230102022103-1231223221131112-3103012033022012-2020131322201223-0011302202233301-1302130221202113-1302013032011121", "registry_path": "docs/guides/data-sources--device_intelligence_summary--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/device_intelligence_summary/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_device_intelligence_summary.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_device_intelligence_summary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_summary/data-source.tf`; digest `sha256:a3b5af58871245c4f3794bc2ed886a4421185806d4cb6b7c6eb5328aa94877da`.

```terraform
# DeviceIntelligenceSummary DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_summary" "example" {
  namespace = "example-value"
}

output "device_intelligence_summary_result" {
  value = data.xcsh_device_intelligence_summary.example
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/examples/)
- [xcsh_device_intelligence_summary](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/device_intelligence_summary/)
