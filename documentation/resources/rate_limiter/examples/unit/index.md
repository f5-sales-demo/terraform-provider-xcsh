---
page_title: "Unit"
subcategory: "Security"
description: "Unit for xcsh_rate_limiter."
xcsh_docs: {"aliases": [], "body_bytes": 1198, "body_sha256": "sha256:fdc5d3021c057b8e3c324fd9b624f08575f6c9fa332d4bd5e56cd2268dbf383b", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:93b1c54c089411766714ddbc2d2a6c1769c9446b2d2ef86dfb9ea72876548c98", "source_path": "examples/resources/xcsh_rate_limiter/unit.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:rate_limiter:example:unit", "parent_id": "xcsh-docs:resources:rate_limiter:examples", "path": "documentation/resources/rate_limiter/examples/unit/index.md", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["unit"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/examples/unit/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Unit for xcsh_rate_limiter.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Unit

Breadcrumbs:

- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/examples/)
- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
