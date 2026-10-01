---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_customer_support_comments."
xcsh_docs: {"aliases": [], "body_bytes": 1138, "body_sha256": "sha256:66874095cd2bec61659014eb7cf67e85bb5119764727dc958df9c5c035035221", "canonical_id": "xcsh-docs:data-sources:customer_support_comments:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:customer_support_comments:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:0d5599711e86330f779d8085568282cccbd3e192a937997854401c5d7fbdb6a3", "source_path": "examples/data-sources/xcsh_customer_support_comments/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:customer_support_comments:example:data-source", "parent_id": "xcsh-docs:data-sources:customer_support_comments:examples", "path": "docs/guides/data-sources--customer_support_comments--example--data-source.md", "provider_name": "customer_support_comments", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/customer_support_comments/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_customer_support_comments.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md)
- [Examples](data-sources--customer_support_comments--examples.md)
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

- [Examples](data-sources--customer_support_comments--examples.md)
- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md)
