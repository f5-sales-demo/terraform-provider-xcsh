---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_data_group."
xcsh_docs: {"aliases": [], "body_bytes": 1195, "body_sha256": "sha256:bee7c8651cfb517fe95c2225c3721f410dcd3389a6b919beb86c133c4af34888", "child_ids": [], "collection_id": "xcsh-docs:resources:data_group:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:1bceaa79a1e1b840495f41f1bf5adac4d3dc173e5487643b627a4cc77fc7f690", "source_path": "examples/resources/xcsh_data_group/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:data_group:example:resource", "parent_id": "xcsh-docs:resources:data_group:examples", "path": "documentation/resources/data_group/examples/resource/index.md", "provider_name": "data_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_group/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_data_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_data_group/resource.tf`; digest `sha256:1bceaa79a1e1b840495f41f1bf5adac4d3dc173e5487643b627a4cc77fc7f690`.

```terraform
# DataGroup Resource Example
# Manages data group in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataGroup configuration
resource "xcsh_data_group" "example" {
  name      = "example-data-group"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/examples/)
- [xcsh_data_group](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_group/)
