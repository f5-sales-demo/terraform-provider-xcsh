---
page_title: "Resource"
subcategory: "Identity"
description: "Resource for xcsh_token."
xcsh_docs: {"aliases": [], "body_bytes": 1175, "body_sha256": "sha256:5a30a46aa8633657e7426fa18a5065387e16decdb0d452a70b698a4301dbc8b6", "child_ids": [], "collection_id": "xcsh-docs:resources:token:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e34df2fe171a3a579b8dd3181a5ec00f7f49f2449740ae1edcf8cdaec5c57bb9", "source_path": "examples/resources/xcsh_token/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:token:example:resource", "parent_id": "xcsh-docs:resources:token:examples", "path": "documentation/resources/token/examples/resource/index.md", "provider_name": "token", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/token/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_token.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tokenCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/token/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/token/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_token/resource.tf`; digest `sha256:e34df2fe171a3a579b8dd3181a5ec00f7f49f2449740ae1edcf8cdaec5c57bb9`.

```terraform
# Token Resource Example
# Manages new token.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Token configuration
resource "xcsh_token" "example" {
  name      = "example-token"
  namespace = "system"
  type      = 1
  site_name = "example-securemesh-site"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/token/examples/)
- [xcsh_token](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/token/)
