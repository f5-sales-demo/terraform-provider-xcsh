---
page_title: "Ai enhancements"
subcategory: "Security"
description: "Ai enhancements for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1186, "body_sha256": "sha256:8119fbe18d0ade9c9968a1a05baebce0ee00659345bb063f04806940bf3fd782", "canonical_id": "xcsh-docs:resources:app_firewall:example:ai-enhancements", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:8a5b5775deccd3da5f7c4e690d3bca03ede520f6e7ee1e00db0fa32e73bad912", "source_path": "examples/resources/xcsh_app_firewall/ai-enhancements.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:ai-enhancements", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "docs/guides/resources--app_firewall--example--ai-enhancements.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["ai-enhancements"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/ai-enhancements/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Ai enhancements for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Ai enhancements

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Examples](resources--app_firewall--examples.md)
- Ai enhancements

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/ai-enhancements.tf`; digest `sha256:8a5b5775deccd3da5f7c4e690d3bca03ede520f6e7ee1e00db0fa32e73bad912`.

```terraform
# AiEnhancements — Acceptance-test-derived Configuration
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
  default_bot_setting        = {}
  default_anonymization      = {}

  enable_ai_enhancements {
    mitigate_high_risk_action = {}
  }
}
```

## Next pages

- [Examples](resources--app_firewall--examples.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
