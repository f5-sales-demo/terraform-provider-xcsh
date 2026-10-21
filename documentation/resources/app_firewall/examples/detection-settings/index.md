---
page_title: "Detection settings"
subcategory: "Security"
description: "Detection settings for xcsh_app_firewall."
xcsh_docs: {"aliases": ["detection-settings"], "body_bytes": 1462, "body_sha256": "sha256:2b74b104051a4283fba9ee5ae6b9635b3169adf990e82835969e6d2eb360a076", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:8dfd23c184c15d1235eb2c949703ae26669c8f88a0eda5530bedf5a3bd8cfce2", "source_path": "examples/resources/xcsh_app_firewall/detection-settings.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:detection-settings", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "documentation/resources/app_firewall/examples/detection-settings/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:5da5abed63484d73ce77467c8c65a96ecdf7e04c3db5d5bd821e0fa97991e015", "provider_type": "resources", "registry_anchor": "canonical-2021221133300300-1122033021031110-1100210202002123-3032001223103320-3110131122302213-3313032232122210-2333311101020021-3133113010310323", "registry_path": "docs/guides/resources--app_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["detection-settings"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/detection-settings/index.txt", "spec_pin_digest": "sha256:d30f9f9a1cd22b9ba948a7409b4aa6907eda743fb73415c72474bd2784e91af1", "summary": "Detection settings for xcsh_app_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.1.5", "schema_components": ["app_firewallCreateRequest"], "target_commit": "85874b2a5cfba24c95e5f7d52218df326d4a0582"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Detection settings

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/examples/)
- Detection settings

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/detection-settings.tf`; digest `sha256:8dfd23c184c15d1235eb2c949703ae26669c8f88a0eda5530bedf5a3bd8cfce2`.

```terraform
# DetectionSettings — Acceptance-test-derived Configuration
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
  name      = "example"
  namespace = "system"

  blocking = {}

  detection_settings {
    default_violation_settings = {}

    enable_suppression      = {}
    enable_threat_campaigns = {}
    signature_selection_setting {
      high_medium_accuracy_signatures = {}
      signature_settings_by_accuracy {
        high_accuracy_action   = "SIG_BLOCK"
        medium_accuracy_action = "SIG_BLOCK"
        low_accuracy_action    = "SIG_BLOCK"
      }
      default_attack_type_settings = {}
    }
  }
}
```
