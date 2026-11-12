---
page_title: "Blocking"
subcategory: "Security"
description: "Blocking for xcsh_app_firewall."
xcsh_docs: {"aliases": ["blocking"], "body_bytes": 1462, "body_sha256": "sha256:5db0ee883ee795b1fdc0232587ae4439861dd65007ffaa37c191798fcc4a5fd8", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:ef7e8d73d30409b783a8406919290ea0de139474cbcea62a1582fcfe0f55c136", "source_path": "examples/resources/xcsh_app_firewall/blocking.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:blocking", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "documentation/resources/app_firewall/examples/blocking/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e62318985676da5dc8f5d2d35ccc26c58c5f11f463e5627995ce95f8658f6b43", "provider_type": "resources", "registry_anchor": "canonical-0332302023310211-0300010022330303-3303322211000332-2011101230233310-0100123200232311-2303111230001210-2231313003213022-2203212330020023", "registry_path": "docs/guides/resources--app_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["blocking"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/blocking/index.txt", "spec_pin_digest": "sha256:b750ce4b9c0ad9c90dad992a893773bacfb3aeff7598bc37194f3037f353c949", "summary": "Blocking for xcsh_app_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.5.2", "schema_components": ["app_firewallCreateRequest"], "target_commit": "18cc909b5ddc97ef9abd48764589d753290a31a4"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Blocking

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/examples/)
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
