---
page_title: "Disable anonymization"
subcategory: "Security"
description: "Disable anonymization for xcsh_app_firewall."
xcsh_docs: {"aliases": ["disable-anonymization"], "body_bytes": 1222, "body_sha256": "sha256:2b1b5e1cc59fe654c11f84ffc7e5351d4be8e1970f0c98172fee7a08aa142387", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:4e7ac5ea3e4dfe4980ba8b5bf719912588c340253be3f414343af180a9e177fc", "source_path": "examples/resources/xcsh_app_firewall/disable-anonymization.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:disable-anonymization", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "documentation/resources/app_firewall/examples/disable-anonymization/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-2110210122122203-3323213231132120-1111132123303212-2213203320203301-3002311131123200-2022321322101310-2123122213001130-0231121221330132", "registry_path": "docs/guides/resources--app_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["disable-anonymization"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/disable-anonymization/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Disable anonymization for xcsh_app_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["app_firewallCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Disable anonymization

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/examples/)
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
