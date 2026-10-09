---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_aws_tgw_site."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1049, "body_sha256": "sha256:506be9c8d5325dbd2f7a5e2c60958e7699cdedb9906fcde22826aceddbcbe81f", "capabilities": ["infrastructure"], "category": "infrastructure", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:aws_tgw_site:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:aef0f2c4d7539fcdc86d3bd364e4b9a22f4422d4d0595364ea8d26b8fc2a6cf7", "source_path": "examples/data-sources/xcsh_aws_tgw_site/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:aws_tgw_site:example:data-source", "parent_id": "xcsh-docs:data-sources:aws_tgw_site:examples", "path": "documentation/data-sources/aws_tgw_site/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "aws_tgw_site", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "data-sources", "registry_anchor": "canonical-2010230202333132-1032132311330303-0132211003312131-1312212020202102-2123320123201032-2003002020012310-0311221100231331-0323201131232302", "registry_path": "docs/guides/data-sources--aws_tgw_site--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/aws_tgw_site/examples/data-source/index.txt", "spec_pin_digest": "sha256:e06a3ea9db6a533295efd5c7a477afc65990ba80c3a998885b97b47262cfe9f1", "summary": "Data source for xcsh_aws_tgw_site.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.4", "schema_components": ["aws_tgw_siteCreateRequest"], "target_commit": "c5ce81d5fb15314a0f9398db954e0da111d89606"}}
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
