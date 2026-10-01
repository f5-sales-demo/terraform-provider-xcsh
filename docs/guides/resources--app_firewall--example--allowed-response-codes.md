---
page_title: "Allowed response codes"
subcategory: "Security"
description: "Allowed response codes for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1284, "body_sha256": "sha256:ab0d9b36ae47b1a541d4531464d9bd522866f032e84a081a59cf6ff1c19a1653", "canonical_id": "xcsh-docs:resources:app_firewall:example:allowed-response-codes", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:9fac3f5328e4e5b2882046677e97a6cbe0fb6c753fef26f8f52956cc720b9fe4", "source_path": "examples/resources/xcsh_app_firewall/allowed-response-codes.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:allowed-response-codes", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "docs/guides/resources--app_firewall--example--allowed-response-codes.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["allowed-response-codes"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/allowed-response-codes/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Allowed response codes for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Allowed response codes

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Examples](resources--app_firewall--examples.md)
- Allowed response codes

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/allowed-response-codes.tf`; digest `sha256:9fac3f5328e4e5b2882046677e97a6cbe0fb6c753fef26f8f52956cc720b9fe4`.

```terraform
# AllowedResponseCodes — Acceptance-test-derived Configuration
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
  blocking                   = {}
  use_default_blocking_page  = {}
  default_bot_setting        = {}
  default_anonymization      = {}

  allowed_response_codes {
    response_code = [200, 204, 301, 302]
  }
}
```

## Next pages

- [Examples](resources--app_firewall--examples.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
