---
page_title: "With limits"
subcategory: "Security"
description: "With limits for xcsh_rate_limiter."
xcsh_docs: {"aliases": ["with-limits"], "body_bytes": 1194, "body_sha256": "sha256:19c82e9ae980251e622482f899095c165066f5641edf3a39708bdf928a98dcc5", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:851a79af35fd3d7c3025e31d8945aa3a5098e4fad6d2cccc80ec89bdb883b6c4", "source_path": "examples/resources/xcsh_rate_limiter/with-limits.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:rate_limiter:example:with-limits", "parent_id": "xcsh-docs:resources:rate_limiter:examples", "path": "documentation/resources/rate_limiter/examples/with-limits/index.md", "product": "distributed-cloud", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-3232112030210220-3003003210301013-3122120123023312-2331110222213302-2311113213213331-3103102230123203-2111132330301032-2021333232220310", "registry_path": "docs/guides/resources--rate_limiter--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-limits"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/examples/with-limits/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "With limits for xcsh_rate_limiter.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With limits

Breadcrumbs:

- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/examples/)
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
