---
page_title: "With rules"
subcategory: ""
description: "With rules for xcsh_user_identification."
xcsh_docs: {"aliases": [], "body_bytes": 1151, "body_sha256": "sha256:f9bf4c5523fd957994e9eb86da67cdd1403c1b28ef64f807a9fc5718b6123da0", "canonical_id": "xcsh-docs:resources:user_identification:example:with-rules", "child_ids": [], "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:b71ec64258f8256d380d65c27c4c945068a8b584275196cad8f9d2b7fcbb430f", "source_path": "examples/resources/xcsh_user_identification/with-rules.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:user_identification:example:with-rules", "parent_id": "xcsh-docs:resources:user_identification:examples", "path": "docs/guides/resources--user_identification--example--with-rules.md", "provider_name": "user_identification", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["with-rules"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/examples/with-rules/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With rules for xcsh_user_identification.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["user_identificationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With rules

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md)
- [Examples](resources--user_identification--examples.md)
- With rules

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/with-rules.tf`; digest `sha256:b71ec64258f8256d380d65c27c4c945068a8b584275196cad8f9d2b7fcbb430f`.

```terraform
# WithRules — Acceptance-test-derived Configuration
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
  name        = "example"
  namespace   = "system"
  description = "User identification with identification rules"

  rules {
    client_ip = {}
  }
}
```

## Next pages

- [Examples](resources--user_identification--examples.md)
- [xcsh_user_identification](../resources/user_identification.md)
