---
page_title: "xcsh_origin_pool"
subcategory: "Load Balancing"
description: "xcsh_origin_pool for xcsh_origin_pool."
xcsh_docs: {"aliases": [], "body_bytes": 1723, "body_sha256": "sha256:24e0f3f8dc65401e7a63d0b886cdf4146c559373d8e12c3cf74a6232aeb8dede", "child_ids": ["xcsh-docs:resources:origin_pool:reference", "xcsh-docs:resources:origin_pool:examples", "xcsh-docs:resources:origin_pool:import", "xcsh-docs:resources:origin_pool:timeouts"], "collection_id": "xcsh-docs:resources:origin_pool:collection", "completeness": "complete", "id": "xcsh-docs:resources:origin_pool:fundamentals", "parent_id": null, "path": "documentation/resources/origin_pool/index.md", "provider_name": "origin_pool", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/origin_pool/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_origin_pool for xcsh_origin_pool.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["origin_poolCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/origin_pool/lifecycle/timeouts/)
