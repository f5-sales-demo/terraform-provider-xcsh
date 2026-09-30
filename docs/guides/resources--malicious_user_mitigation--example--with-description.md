---
page_title: "With description"
subcategory: ""
description: "With description for xcsh_malicious_user_mitigation."
xcsh_docs: {"aliases": [], "body_bytes": 1385, "body_sha256": "sha256:f853fa786d031a396564bae85dc7a2a593a7869da386b8bdf676111e14a19adb", "canonical_id": "xcsh-docs:resources:malicious_user_mitigation:example:with-description", "child_ids": [], "collection_id": "xcsh-docs:resources:malicious_user_mitigation:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:a0561f064cef4a0bf4898a53eed03d3f77741d1f6b9b0474fd171ea08dcd9202", "source_path": "examples/resources/xcsh_malicious_user_mitigation/with-description.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:malicious_user_mitigation:example:with-description", "parent_id": "xcsh-docs:resources:malicious_user_mitigation:examples", "path": "docs/guides/resources--malicious_user_mitigation--example--with-description.md", "provider_name": "malicious_user_mitigation", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["with-description"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/malicious_user_mitigation/examples/with-description/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "With description for xcsh_malicious_user_mitigation.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["malicious_user_mitigationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# With description

Breadcrumbs:

- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md)
- [Examples](resources--malicious_user_mitigation--examples.md)
- With description

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_malicious_user_mitigation/with-description.tf`; digest `sha256:a0561f064cef4a0bf4898a53eed03d3f77741d1f6b9b0474fd171ea08dcd9202`.

```terraform
# WithDescription — Acceptance-test-derived Configuration
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
  depends_on  = [time_sleep.wait_for_namespace]
  name        = "example-value"
  namespace   = xcsh_namespace.test.name
  description = "example-description"
}
```

## Next pages

- [Examples](resources--malicious_user_mitigation--examples.md)
- [xcsh_malicious_user_mitigation](../resources/malicious_user_mitigation.md)
