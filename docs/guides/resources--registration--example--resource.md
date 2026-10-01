---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_registration."
xcsh_docs: {"aliases": [], "body_bytes": 1125, "body_sha256": "sha256:f5cb9aef5e0905c57c86086ffba7ed3bbf4e8318411b296f379f63d8efe16c6a", "canonical_id": "xcsh-docs:resources:registration:example:resource", "child_ids": [], "collection_id": "xcsh-docs:resources:registration:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:63b17b62e7eb3ba883d15945e8c1f68dbbb4f71a319c4ef34eeeacc2e2bd25dc", "source_path": "examples/resources/xcsh_registration/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:registration:example:resource", "parent_id": "xcsh-docs:resources:registration:examples", "path": "docs/guides/resources--registration--example--resource.md", "provider_name": "registration", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/registration/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_registration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["registrationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Resource

Breadcrumbs:

- [xcsh_registration](../resources/registration.md)
- [Examples](resources--registration--examples.md)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_registration/resource.tf`; digest `sha256:63b17b62e7eb3ba883d15945e8c1f68dbbb4f71a319c4ef34eeeacc2e2bd25dc`.

```terraform
# Registration Resource Example
# Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this message, never used by users.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Registration configuration
resource "xcsh_registration" "example" {
  name      = "example-registration"
  namespace = "staging"

  token = "example-value"
}
```

## Next pages

- [Examples](resources--registration--examples.md)
- [xcsh_registration](../resources/registration.md)
