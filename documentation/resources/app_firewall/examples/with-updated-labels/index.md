---
page_title: "With updated labels"
subcategory: "Security"
description: "With updated labels for xcsh_app_firewall."
xcsh_docs: {"aliases": ["with-updated-labels"], "body_bytes": 1340, "body_sha256": "sha256:773a7fcc57d3613d5d3415d457e4d7d1595f3fedff26778715231397cebf5de8", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:0bbaf48e9cd535c6fd630956b0756feb6f88090c3cb140fc62c351e03346e5cd", "source_path": "examples/resources/xcsh_app_firewall/with-updated-labels.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:with-updated-labels", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "documentation/resources/app_firewall/examples/with-updated-labels/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:7e724befbd28dae1d544e2cc8bbc72d0e62fdd0382374fb6e98044bf0ff0e829", "provider_type": "resources", "registry_anchor": "canonical-3220033311201323-0220231123202100-3320212210212113-3220330013122010-1032333323302330-0330221322212032-2110300111313030-1023220322032102", "registry_path": "docs/guides/resources--app_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-updated-labels"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/with-updated-labels/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "With updated labels for xcsh_app_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["app_firewallCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With updated labels

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/examples/)
- With updated labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/with-updated-labels.tf`; digest `sha256:0bbaf48e9cd535c6fd630956b0756feb6f88090c3cb140fc62c351e03346e5cd`.

```terraform
# WithUpdatedLabels — Acceptance-test-derived Configuration
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

resource "xcsh_app_firewall" "test" {
  name        = "example"
  namespace   = "system"
  description = "Test application firewall"

  labels = {
    environment = "staging"
    team        = "platform"
  }

  default_detection_settings = {}
  allow_all_response_codes   = {}
  blocking                   = {}
  use_default_blocking_page  = {}
  default_bot_setting        = {}
  default_anonymization      = {}
}
```
