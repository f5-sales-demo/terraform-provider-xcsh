---
page_title: "Data source"
subcategory: "Infrastructure"
description: "Data source for xcsh_aws_vpc_site."
xcsh_docs: {"aliases": [], "body_bytes": 1277, "body_sha256": "sha256:3480386c1ee22ee8ce4f1a65bad664f52923fa71a8942eabfb68da427c6c0932", "child_ids": [], "collection_id": "xcsh-docs:data-sources:aws_vpc_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:255a62b4362db5d970e6f776cb0dcab3af764d3d9fda6ded76847f22ef47dd92", "source_path": "examples/data-sources/xcsh_aws_vpc_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:aws_vpc_site:example:data-source", "parent_id": "xcsh-docs:data-sources:aws_vpc_site:examples", "path": "documentation/data-sources/aws_vpc_site/examples/data-source/index.md", "provider_name": "aws_vpc_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_vpc_site/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_aws_vpc_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["aws_vpc_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_aws_vpc_site/data-source.tf`; digest `sha256:255a62b4362db5d970e6f776cb0dcab3af764d3d9fda6ded76847f22ef47dd92`.

```terraform
# AWSVPCSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AWSVPCSite by name
data "xcsh_aws_vpc_site" "example" {
  name      = "example-aws-vpc-site"
  namespace = "staging"
}

output "aws_vpc_site_id" {
  value = data.xcsh_aws_vpc_site.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/examples/)
- [xcsh_aws_vpc_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_vpc_site/)
