---
page_title: "With limits"
subcategory: "Security"
description: "With limits for xcsh_rate_limiter."
xcsh_docs: {"aliases": [], "body_bytes": 1210, "body_sha256": "sha256:a0f6e420a156301a0f03dba07b11c8c88f21ba50e6028f27d70b6624f26bd565", "canonical_id": "xcsh-docs:resources:rate_limiter:example:with-limits", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:851a79af35fd3d7c3025e31d8945aa3a5098e4fad6d2cccc80ec89bdb883b6c4", "source_path": "examples/resources/xcsh_rate_limiter/with-limits.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:rate_limiter:example:with-limits", "parent_id": "xcsh-docs:resources:rate_limiter:examples", "path": "docs/guides/resources--rate_limiter--example--with-limits.md", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["with-limits"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/examples/with-limits/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With limits for xcsh_rate_limiter.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With limits

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md)
- [Examples](resources--rate_limiter--examples.md)
- With limits

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/with-limits.tf`; digest `sha256:851a79af35fd3d7c3025e31d8945aa3a5098e4fad6d2cccc80ec89bdb883b6c4`.

```terraform
# WithLimits — Acceptance-test-derived Configuration
# Extracted from an acceptance test helper.
# No new live API validation is claimed.

terraform {
  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

resource "xcsh_rate_limiter" "test" {
  name        = "example"
  namespace   = "system"
  description = "Rate limiter with limits configuration"

  limits {
    total_number      = 100
    unit              = "MINUTE"
    burst_multiplier  = 2
    period_multiplier = 1

    leaky_bucket = {}
  }
}
```

## Next pages

- [Examples](resources--rate_limiter--examples.md)
- [xcsh_rate_limiter](../resources/rate_limiter.md)
