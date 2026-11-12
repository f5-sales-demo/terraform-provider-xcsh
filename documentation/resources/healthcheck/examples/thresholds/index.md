---
page_title: "Thresholds"
subcategory: "Monitoring"
description: "Thresholds for xcsh_healthcheck."
xcsh_docs: {"aliases": ["thresholds"], "body_bytes": 1102, "body_sha256": "sha256:ff2977543b596cd01a8dfab5e32938fe1ade1fc9579c7a0371336afc4eeea638", "capabilities": ["load-balancing"], "category": "load-balancing", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:healthcheck:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:eb7ffc7c4c13b7cf863658813f0b685ce60bc365c197904650a919fd4f6c683f", "source_path": "examples/resources/xcsh_healthcheck/thresholds.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:healthcheck:example:thresholds", "parent_id": "xcsh-docs:resources:healthcheck:examples", "path": "documentation/resources/healthcheck/examples/thresholds/index.md", "product": "distributed-cloud", "provider_name": "healthcheck", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-1022121031301321-3032210032102310-2003203223112330-3320322300010310-0101222300012320-0231110320011312-3231113132211233-1102111121110131", "registry_path": "docs/guides/resources--healthcheck--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["thresholds"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/healthcheck/examples/thresholds/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Thresholds for xcsh_healthcheck.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["healthcheckCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Thresholds

Breadcrumbs:

- [xcsh_healthcheck](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/healthcheck/examples/)
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
