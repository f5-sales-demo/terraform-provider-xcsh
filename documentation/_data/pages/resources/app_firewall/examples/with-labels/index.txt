---
page_title: "With labels"
subcategory: "Security"
description: "With labels for xcsh_app_firewall."
xcsh_docs: {"aliases": ["with-labels"], "body_bytes": 1449, "body_sha256": "sha256:75fca453babbaf6af48c00130b409a4664f8593b22ff67eb3f9370e8ddaff435", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:51f4bc6ac4257ec69ec14e7e0c37f3b9cd582e2461dff7632a889f943d97b42f", "source_path": "examples/resources/xcsh_app_firewall/with-labels.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:with-labels", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "documentation/resources/app_firewall/examples/with-labels/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:5a7fb41daf7683904c87458d3d7c40e4f3e095bd9d9aff0ef4c2c67cb6c9a8b5", "provider_type": "resources", "registry_anchor": "canonical-2030311032102122-1000300211002213-0320313220310320-3003132331232001-0300220121323123-1113011310213301-2020320301333203-0321113100013212", "registry_path": "docs/guides/resources--app_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-labels"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/with-labels/index.txt", "spec_pin_digest": "sha256:fb3399d426b86fc806bdce295180d1446d1c05db41b1ae48575b9b6c2409bc5e", "summary": "With labels for xcsh_app_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.1", "schema_components": ["app_firewallCreateRequest"], "target_commit": "af922688a0dab75542a6bd0181ddd80fee4c8c2c"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With labels

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/examples/)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/with-labels.tf`; digest `sha256:51f4bc6ac4257ec69ec14e7e0c37f3b9cd582e2461dff7632a889f943d97b42f`.

```terraform
# WithLabels — Acceptance-test-derived Configuration
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
    environment = "test"
    team        = "security"
  }

  # Use default detection settings
  default_detection_settings = {}

  # Allow all response codes
  allow_all_response_codes = {}

  # Blocking mode
  blocking = {}

  # Use default blocking page
  use_default_blocking_page = {}

  # Use default bot settings
  default_bot_setting = {}

  # Use default anonymization
  default_anonymization = {}
}
```
