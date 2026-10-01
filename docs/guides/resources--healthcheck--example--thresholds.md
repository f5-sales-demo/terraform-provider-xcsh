---
page_title: "Thresholds"
subcategory: "Monitoring"
description: "Thresholds for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 1115, "body_sha256": "sha256:a401dd65e1c5c9c33190b26f2e825262e6b6f8b28278932cc114f8d30cfcef8c", "canonical_id": "xcsh-docs:resources:healthcheck:example:thresholds", "child_ids": [], "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:eb7ffc7c4c13b7cf863658813f0b685ce60bc365c197904650a919fd4f6c683f", "source_path": "examples/resources/xcsh_healthcheck/thresholds.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:healthcheck:example:thresholds", "parent_id": "xcsh-docs:resources:healthcheck:examples", "path": "docs/guides/resources--healthcheck--example--thresholds.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["thresholds"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/examples/thresholds/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Thresholds for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Thresholds

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md)
- [Examples](resources--healthcheck--examples.md)
- Thresholds

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/thresholds.tf`; digest `sha256:eb7ffc7c4c13b7cf863658813f0b685ce60bc365c197904650a919fd4f6c683f`.

```terraform
# Thresholds — Acceptance-test-derived Configuration
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

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 3
  unhealthy_threshold = 2
  timeout             = 5
  interval            = 15

  tcp_health_check {}
}
```

## Next pages

- [Examples](resources--healthcheck--examples.md)
- [xcsh_healthcheck](../resources/healthcheck.md)
