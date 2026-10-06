---
page_title: "Http http2"
subcategory: "Monitoring"
description: "Http http2 for xcsh_healthcheck."
xcsh_docs: {"aliases": ["http-http2"], "body_bytes": 1193, "body_sha256": "sha256:2bd3c7b5b31e10cdcf92d602d67aa4f1faa8aac7be323fd6324c5be030231427", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:7586adcf4b2cb6423fa64e91dd0205c03d4130e45c575432d8fff8d1cf0d1e8f", "source_path": "examples/resources/xcsh_healthcheck/http-http2.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:healthcheck:example:http-http2", "parent_id": "xcsh-docs:resources:healthcheck:examples", "path": "documentation/resources/healthcheck/examples/http-http2/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-1310130002020132-0022010320210002-3021310133232030-0222300211223012-3231010102122112-0322031231013030-1321313213113031-2020333321313202", "registry_path": "docs/guides/resources--healthcheck--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["http-http2"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/examples/http-http2/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "Http http2 for xcsh_healthcheck.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["healthcheckCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Http http2

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/examples/)
- Http http2

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_healthcheck/http-http2.tf`; digest `sha256:7586adcf4b2cb6423fa64e91dd0205c03d4130e45c575432d8fff8d1cf0d1e8f`.

```terraform
# HttpHttp2 — Acceptance-test-derived Configuration
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
    path        = "example-value"
    host_header = "example.com"
    use_http2   = true
  }
}
```
