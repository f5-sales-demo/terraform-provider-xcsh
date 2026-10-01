---
page_title: "Allow list"
subcategory: "Security"
description: "Allow list for xcsh_service_policy."
xcsh_docs: {"aliases": [], "body_bytes": 1411, "body_sha256": "sha256:e40e126aaee7526ea499a5acc1ec25e2edda58c766c5a6fcda16d90c1165d9de", "child_ids": [], "collection_id": "xcsh-docs:resources:service_policy:collection", "completeness": "complete", "evidence": {"attribution": "Acceptance-test-derived fixture; no new live API execution is claimed.", "outcome": "valid configuration", "sha256": "sha256:e3e5a70e5d1b6c7d8436b9826fb90da18093ff0f1b51c42b00f77c75bb40e353", "source_path": "examples/resources/xcsh_service_policy/allow-list.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:service_policy:example:allow-list", "parent_id": "xcsh-docs:resources:service_policy:examples", "path": "documentation/resources/service_policy/examples/allow-list/index.md", "provider_name": "service_policy", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "role": "example", "schema_path": ["allow-list"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/service_policy/examples/allow-list/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Allow list for xcsh_service_policy.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["service_policyCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Allow list

Breadcrumbs:

- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/examples/)
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

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/examples/)
- [xcsh_service_policy](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/service_policy/)
