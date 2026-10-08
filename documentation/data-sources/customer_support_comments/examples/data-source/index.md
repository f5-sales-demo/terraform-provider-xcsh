---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_customer_support_comments."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1077, "body_sha256": "sha256:c6ae79dd50dae94018c56b3bc28f848794966f692bc6a6ba8dacbff1191ee46f", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:customer_support_comments:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0d5599711e86330f779d8085568282cccbd3e192a937997854401c5d7fbdb6a3", "source_path": "examples/data-sources/xcsh_customer_support_comments/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:customer_support_comments:example:data-source", "parent_id": "xcsh-docs:data-sources:customer_support_comments:examples", "path": "documentation/data-sources/customer_support_comments/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "customer_support_comments", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-2332130320302213-2223033010100300-2011303322211302-0303123320302333-0320133332211322-2223010030013210-3130110200121011-1123323322213321", "registry_path": "docs/guides/data-sources--customer_support_comments--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/customer_support_comments/examples/data-source/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Data source for xcsh_customer_support_comments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": [], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_customer_support_comments/data-source.tf`; digest `sha256:0d5599711e86330f779d8085568282cccbd3e192a937997854401c5d7fbdb6a3`.

```terraform
# CustomerSupportComments DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_customer_support_comments" "example" {
  name = "example-value"
}

output "customer_support_comments_result" {
  value = data.xcsh_customer_support_comments.example
}
```
