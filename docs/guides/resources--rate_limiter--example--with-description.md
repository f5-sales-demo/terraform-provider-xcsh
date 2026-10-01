---
page_title: "With description"
subcategory: "Security"
description: "With description for xcsh_rate_limiter."
xcsh_docs: {"aliases": [], "body_bytes": 1053, "body_sha256": "sha256:9d900e64038b41e9caa494a38e4000284239b5461eaaa71dc7641f03b906fcc7", "canonical_id": "xcsh-docs:resources:rate_limiter:example:with-description", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:87a15fe7b89798f0857db8053a2641296d700d8b857fd78fbb80b8d1c50c50a9", "source_path": "examples/resources/xcsh_rate_limiter/with-description.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:rate_limiter:example:with-description", "parent_id": "xcsh-docs:resources:rate_limiter:examples", "path": "docs/guides/resources--rate_limiter--example--with-description.md", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["with-description"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/examples/with-description/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With description for xcsh_rate_limiter.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With description

Breadcrumbs:

- [xcsh_rate_limiter](../resources/rate_limiter.md)
- [Examples](resources--rate_limiter--examples.md)
- With description

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/with-description.tf`; digest `sha256:87a15fe7b89798f0857db8053a2641296d700d8b857fd78fbb80b8d1c50c50a9`.

```terraform
# WithDescription — Acceptance-test-derived Configuration
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
  description = "example-value"
}
```

## Next pages

- [Examples](resources--rate_limiter--examples.md)
- [xcsh_rate_limiter](../resources/rate_limiter.md)
