---
page_title: "Deny list"
subcategory: "Security"
description: "Deny list for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1132, "body_sha256": "sha256:4657580d153810506c63440e12676633c9e32603c92d32fd4f53291021c4e1bf", "canonical_id": "xcsh-docs:resources:service_policy:example:deny-list", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:6000e015862d1bb25f8e6b5bc94664685fdab920f11f2608c4f47416c65ff393", "source_path": "examples/resources/xcsh_service_policy/deny-list.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:service_policy:example:deny-list", "parent_id": "xcsh-docs:resources:service_policy:examples", "path": "docs/guides/resources--service_policy--example--deny-list.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["deny-list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/examples/deny-list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Deny list for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Deny list

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Examples](resources--service_policy--examples.md)
- Deny list

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/deny-list.tf`; digest `sha256:6000e015862d1bb25f8e6b5bc94664685fdab920f11f2608c4f47416c65ff393`.

```terraform
# DenyList — Acceptance-test-derived Configuration
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
  name      = "example"
  namespace = "system"

  deny_list {
    prefix_list {
      prefixes = ["172.16.0.0/12"]
    }
    default_action_allow = {}
  }

  any_server = {}
}
```

## Next pages

- [Examples](resources--service_policy--examples.md)
- [xcsh_service_policy](../resources/service_policy.md)
