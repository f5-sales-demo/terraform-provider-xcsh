---
page_title: "Tls fingerprint"
subcategory: ""
description: "Tls fingerprint for xcsh_user_identification."
xcsh_docs: {"aliases": ["tls-fingerprint"], "body_bytes": 1072, "body_sha256": "sha256:3fa914f3069c2f6ff4f35cf209e07c626ef53b50107f23b5a5a0a27c94503a43", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:18aace91bd93f5201f1f6ed23d33d48d6c8f66a90c3a44614cea4e6f6f5a4a09", "source_path": "examples/resources/xcsh_user_identification/tls-fingerprint.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:user_identification:example:tls-fingerprint", "parent_id": "xcsh-docs:resources:user_identification:examples", "path": "documentation/resources/user_identification/examples/tls-fingerprint/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:c6e4ccf56e8887c4192e662ded7285f98ad804e3254075567df8f73a44375945", "provider_type": "resources", "registry_anchor": "canonical-0331220123100321-2232002120000133-0000022303231133-3302321233120223-2331211231102231-1232333102133030-3300300322001211-2131300102011010", "registry_path": "docs/guides/resources--user_identification--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["tls-fingerprint"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/examples/tls-fingerprint/index.txt", "spec_pin_digest": "sha256:ba30301dbe17345dfb4303bb66013effc7812eafc55ed39ca25a02278fc1ac8d", "summary": "Tls fingerprint for xcsh_user_identification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.5", "schema_components": ["user_identificationCreateRequest"], "target_commit": "d0d1579f49a5777ad35f6cd777a0f12db4b2442f"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Tls fingerprint

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/examples/)
- Tls fingerprint

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/tls-fingerprint.tf`; digest `sha256:18aace91bd93f5201f1f6ed23d33d48d6c8f66a90c3a44614cea4e6f6f5a4a09`.

```terraform
# TlsFingerprint — Acceptance-test-derived Configuration
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

resource "xcsh_user_identification" "test" {
  name      = "example"
  namespace = "system"

  rules {
    tls_fingerprint = {}
  }
}
```
