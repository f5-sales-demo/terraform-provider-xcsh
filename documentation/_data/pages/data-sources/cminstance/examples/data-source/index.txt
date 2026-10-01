---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cminstance."
xcsh_docs: {"aliases": [], "body_bytes": 1255, "body_sha256": "sha256:eacef67085860bda96847e14efd433dfc2f5ab01039589930819108b7fb0ab17", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cminstance:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:833c4b202ec741b02e9ef290bb6512a37d3156dd48242949d7ad995d973193c7", "source_path": "examples/data-sources/xcsh_cminstance/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cminstance:example:data-source", "parent_id": "xcsh-docs:data-sources:cminstance:examples", "path": "documentation/data-sources/cminstance/examples/data-source/index.md", "provider_name": "cminstance", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cminstance/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_cminstance.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cminstanceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cminstance/data-source.tf`; digest `sha256:833c4b202ec741b02e9ef290bb6512a37d3156dd48242949d7ad995d973193c7`.

```terraform
# Cminstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Cminstance by name
data "xcsh_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "staging"
}

output "cminstance_id" {
  value = data.xcsh_cminstance.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/examples/)
- [xcsh_cminstance](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cminstance/)
