---
page_title: "All attributes"
subcategory: "Monitoring"
description: "All attributes for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 1405, "body_sha256": "sha256:90efaed8baa7882f4af65db5f03151383c2b900a716cb8147da9c23c2e746565", "child_ids": [], "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:7b4a2a1e5d67895b93aa7dcead7285790016b97f45b00fc3456985658bd0bba4", "source_path": "examples/resources/xcsh_healthcheck/all-attributes.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:healthcheck:example:all-attributes", "parent_id": "xcsh-docs:resources:healthcheck:examples", "path": "documentation/resources/healthcheck/examples/all-attributes/index.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["all-attributes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/examples/all-attributes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "All attributes for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# All attributes

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/examples/)
- All attributes

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/all-attributes.tf`; digest `sha256:7b4a2a1e5d67895b93aa7dcead7285790016b97f45b00fc3456985658bd0bba4`.

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

resource "xcsh_healthcheck" "test" {
  name      = "example"
  namespace = "system"

  healthy_threshold   = 1
  unhealthy_threshold = 2
  timeout             = 3
  interval            = 5

  labels = {
    environment = "test"
    managed_by  = "terraform-acceptance-test"
  }

  annotations = {
    purpose = "acceptance-testing"
    owner   = "ci-cd"
  }

  tcp_health_check {}
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/examples/)
- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
