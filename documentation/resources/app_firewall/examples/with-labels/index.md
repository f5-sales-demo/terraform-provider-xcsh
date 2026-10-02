---
page_title: "With labels"
subcategory: "Security"
description: "With labels for xcsh_app_firewall."
xcsh_docs: {"aliases": ["with-labels"], "body_bytes": 1671, "body_sha256": "sha256:0ea64d620ba3ad9417227f57ce923b2eca3ce19182a6c5e97e77d7f55e6a914e", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:51f4bc6ac4257ec69ec14e7e0c37f3b9cd582e2461dff7632a889f943d97b42f", "source_path": "examples/resources/xcsh_app_firewall/with-labels.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:with-labels", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "documentation/resources/app_firewall/examples/with-labels/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-2030311032102122-1000300211002213-0320313220310320-3003132331232001-0300220121323123-1113011310213301-2020320301333203-0321113100013212", "registry_path": "docs/guides/resources--app_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-labels"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/with-labels/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "With labels for xcsh_app_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/examples/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
