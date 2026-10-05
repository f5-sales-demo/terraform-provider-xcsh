---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1277, "body_sha256": "sha256:5dc00d960ea19c70b200b745e7e1c499fe60e1c30165e2f2a569b73874486a90", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:aef0f2c4d7539fcdc86d3bd364e4b9a22f4422d4d0595364ea8d26b8fc2a6cf7", "source_path": "examples/data-sources/xcsh_aws_tgw_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:aws_tgw_site:example:data-source", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:examples", "path": "documentation/data-sources/aws_tgw_site/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-2010230202333132-1032132311330303-0132211003312131-1312212020202102-2123320123201032-2003002020012310-0311221100231331-0323201131232302", "registry_path": "docs/guides/data-sources--aws_tgw_site--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/examples/data-source/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Data source for xcsh_aws_tgw_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/examples/)
- [xcsh_aws_tgw_site](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/aws_tgw_site/)
