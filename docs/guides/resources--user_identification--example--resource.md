---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_user_identification."
xcsh_docs: {"aliases": [], "body_bytes": 1167, "body_sha256": "sha256:2d760cbcdeb247ab262dadeb5efdac9962e7a0b9b69fdf606c51aabfad9176a5", "canonical_id": "xcsh-docs:resources:user_identification:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:user_identification:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a74347515a80c8e0aaaac40831664823cc6a3928018d3932b931b768606d6cb2", "source_path": "examples/resources/xcsh_user_identification/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:user_identification:example:resource", "parent_id": "xcsh-docs:resources:user_identification:examples", "path": "docs/guides/resources--user_identification--example--resource.md", "provider_name": "user_identification", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/user_identification/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_user_identification.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["user_identificationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_user_identification](../resources/user_identification.md)
- [Examples](resources--user_identification--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_user_identification/resource.tf`; digest `sha256:a74347515a80c8e0aaaac40831664823cc6a3928018d3932b931b768606d6cb2`.

```terraform
# UserIdentification Resource Example
# Manages user_identification creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic UserIdentification configuration
resource "xcsh_user_identification" "example" {
  name      = "example-user-identification"
  namespace = "staging"
}
```

## Next pages

- [Examples](resources--user_identification--examples.md)
- [xcsh_user_identification](../resources/user_identification.md)
