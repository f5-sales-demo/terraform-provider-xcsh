---
page_title: "Data source"
subcategory: "Load Balancing"
description: "Data source for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1060, "body_sha256": "sha256:5a4f3d94004e8dd33202d50b8bae13d255f2322ef1aacd5ea5ccb68b3928afbe", "canonical_id": "xcsh-docs:data-sources:origin_pool:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:cae409ea2c1cde6f7ffac297039f556a77e9684a94017d350d26cb02cd85782a", "source_path": "examples/data-sources/xcsh_origin_pool/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:origin_pool:example:data-source", "parent_id": "xcsh-docs:data-sources:origin_pool:examples", "path": "docs/guides/data-sources--origin_pool--example--data-source.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_origin_pool](../data-sources/origin_pool.md)
- [Examples](data-sources--origin_pool--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_origin_pool/data-source.tf`; digest `sha256:cae409ea2c1cde6f7ffac297039f556a77e9684a94017d350d26cb02cd85782a`.

```terraform
# OriginPool Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing OriginPool by name
data "xcsh_origin_pool" "example" {
  name      = "example-origin-pool"
  namespace = "staging"
}

output "origin_pool_id" {
  value = data.xcsh_origin_pool.example.id
}
```

## Next pages

- [Examples](data-sources--origin_pool--examples.md)
- [xcsh_origin_pool](../data-sources/origin_pool.md)
