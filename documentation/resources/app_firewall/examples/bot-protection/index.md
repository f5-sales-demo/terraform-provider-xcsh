---
page_title: "Bot protection"
subcategory: "Security"
description: "Bot protection for xcsh_app_firewall."
xcsh_docs: {"aliases": ["bot-protection"], "body_bytes": 1528, "body_sha256": "sha256:a04ef2fdd6a6347415b9d061932ebb38438c139c0fcf3ae377b7f8eb2785ec60", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:6e9c8727741ef612628e60f244a11489e11b01d40ff73ee72db9aa6ea5f61164", "source_path": "examples/resources/xcsh_app_firewall/bot-protection.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:bot-protection", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "documentation/resources/app_firewall/examples/bot-protection/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "resources", "registry_anchor": "canonical-0211211133000211-3001323131330100-2101220020033303-3332001321322033-1322203210020101-3323033020022330-2003201020112222-3321210302020120", "registry_path": "docs/guides/resources--app_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["bot-protection"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/bot-protection/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Bot protection for xcsh_app_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
