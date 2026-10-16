---
page_title: "With annotations"
subcategory: ""
description: "With annotations for xcsh_user_identification."
xcsh_docs: {"aliases": ["with-annotations"], "body_bytes": 1127, "body_sha256": "sha256:77385ef41150aa2c5a2eaefe3b823d5521e8dadf69050a39ec318c301d07e7a1", "capabilities": ["security"], "category": "security", "child_ids": [], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:903200cccd3aa8c5abe2eebfb30b2e2efe55bae4a37636c3dae4de46ed18a25f", "source_path": "examples/resources/xcsh_user_identification/with-annotations.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:user_identification:example:with-annotations", "parent_id": "xcsh-docs:resources:user_identification:examples", "path": "documentation/resources/user_identification/examples/with-annotations/index.md", "product": "distributed-cloud", "provider_name": "user_identification", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "resources", "registry_anchor": "canonical-3311100230323221-0011100032132010-0321103010011311-2010100032303113-0111130122101232-3121030303201011-2223132312110233-0211103212022323", "registry_path": "docs/guides/resources--user_identification--examples--group-001.md", "relationships": [], "retrieval_version": 1, "role": "example", "schema_path": ["with-annotations"], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/examples/with-annotations/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "With annotations for xcsh_user_identification.", "tasks": ["configuration"], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["user_identificationCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With annotations

Breadcrumbs:

- [xcsh_user_identification](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/user_identification/examples/)
- With annotations

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/with-annotations.tf`; digest `sha256:903200cccd3aa8c5abe2eebfb30b2e2efe55bae4a37636c3dae4de46ed18a25f`.

```terraform
# WithAnnotations — Acceptance-test-derived Configuration
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

  annotations = {
    example-key = "example-value"
  }

  rules {
    client_ip = {}
  }
}
```
