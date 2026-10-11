---
page_title: "Allowed response codes"
subcategory: "Security"
description: "Allowed response codes for xcsh_app_firewall."
xcsh_docs: {"aliases": ["allowed-response-codes"], "body_bytes": 1268, "body_sha256": "sha256:cfc9da9a99435eec927e909bc1fb61e16d2f6e0f18bba07a04960564940b0159", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:9fac3f5328e4e5b2882046677e97a6cbe0fb6c753fef26f8f52956cc720b9fe4", "source_path": "examples/resources/xcsh_app_firewall/allowed-response-codes.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:allowed-response-codes", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "documentation/resources/app_firewall/examples/allowed-response-codes/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-0203213320003220-1301032123212000-1332232001222030-1103103320122132-3122323103311322-3303032233323020-0212323202330310-3020113212303031", "registry_path": "docs/guides/resources--app_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["allowed-response-codes"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/allowed-response-codes/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Allowed response codes for xcsh_app_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["app_firewallCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Allowed response codes

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/examples/)
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
