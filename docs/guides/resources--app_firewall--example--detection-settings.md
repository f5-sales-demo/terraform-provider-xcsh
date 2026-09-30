---
page_title: "Detection settings"
subcategory: "Security"
description: "Detection settings for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1390, "body_sha256": "sha256:934c840266a988402d7b0e78173cd8e619b8bacc0e8588ff85032170d1605a84", "canonical_id": "xcsh-docs:resources:app_firewall:example:detection-settings", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:50e5c6c8e8107cae5d8fe6c732ac501ab51251e38c9d54deab9761fd88fcb6f2", "source_path": "examples/resources/xcsh_app_firewall/detection-settings.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:detection-settings", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "docs/guides/resources--app_firewall--example--detection-settings.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["detection-settings"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/detection-settings/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Detection settings for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Detection settings

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Examples](resources--app_firewall--examples.md)
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

- [Examples](resources--app_firewall--examples.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
