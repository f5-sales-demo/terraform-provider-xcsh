---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_image."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1208, "body_sha256": "sha256:ce80a887ff6a46405979849d6b9b0cd4bdc30649d22be74ca64c1a66cc6f9913", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:site_image:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f749f6969ede8da33793449ff8e0b7f889a9d0fc592de94c4d8d386d5af43af8", "source_path": "examples/data-sources/xcsh_site_image/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_image:example:data-source", "parent_id": "xcsh-docs:data-sources:site_image:examples", "path": "documentation/data-sources/site_image/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "site_image", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-1132221030332110-3321221020101213-1111200220022233-2000303203110102-2022010032000233-3111110123220010-1230231003203232-0130102100010330", "registry_path": "docs/guides/data-sources--site_image--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_image/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_site_image.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
