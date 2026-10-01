---
page_title: "With annotations"
subcategory: ""
description: "With annotations for xcsh_user_identification."
xcsh_docs: {"aliases": [], "body_bytes": 1164, "body_sha256": "sha256:b6655368c107cf9944f11ac79dccaa03a79453548f449d759ababb22bd19c3ee", "canonical_id": "xcsh-docs:resources:user_identification:example:with-annotations", "child_ids": [], "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:903200cccd3aa8c5abe2eebfb30b2e2efe55bae4a37636c3dae4de46ed18a25f", "source_path": "examples/resources/xcsh_user_identification/with-annotations.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:user_identification:example:with-annotations", "parent_id": "xcsh-docs:resources:user_identification:examples", "path": "docs/guides/resources--user_identification--example--with-annotations.md", "provider_name": "user_identification", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["with-annotations"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/examples/with-annotations/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With annotations for xcsh_user_identification.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["user_identificationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With annotations

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md)
- [Examples](resources--user_identification--examples.md)
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

## Next pages

- [Examples](resources--user_identification--examples.md)
- [xcsh_user_identification](../resources/user_identification.md)
