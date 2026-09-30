---
page_title: "With labels"
subcategory: ""
description: "With labels for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": [], "body_bytes": 1376, "body_sha256": "sha256:bba1943441d1120b0328858e58ff732106949d25f759c3f6b6528c7efab35bc9", "canonical_id": "xcsh-docs:resources:malicious_user_mitigation:example:with-labels", "child_ids": [], "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:05d2d99ea2c9b4398348e89d32380bd3c6c3cf7e2154f51c8b74dcc85481aec2", "source_path": "examples/resources/xcsh_malicious_user_mitigation/with-labels.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:malicious_user_mitigation:example:with-labels", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:examples", "path": "docs/guides/resources--malicious_user_mitigation--example--with-labels.md", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["with-labels"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/examples/with-labels/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With labels for xcsh_malicious_user_mitigation.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# With labels

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md)
- [Examples](resources--malicious_user_mitigation--examples.md)
- With labels

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/with-labels.tf`; digest `sha256:05d2d99ea2c9b4398348e89d32380bd3c6c3cf7e2154f51c8b74dcc85481aec2`.

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
    time = {
      source  = "hashicorp/time"
      version = "= 0.13.1"
    }
  }
}

resource "xcsh_namespace" "test" {
  name = "example"
}

resource "time_sleep" "wait_for_namespace" {
  depends_on      = [xcsh_namespace.test]
  create_duration = "5s"
}

resource "xcsh_malicious_user_mitigation" "test" {
  depends_on = [time_sleep.wait_for_namespace]
  name       = "example-value"
  namespace  = xcsh_namespace.test.name

  labels = {
    example-key = "example-value"
  }
}
```

## Next pages

- [Examples](resources--malicious_user_mitigation--examples.md)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md)
