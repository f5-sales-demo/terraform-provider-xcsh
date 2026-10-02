---
page_title: "Blocking"
subcategory: "Security"
description: "Blocking for xcsh_app_firewall."
xcsh_docs: {"aliases": ["blocking"], "body_bytes": 1684, "body_sha256": "sha256:4a690de4cc580b52581c56715898c71976615f89924c3ed5288e26b90920f1ff", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:3217787f954272d9dc634bbb1180c6a67657552249a3e4198b15d116afcf5824", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:ef7e8d73d30409b783a8406919290ea0de139474cbcea62a1582fcfe0f55c136", "source_path": "examples/resources/xcsh_app_firewall/blocking.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:blocking", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "documentation/resources/app_firewall/examples/blocking/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "resources", "registry_anchor": "canonical-0332302023310211-0300010022330303-3303322211000332-2011101230233310-0100123200232311-2303111230001210-2231313003213022-2203212330020023", "registry_path": "docs/guides/resources--app_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["blocking"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/blocking/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Blocking for xcsh_app_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["app_firewallCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/examples/)
- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
