---
page_title: "With jitter"
subcategory: "Monitoring"
description: "With jitter for xcsh_healthcheck."
xcsh_docs: {"aliases": ["with-jitter"], "body_bytes": 1131, "body_sha256": "sha256:50da9b8a915460754afb8d2e5a38521d97cd43724ba63259f82258a87c1da09e", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:b066d263997260c6a34ca816fcfdef53d425a0fd6680007fb1a8a2d0ba72d4fe", "source_path": "examples/resources/xcsh_healthcheck/with-jitter.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:healthcheck:example:with-jitter", "parent_id": "xcsh-docs:resources:healthcheck:examples", "path": "documentation/resources/healthcheck/examples/with-jitter/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-1213303100211112-0310221121330321-0300210312002303-1112302300120312-0120211313122313-0120130232221320-0220102330331211-1000013020212032", "registry_path": "docs/guides/resources--healthcheck--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-jitter"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/examples/with-jitter/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "With jitter for xcsh_healthcheck.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["healthcheckCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With jitter

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/examples/)
- With jitter

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/with-jitter.tf`; digest `sha256:b066d263997260c6a34ca816fcfdef53d425a0fd6680007fb1a8a2d0ba72d4fe`.

```terraform
# WithJitter — Acceptance-test-derived Configuration
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
  jitter_percent      = 30

  tcp_health_check {}
}
```
