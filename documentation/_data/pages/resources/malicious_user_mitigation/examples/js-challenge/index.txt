---
page_title: "Js challenge"
subcategory: ""
description: "Js challenge for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": ["js-challenge"], "body_bytes": 1471, "body_sha256": "sha256:919438865fb74546f3c751c9ea724bbc64c4ffe7ff5879e07d009e4409896079", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:31cce82837db63f5b4c06f4195de3f808d4789b7437b6ae8148fa42b9b7c4194", "source_path": "examples/resources/xcsh_malicious_user_mitigation/js-challenge.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:malicious_user_mitigation:example:js-challenge", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:examples", "path": "documentation/resources/malicious_user_mitigation/examples/js-challenge/index.md", "product": "distributed-cloud", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-3213332111110303-0233120032031331-3130132222023022-3233112232230001-1110123013010013-2020020120033103-1130312223101333-3203101321203012", "registry_path": "docs/guides/resources--malicious_user_mitigation--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["js-challenge"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/examples/js-challenge/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Js challenge for xcsh_malicious_user_mitigation.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Js challenge

Breadcrumbs:

- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/examples/)
- Js challenge

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/js-challenge.tf`; digest `sha256:31cce82837db63f5b4c06f4195de3f808d4789b7437b6ae8148fa42b9b7c4194`.

```terraform
# JsChallenge — Acceptance-test-derived Configuration
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

resource "xcsh_malicious_user_mitigation" "test" {
  name      = "example"
  namespace = "system"

  mitigation_type {
    rules {
      threat_level {
        medium = {}
      }
      mitigation_action {
        javascript_challenge = {}
      }
    }
  }
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/examples/)
- [xcsh_malicious_user_mitigation](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/malicious_user_mitigation/)
