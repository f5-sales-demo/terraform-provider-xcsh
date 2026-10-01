---
page_title: "xcsh_rate_limiter_policy"
subcategory: "Security"
description: "xcsh_rate_limiter_policy for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1718, "body_sha256": "sha256:7319c02dd7b600089292ae6502313fcbfbe9aa368219e872bbb6170f7ba3db92", "child_ids": ["xcsh-docs:resources:rate_limiter_policy:reference", "xcsh-docs:resources:rate_limiter_policy:examples", "xcsh-docs:resources:rate_limiter_policy:import", "xcsh-docs:resources:rate_limiter_policy:timeouts"], "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "id": "xcsh-docs:resources:rate_limiter_policy:fundamentals", "parent_id": null, "path": "documentation/resources/rate_limiter_policy/index.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_rate_limiter_policy for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_rate_limiter_policy

Breadcrumbs:

- xcsh_rate_limiter_policy

Manages a Rate Limiter Policy resource in F5 Distributed Cloud for rate limiter policy create
specification. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# RateLimiterPolicy Resource Example
# Manages a Rate Limiter Policy resource in F5 Distributed Cloud for rate limiter policy create specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic RateLimiterPolicy configuration
resource "xcsh_rate_limiter_policy" "example" {
  name      = "example-rate-limiter-policy"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/lifecycle/timeouts/)
