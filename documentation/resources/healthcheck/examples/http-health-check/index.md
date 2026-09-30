---
page_title: "Http health check"
subcategory: "Monitoring"
description: "Http health check for xcsh_healthcheck."
xcsh_docs: {"aliases": [], "body_bytes": 1311, "body_sha256": "sha256:a27546a9a32209fe5892652d82e2e733f2dfcb5d8c4267ade003e982b40102bc", "child_ids": [], "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:d687dc5e245ad34528fd20feda8e0150dd05ca683b1bddfd84c19f770f54453d", "source_path": "examples/resources/xcsh_healthcheck/http-health-check.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:healthcheck:example:http-health-check", "parent_id": "xcsh-docs:resources:healthcheck:examples", "path": "documentation/resources/healthcheck/examples/http-health-check/index.md", "provider_name": "healthcheck", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["http-health-check"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/examples/http-health-check/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Http health check for xcsh_healthcheck.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["healthcheckCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Http health check

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/examples/)
- Http health check

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/http-health-check.tf`; digest `sha256:d687dc5e245ad34528fd20feda8e0150dd05ca683b1bddfd84c19f770f54453d`.

```terraform
# HttpHealthCheck — Acceptance-test-derived Configuration
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

  http_health_check {
    path        = "/health"
    host_header = "example.com"
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/examples/)
- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
