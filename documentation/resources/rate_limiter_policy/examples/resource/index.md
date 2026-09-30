---
page_title: "Resource"
subcategory: "Security"
description: "Resource for xcsh_rate_limiter_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1261, "body_sha256": "sha256:a65ef1898f3f89a1f1d14e7cf1a2808ebc22698aebe6436a7b6da5da968ca394", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter_policy:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:d8a1a676828e8f463f6383201f395371ac9e0d89d99902337c84f2ee8a150c35", "source_path": "examples/resources/xcsh_rate_limiter_policy/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:rate_limiter_policy:example:resource", "parent_id": "xcsh-docs:resources:rate_limiter_policy:examples", "path": "documentation/resources/rate_limiter_policy/examples/resource/index.md", "provider_name": "rate_limiter_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter_policy/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_rate_limiter_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiter_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/examples/)
- [xcsh_rate_limiter_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter_policy/)
