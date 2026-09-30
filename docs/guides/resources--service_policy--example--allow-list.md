---
page_title: "Allow list"
subcategory: "Security"
description: "Allow list for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1106, "body_sha256": "sha256:ff6aa1d827b81dc95ee232cb4783d0247f640abbfb595cf0f3ac4115dc7f9712", "canonical_id": "xcsh-docs:resources:service_policy:example:allow-list", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:e3e5a70e5d1b6c7d8436b9826fb90da18093ff0f1b51c42b00f77c75bb40e353", "source_path": "examples/resources/xcsh_service_policy/allow-list.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:service_policy:example:allow-list", "parent_id": "xcsh-docs:resources:service_policy:examples", "path": "docs/guides/resources--service_policy--example--allow-list.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["allow-list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/examples/allow-list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Allow list for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Allow list

Breadcrumbs:

- [xcsh_service_policy](../resources/service_policy.md)
- [Examples](resources--service_policy--examples.md)
- Allow list

Acceptance-test-derived fixture; no new live API execution is claimed.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_service_policy/allow-list.tf`; digest `sha256:e3e5a70e5d1b6c7d8436b9826fb90da18093ff0f1b51c42b00f77c75bb40e353`.

```terraform
# AllowList — Acceptance-test-derived Configuration
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

  # Allow list with IP prefix
  allow_list {
    prefix_list {
      prefixes = ["10.0.0.0/8", "192.168.0.0/16"]
    }
    default_action_deny = {}
  }

  # Apply to any server
  any_server = {}
}
```

## Next pages

- [Examples](resources--service_policy--examples.md)
- [xcsh_service_policy](../resources/service_policy.md)
