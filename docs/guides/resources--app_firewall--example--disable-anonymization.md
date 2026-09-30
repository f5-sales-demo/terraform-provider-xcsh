---
page_title: "Disable anonymization"
subcategory: "Security"
description: "Disable anonymization for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1139, "body_sha256": "sha256:669e27cbf5d22f2d6c254de5278d6041e4840d3b05841c564faf7eee76068e27", "canonical_id": "xcsh-docs:resources:app_firewall:example:disable-anonymization", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:4e7ac5ea3e4dfe4980ba8b5bf719912588c340253be3f414343af180a9e177fc", "source_path": "examples/resources/xcsh_app_firewall/disable-anonymization.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:disable-anonymization", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "docs/guides/resources--app_firewall--example--disable-anonymization.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["disable-anonymization"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/disable-anonymization/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Disable anonymization for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Disable anonymization

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Examples](resources--app_firewall--examples.md)
- Disable anonymization

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/disable-anonymization.tf`; digest `sha256:4e7ac5ea3e4dfe4980ba8b5bf719912588c340253be3f414343af180a9e177fc`.

```terraform
# DisableAnonymization — Acceptance-test-derived Configuration
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

  disable_anonymization = {}
}
```

## Next pages

- [Examples](resources--app_firewall--examples.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
