---
page_title: "Blocking"
subcategory: "Security"
description: "Blocking for xcsh_app_firewall."
xcsh_docs: {"aliases": [], "body_bytes": 1379, "body_sha256": "sha256:400b9768c2ab77c4e9d087f227358f9aaf937e0d47fe96ebd2127fe14c51c08b", "canonical_id": "xcsh-docs:resources:app_firewall:example:blocking", "child_ids": [], "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:ef7e8d73d30409b783a8406919290ea0de139474cbcea62a1582fcfe0f55c136", "source_path": "examples/resources/xcsh_app_firewall/blocking.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:blocking", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "docs/guides/resources--app_firewall--example--blocking.md", "provider_name": "app_firewall", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["blocking"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/blocking/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Blocking for xcsh_app_firewall.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_firewallCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Blocking

Breadcrumbs:

- [xcsh_app_firewall](../resources/app_firewall.md)
- [Examples](resources--app_firewall--examples.md)
- Blocking

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/blocking.tf`; digest `sha256:ef7e8d73d30409b783a8406919290ea0de139474cbcea62a1582fcfe0f55c136`.

```terraform
# Blocking — Acceptance-test-derived Configuration
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

  # Use default detection settings
  default_detection_settings = {}

  # Blocking mode - actively block malicious requests
  blocking = {}

  # allow_all_response_codes / use_default_blocking_page / default_bot_setting /
  # default_anonymization are server-default oneof markers the provider import-suppresses.
  # Declaring them makes the config import-unclean (config has them, imported state does
  # not), so they are intentionally omitted here — the server still materializes them.
}
```

## Next pages

- [Examples](resources--app_firewall--examples.md)
- [xcsh_app_firewall](../resources/app_firewall.md)
