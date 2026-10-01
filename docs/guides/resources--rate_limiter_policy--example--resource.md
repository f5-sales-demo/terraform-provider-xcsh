---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1154, "body_sha256": "sha256:3087ba9589050a44c7df0cbd40761251b029b086077101d8fbb982fd40b7143f", "canonical_id": "xcsh-docs:resources:rate_limiter_policy:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d8a1a676828e8f463f6383201f395371ac9e0d89d99902337c84f2ee8a150c35", "source_path": "examples/resources/xcsh_rate_limiter_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:rate_limiter_policy:example:resource", "parent_id": "xcsh-docs:resources:rate_limiter_policy:examples", "path": "docs/guides/resources--rate_limiter_policy--example--resource.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
- [Examples](resources--rate_limiter_policy--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter_policy/resource.tf`; digest `sha256:d8a1a676828e8f463f6383201f395371ac9e0d89d99902337c84f2ee8a150c35`.

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

## Next pages

- [Examples](resources--rate_limiter_policy--examples.md)
- [xcsh_rate_limiter_policy](../resources/rate_limiter_policy.md)
