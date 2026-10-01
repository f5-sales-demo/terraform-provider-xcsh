---
page_title: "With labels"
subcategory: "Security"
description: "With labels for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1220, "body_sha256": "sha256:fd5d9f82b0d9d99fb9fa44c2ed843bd5c4a463d4baca1b0a43f45bc4efd05c1e", "canonical_id": "xcsh-docs:resources:service_policy:example:with-labels", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:27966c953ab74ec12398f4c688232398a984b67f1bec6a9179d66139d051b625", "source_path": "examples/resources/xcsh_service_policy/with-labels.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:service_policy:example:with-labels", "parent_id": "xcsh-docs:resources:service_policy:examples", "path": "docs/guides/resources--service_policy--example--with-labels.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["with-labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/examples/with-labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With labels for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# With labels

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Examples](resources--service_policy--examples.md)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/with-labels.tf`; digest `sha256:27966c953ab74ec12398f4c688232398a984b67f1bec6a9179d66139d051b625`.

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

resource "xcsh_service_policy" "test" {
  name        = "example"
  namespace   = "system"
  description = "Test service policy"

  labels = {
    environment = "test"
    team        = "security"
  }

  # Allow all requests
  allow_all_requests = {}

  # Apply to any server
  any_server = {}
}
```

## Next pages

- [Examples](resources--service_policy--examples.md)
- [xcsh_service_policy](../resources/service_policy.md)
