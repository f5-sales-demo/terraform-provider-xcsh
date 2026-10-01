---
page_title: "With updated labels"
subcategory: "Security"
description: "With updated labels for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1562, "body_sha256": "sha256:bdd696ec2d2a7ca02776968161db1c9904036671d9700253c108cbf93546470e", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:0bbaf48e9cd535c6fd630956b0756feb6f88090c3cb140fc62c351e03346e5cd", "source_path": "examples/resources/xcsh_app_firewall/with-updated-labels.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:with-updated-labels", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "documentation/resources/app_firewall/examples/with-updated-labels/index.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["with-updated-labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/with-updated-labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With updated labels for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/examples/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
