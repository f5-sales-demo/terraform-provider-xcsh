---
page_title: "Resource"
subcategory: "Identity"
description: "Resource for xcsh_token."
xcsh_docs: {"aliases": ["resource"], "body_bytes": 1175, "body_sha256": "sha256:5a30a46aa8633657e7426fa18a5065387e16decdb0d452a70b698a4301dbc8b6", "capabilities": ["identity"], "category": "identity", "child_ids": [], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:token:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:e34df2fe171a3a579b8dd3181a5ec00f7f49f2449740ae1edcf8cdaec5c57bb9", "source_path": "examples/resources/xcsh_token/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:token:example:resource", "parent_id": "xcsh-docs:resources:token:examples", "path": "documentation/resources/token/examples/resource/index.md", "product": "distributed-cloud", "provider_name": "token", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2101330310331202-0332121202233333-3312120320221103-2222113332330330-0203101000211121-0112012123211223-2321302231333312-3120102111310332", "registry_path": "docs/guides/resources--token--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["resource"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/token/examples/resource/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Resource for xcsh_token.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["tokenCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
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
