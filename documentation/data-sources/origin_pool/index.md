---
page_title: "xcsh_origin_pool"
subcategory: "Load Balancing"
description: "xcsh_origin_pool for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1470, "body_sha256": "sha256:d9a5842cfca2c8bdd1d8174e1512a1907c7f0080f78f377c78fd1f7b36cf5369", "child_ids": ["xcsh-docs:data-sources:origin_pool:reference", "xcsh-docs:data-sources:origin_pool:examples"], "collection_id": "xcsh-docs:data-sources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:origin_pool:fundamentals", "parent_id": null, "path": "documentation/data-sources/origin_pool/index.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/origin_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_origin_pool for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_origin_pool

Breadcrumbs:

- xcsh_origin_pool

Manages a Origin Pool resource in F5 Distributed Cloud for defining backend server pools for load
balancer targets.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `healthcheck`.

- healthcheck: Monitor origin server health

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/origin_pool/examples/)
