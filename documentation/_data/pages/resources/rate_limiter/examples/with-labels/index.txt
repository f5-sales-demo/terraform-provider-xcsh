---
page_title: "With labels"
subcategory: "Security"
description: "With labels for xcsh_rate_limiter."
xcsh_docs: {"aliases": [], "body_bytes": 1156, "body_sha256": "sha256:2ef3b12098cf939760a448ed8664a6eee010486ba4da604b81dedfc611501e70", "child_ids": [], "collection_id": "xcsh-docs:resources:rate_limiter:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:5609348e52d35b34f6d604be53430c08f2fce483ae5ffb0505d915c66d61db51", "source_path": "examples/resources/xcsh_rate_limiter/with-labels.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:rate_limiter:example:with-labels", "parent_id": "xcsh-docs:resources:rate_limiter:examples", "path": "documentation/resources/rate_limiter/examples/with-labels/index.md", "provider_name": "rate_limiter", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["with-labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/rate_limiter/examples/with-labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With labels for xcsh_rate_limiter.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["rate_limiterCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# With labels

Breadcrumbs:

- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/examples/)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_rate_limiter/with-labels.tf`; digest `sha256:5609348e52d35b34f6d604be53430c08f2fce483ae5ffb0505d915c66d61db51`.

```terraform
# WithLabels — Acceptance-test-derived Configuration
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

  labels = {
    example-key = "example-value"
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/examples/)
- [xcsh_rate_limiter](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/rate_limiter/)
