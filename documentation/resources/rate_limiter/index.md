---
page_title: "xcsh_rate_limiter"
subcategory: "Security"
description: "xcsh_rate_limiter for xcsh_rate_limiter."
xcsh_docs: {"aliases": [], "body_bytes": 1744, "body_sha256": "sha256:3ddfb96e52784759780eb85643b6e313a1ac710cf970104d9b7ab783a7e5959e", "child_ids": ["xcsh-docs:resources:rate_limiter:reference", "xcsh-docs:resources:rate_limiter:examples", "xcsh-docs:resources:rate_limiter:import", "xcsh-docs:resources:rate_limiter:timeouts"], "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter:fundamentals", "parent_id": null, "path": "documentation/resources/rate_limiter/index.md", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_rate_limiter for xcsh_rate_limiter.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_rate_limiter

Breadcrumbs:

- xcsh_rate_limiter

Manages rate\_limiter creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `rate_limiter_policy`.

- rate_limiter_policy: Detailed rate limiting rules

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# RateLimiter Resource Example
# Manages rate_limiter creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic RateLimiter configuration
resource "xcsh_rate_limiter" "example" {
  name      = "example-rate-limiter"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/lifecycle/timeouts/)
