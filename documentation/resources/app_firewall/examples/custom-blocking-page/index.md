---
page_title: "Custom blocking page"
subcategory: "Security"
description: "Custom blocking page for xcsh_app_firewall."
xcsh_docs: {"aliases": ["custom-blocking-page"], "body_bytes": 1297, "body_sha256": "sha256:95b4f6d4ec1e979c70aa675049deb10b605dab93330e2a21c6a40de5aba2188b", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:app_firewall:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:f2df3fea5ad8e1732d977424427aaf5de894118f0229d9185c0686db929c5462", "source_path": "examples/resources/xcsh_app_firewall/custom-blocking-page.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:app_firewall:example:custom-blocking-page", "parent_id": "xcsh-docs:resources:app_firewall:examples", "path": "documentation/resources/app_firewall/examples/custom-blocking-page/index.md", "product": "distributed-cloud", "provider_name": "app_firewall", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "resources", "registry_anchor": "canonical-3030200331103100-3101110011212013-0333230030101303-0201233212122322-0102010113030120-1031331022311321-3312212132122333-3303112223121131", "registry_path": "docs/guides/resources--app_firewall--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["custom-blocking-page"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_firewall/examples/custom-blocking-page/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Custom blocking page for xcsh_app_firewall.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["app_firewallCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Custom blocking page

Breadcrumbs:

- [xcsh_app_firewall](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_firewall/examples/)
- Custom blocking page

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_app_firewall/custom-blocking-page.tf`; digest `sha256:f2df3fea5ad8e1732d977424427aaf5de894118f0229d9185c0686db929c5462`.

```terraform
# CustomBlockingPage — Acceptance-test-derived Configuration
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
  default_bot_setting        = {}
  default_anonymization      = {}

  blocking_page {
    blocking_page = "https://example.com/blocked.html"
    response_code = "Forbidden"
  }
}
```
