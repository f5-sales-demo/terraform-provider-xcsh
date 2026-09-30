---
page_title: "With labels"
subcategory: "Monitoring"
description: "With labels for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 1110, "body_sha256": "sha256:c1e1727dd5874e75fc45345441771a5d04f66d1525bb66229b17a52e12e6faf9", "canonical_id": "xcsh-docs:resources:healthcheck:example:with-labels", "child_ids": [], "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:9d101becb63d5e10a0b0292248bcfa96ec4395ca9743eb0d33275bb701472643", "source_path": "examples/resources/xcsh_healthcheck/with-labels.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:healthcheck:example:with-labels", "parent_id": "xcsh-docs:resources:healthcheck:examples", "path": "docs/guides/resources--healthcheck--example--with-labels.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["with-labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/examples/with-labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With labels for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# With labels

Breadcrumbs:

- [xcsh_healthcheck](../resources/healthcheck.md)
- [Examples](resources--healthcheck--examples.md)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/with-labels.tf`; digest `sha256:9d101becb63d5e10a0b0292248bcfa96ec4395ca9743eb0d33275bb701472643`.

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

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  labels = {
    environment = "example-value"
    managed_by  = "example-description"
  }

  tcp_health_check {}
}
```

## Next pages

- [Examples](resources--healthcheck--examples.md)
- [xcsh_healthcheck](../resources/healthcheck.md)
