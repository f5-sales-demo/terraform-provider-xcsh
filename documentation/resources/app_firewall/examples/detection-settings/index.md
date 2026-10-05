---
page_title: "Detection settings"
subcategory: "Security"
description: "Detection settings for xcsh_app_firewall."
xcsh_docs: {"aliases": ["detection-settings"], "body_bytes": 1695, "body_sha256": "sha256:cfbb2388d9f047b5c25bd4d7ce5e5f27b3fe48afff7e5254952f3181bccf0a95", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:50e5c6c8e8107cae5d8fe6c732ac501ab51251e38c9d54deab9761fd88fcb6f2", "source_path": "examples/resources/xcsh_app_firewall/detection-settings.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:detection-settings", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "documentation/resources/app_firewall/examples/detection-settings/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-2021221133300300-1122033021031110-1100210202002123-3032001223103320-3110131122302213-3313032232122210-2333311101020021-3133113010310323", "registry_path": "docs/guides/resources--app_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["detection-settings"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/detection-settings/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Detection settings for xcsh_app_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
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

Source: `examples/resources/xcsh_app_firewall/detection-settings.tf`; digest `sha256:50e5c6c8e8107cae5d8fe6c732ac501ab51251e38c9d54deab9761fd88fcb6f2`.

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

  allow_all_response_codes  = {}
  blocking                  = {}
  use_default_blocking_page = {}
  default_bot_setting       = {}
  default_anonymization     = {}

  detection_settings {
    default_violation_settings = {}
    default_bot_setting        = {}
    enable_suppression         = {}
    enable_threat_campaigns    = {}
    signature_selection_setting {
      high_medium_accuracy_signatures = {}
      default_attack_type_settings    = {}
    }
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/examples/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
