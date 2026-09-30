---
page_title: "Resource"
subcategory: "Load Balancing"
description: "Resource for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1182, "body_sha256": "sha256:24c43ad03bdf05ae728d91b4d9edad17d53e3a46ac04a467d0d59df5df20df51", "child_ids": [], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d9f0fe93931beb085130f6173ace667305af3adbe1e9448009e07911a7af8fa0", "source_path": "examples/resources/xcsh_origin_pool/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:origin_pool:example:resource", "parent_id": "xcsh-docs:resources:origin_pool:examples", "path": "documentation/resources/origin_pool/examples/resource/index.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_origin_pool/resource.tf`; digest `sha256:d9f0fe93931beb085130f6173ace667305af3adbe1e9448009e07911a7af8fa0`.

```terraform
# OriginPool Resource Example
# Manages a Origin Pool resource in F5 Distributed Cloud for defining backend server pools for load balancer targets.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic OriginPool configuration
resource "xcsh_origin_pool" "example" {
  name      = "example-origin-pool"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/examples/)
- [xcsh_origin_pool](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/)
