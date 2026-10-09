---
page_title: "Http health check"
subcategory: "Monitoring"
description: "Http health check for xcsh_healthcheck."
xcsh_docs: {"aliases": ["http-health-check"], "body_bytes": 1191, "body_sha256": "sha256:7e554d175971195895d53ac5b80562aef06a19dc11668522ff8994e44f18f873", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:d687dc5e245ad34528fd20feda8e0150dd05ca683b1bddfd84c19f770f54453d", "source_path": "examples/resources/xcsh_healthcheck/http-health-check.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:healthcheck:example:http-health-check", "parent_id": "xcsh-docs:resources:healthcheck:examples", "path": "documentation/resources/healthcheck/examples/http-health-check/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-0112102133233221-2002233010320103-1032320003320102-3313223330322100-2130230101013030-1031011312021212-1021223231113210-0331233020020330", "registry_path": "docs/guides/resources--healthcheck--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["http-health-check"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/examples/http-health-check/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Http health check for xcsh_healthcheck.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["healthcheckCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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
