---
page_title: "Bot protection"
subcategory: "Security"
description: "Bot protection for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1528, "body_sha256": "sha256:a04ef2fdd6a6347415b9d061932ebb38438c139c0fcf3ae377b7f8eb2785ec60", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:6e9c8727741ef612628e60f244a11489e11b01d40ff73ee72db9aa6ea5f61164", "source_path": "examples/resources/xcsh_app_firewall/bot-protection.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:bot-protection", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "documentation/resources/app_firewall/examples/bot-protection/index.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["bot-protection"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/bot-protection/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Bot protection for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Bot protection

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/examples/)
- Bot protection

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/bot-protection.tf`; digest `sha256:6e9c8727741ef612628e60f244a11489e11b01d40ff73ee72db9aa6ea5f61164`.

```terraform
# BotProtection — Acceptance-test-derived Configuration
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

  default_detection_settings = {}
  allow_all_response_codes   = {}
  blocking                   = {}
  use_default_blocking_page  = {}
  default_anonymization      = {}

  bot_protection_setting {
    good_bot_action       = "REPORT"
    malicious_bot_action  = "BLOCK"
    suspicious_bot_action = "REPORT"
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/examples/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
