---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_image."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1208, "body_sha256": "sha256:ce80a887ff6a46405979849d6b9b0cd4bdc30649d22be74ca64c1a66cc6f9913", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_image:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f749f6969ede8da33793449ff8e0b7f889a9d0fc592de94c4d8d386d5af43af8", "source_path": "examples/data-sources/xcsh_site_image/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_image:example:data-source", "parent_id": "xcsh-docs:data-sources:site_image:examples", "path": "documentation/data-sources/site_image/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "site_image", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1132221030332110-3321221020101213-1111200220022233-2000303203110102-2022010032000233-3111110123220010-1230231003203232-0130102100010330", "registry_path": "docs/guides/data-sources--site_image--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_image/examples/data-source/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Data source for xcsh_site_image.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_image](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_site_image/data-source.tf`; digest `sha256:f749f6969ede8da33793449ff8e0b7f889a9d0fc592de94c4d8d386d5af43af8`.

```terraform
# SiteImage DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_site_image" "example" {
  site_name = "example-value"
}

output "site_image_result" {
  value     = data.xcsh_site_image.example
  sensitive = true
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/examples/)
- [xcsh_site_image](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/site_image/)
