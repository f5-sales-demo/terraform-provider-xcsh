---
page_title: "All attributes"
subcategory: "Security"
description: "All attributes for xcsh_rate_limiter."
xcsh_docs: {"aliases": [], "body_bytes": 1213, "body_sha256": "sha256:3def15dac9766e3a752790aa1a2a3541d7d8d4e37044e2fc5759d7ab8897614f", "canonical_id": "xcsh-docs:resources:rate_limiter:example:all-attributes", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:0fa6c0ef5d6c3a83afac2353f4434ae5a9fdbf3ec95f2509bae99373fcdd2e92", "source_path": "examples/resources/xcsh_rate_limiter/all-attributes.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:rate_limiter:example:all-attributes", "parent_id": "xcsh-docs:resources:rate_limiter:examples", "path": "docs/guides/resources--rate_limiter--example--all-attributes.md", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["all-attributes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/examples/all-attributes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "All attributes for xcsh_rate_limiter.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# All attributes

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md)
- [Examples](resources--rate_limiter--examples.md)
- All attributes

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/all-attributes.tf`; digest `sha256:0fa6c0ef5d6c3a83afac2353f4434ae5a9fdbf3ec95f2509bae99373fcdd2e92`.

```terraform
# AllAttributes — Acceptance-test-derived Configuration
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
  description = "Test rate limiter with all attributes"
  disable     = false

  labels = {
    environment = "test"
    team        = "engineering"
  }

  annotations = {
    purpose = "testing"
  }
}
```

## Next pages

- [Examples](resources--rate_limiter--examples.md)
- [xcsh_rate_limiter](../resources/rate_limiter.md)
