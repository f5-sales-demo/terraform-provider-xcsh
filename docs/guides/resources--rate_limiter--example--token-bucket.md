---
page_title: "Token bucket"
subcategory: "Security"
description: "Token bucket for xcsh_rate_limiter."
xcsh_docs: {"aliases": [], "body_bytes": 1024, "body_sha256": "sha256:b5d007b947fa27165b25abced7a07d512ef236ac6de3a99752cafd86f5e764c6", "canonical_id": "xcsh-docs:resources:rate_limiter:example:token-bucket", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:12c5b85092b1f6303b0c05874db596bd776622def34147f387a7ecad55d82778", "source_path": "examples/resources/xcsh_rate_limiter/token-bucket.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:rate_limiter:example:token-bucket", "parent_id": "xcsh-docs:resources:rate_limiter:examples", "path": "docs/guides/resources--rate_limiter--example--token-bucket.md", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["token-bucket"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/examples/token-bucket/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Token bucket for xcsh_rate_limiter.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Token bucket

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md)
- [Examples](resources--rate_limiter--examples.md)
- Token bucket

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/token-bucket.tf`; digest `sha256:12c5b85092b1f6303b0c05874db596bd776622def34147f387a7ecad55d82778`.

```terraform
# TokenBucket — Acceptance-test-derived Configuration
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
    total_number     = 50
    unit             = "SECOND"
    burst_multiplier = 5

    token_bucket = {}
  }
}
```

## Next pages

- [Examples](resources--rate_limiter--examples.md)
- [xcsh_rate_limiter](../resources/rate_limiter.md)
