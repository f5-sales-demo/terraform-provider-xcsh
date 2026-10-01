---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": [], "body_bytes": 1071, "body_sha256": "sha256:617c8ebbe2b3054befe6d5dc1571a665cbc6a892ebe121843fe3c4450d310c96", "canonical_id": "xcsh-docs:data-sources:aws_tgw_site:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:aef0f2c4d7539fcdc86d3bd364e4b9a22f4422d4d0595364ea8d26b8fc2a6cf7", "source_path": "examples/data-sources/xcsh_aws_tgw_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:aws_tgw_site:example:data-source", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:examples", "path": "docs/guides/data-sources--aws_tgw_site--example--data-source.md", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_aws_tgw_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
- [Examples](data-sources--aws_tgw_site--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_aws_tgw_site/data-source.tf`; digest `sha256:aef0f2c4d7539fcdc86d3bd364e4b9a22f4422d4d0595364ea8d26b8fc2a6cf7`.

```terraform
# AWSTGWSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AWSTGWSite by name
data "xcsh_aws_tgw_site" "example" {
  name      = "example-aws-tgw-site"
  namespace = "staging"
}

output "aws_tgw_site_id" {
  value = data.xcsh_aws_tgw_site.example.id
}
```

## Next pages

- [Examples](data-sources--aws_tgw_site--examples.md)
- [xcsh_aws_tgw_site](../data-sources/aws_tgw_site.md)
