---
page_title: "Tls fingerprint"
subcategory: ""
description: "Tls fingerprint for xcsh_user_identification."
xcsh_docs: {"aliases": ["tls-fingerprint"], "body_bytes": 1315, "body_sha256": "sha256:a13f68fbd29a9fc64d5dfb17def5f5bca5dccb00788db57a7916c86af595b8d9", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:18aace91bd93f5201f1f6ed23d33d48d6c8f66a90c3a44614cea4e6f6f5a4a09", "source_path": "examples/resources/xcsh_user_identification/tls-fingerprint.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:user_identification:example:tls-fingerprint", "parent_id": "xcsh-docs:resources:user_identification:examples", "path": "documentation/resources/user_identification/examples/tls-fingerprint/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0331220123100321-2232002120000133-0000022303231133-3302321233120223-2331211231102231-1232333102133030-3300300322001211-2131300102011010", "registry_path": "docs/guides/resources--user_identification--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["tls-fingerprint"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/examples/tls-fingerprint/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Tls fingerprint for xcsh_user_identification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["user_identificationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/examples/)
- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
