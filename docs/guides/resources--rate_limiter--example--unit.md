---
page_title: "Unit"
subcategory: "Security"
description: "Unit for xcsh_rate_limiter."
xcsh_docs: {"aliases": [], "body_bytes": 1091, "body_sha256": "sha256:ce7f2df8103687b901743c4a0cf8035ab1075bc86495c8df3509599923281314", "canonical_id": "xcsh-docs:resources:rate_limiter:example:unit", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:93b1c54c089411766714ddbc2d2a6c1769c9446b2d2ef86dfb9ea72876548c98", "source_path": "examples/resources/xcsh_rate_limiter/unit.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:rate_limiter:example:unit", "parent_id": "xcsh-docs:resources:rate_limiter:examples", "path": "docs/guides/resources--rate_limiter--example--unit.md", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["unit"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/examples/unit/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Unit for xcsh_rate_limiter.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Unit

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md)
- [Examples](resources--rate_limiter--examples.md)
- Unit

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/unit.tf`; digest `sha256:93b1c54c089411766714ddbc2d2a6c1769c9446b2d2ef86dfb9ea72876548c98`.

```terraform
# Unit — Acceptance-test-derived Configuration
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
  name      = "example"
  namespace = "system"

  limits {
    total_number     = 3
    unit             = "MINUTE"
    burst_multiplier = 2

    leaky_bucket = {}
  }
}
```

## Next pages

- [Examples](resources--rate_limiter--examples.md)
- [xcsh_rate_limiter](../resources/rate_limiter.md)
