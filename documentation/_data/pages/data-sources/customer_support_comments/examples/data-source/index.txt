---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_customer_support_comments."
xcsh_docs: {"aliases": ["data-source"], "body_bytes": 1344, "body_sha256": "sha256:6e2c5fa8add1b4d2653365c6be34c604e0c51e741dafd77e64d25097ce5d37c9", "capabilities": [], "category": null, "child_ids": [], "classification": {"rules_sha256": "sha256:98dbf5280bf8376a0c75c1391250db9db4ac0d293d42bdd831c20f7514925adb", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:customer_support_comments:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0d5599711e86330f779d8085568282cccbd3e192a937997854401c5d7fbdb6a3", "source_path": "examples/data-sources/xcsh_customer_support_comments/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:customer_support_comments:example:data-source", "parent_id": "xcsh-docs:data-sources:customer_support_comments:examples", "path": "documentation/data-sources/customer_support_comments/examples/data-source/index.md", "product": "distributed-cloud", "provider_name": "customer_support_comments", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2332130320302213-2223033010100300-2011303322211302-0303123320302333-0320133332211322-2223010030013210-3130110200121011-1123323322213321", "registry_path": "docs/guides/data-sources--customer_support_comments--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["data-source"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/customer_support_comments/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_customer_support_comments.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/examples/)
- [xcsh_customer_support_comments](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/customer_support_comments/)
