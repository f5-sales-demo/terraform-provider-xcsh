---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_site_image."
xcsh_docs: {"aliases": [], "body_bytes": 1002, "body_sha256": "sha256:6ac411ec2b0533dec293e177183d99cf119d1ca54f80670e849d4738a1efb9a4", "canonical_id": "xcsh-docs:data-sources:site_image:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:site_image:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:f749f6969ede8da33793449ff8e0b7f889a9d0fc592de94c4d8d386d5af43af8", "source_path": "examples/data-sources/xcsh_site_image/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:site_image:example:data-source", "parent_id": "xcsh-docs:data-sources:site_image:examples", "path": "docs/guides/data-sources--site_image--example--data-source.md", "provider_name": "site_image", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/site_image/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_site_image.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_site_image](../data-sources/site_image.md)
- [Examples](data-sources--site_image--examples.md)
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

- [Examples](data-sources--site_image--examples.md)
- [xcsh_site_image](../data-sources/site_image.md)
