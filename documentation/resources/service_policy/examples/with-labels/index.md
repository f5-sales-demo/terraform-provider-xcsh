---
page_title: "With labels"
subcategory: "Security"
description: "With labels for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1327, "body_sha256": "sha256:02b3e376f14878f474021a9f87b8a8cd46061ebacb61464fcf1d8f97619ebd43", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:27966c953ab74ec12398f4c688232398a984b67f1bec6a9179d66139d051b625", "source_path": "examples/resources/xcsh_service_policy/with-labels.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:service_policy:example:with-labels", "parent_id": "xcsh-docs:resources:service_policy:examples", "path": "documentation/resources/service_policy/examples/with-labels/index.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["with-labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/examples/with-labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With labels for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# With labels

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/examples/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
