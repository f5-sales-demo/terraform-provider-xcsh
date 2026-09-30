---
page_title: "Resource"
subcategory: ""
description: "Resource for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 1247, "body_sha256": "sha256:99f7ee3ace5d8e482b1bfba7844a55ea5baa1c89d27619c3f9fdb6205b9fb26a", "child_ids": [], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:868be3bcffc6df67dc51247c65c90b154557ee6fbafff96f6968ad86d503ff19", "source_path": "examples/resources/xcsh_protected_application/resource.tf", "validation": "terraform validate"}, "id": "xcsh-docs:resources:protected_application:example:resource", "parent_id": "xcsh-docs:resources:protected_application:examples", "path": "documentation/resources/protected_application/examples/resource/index.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "example", "schema_path": ["resource"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/examples/resource/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Resource for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Resource

Breadcrumbs:

- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/examples/)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_protected_application/resource.tf`; digest `sha256:868be3bcffc6df67dc51247c65c90b154557ee6fbafff96f6968ad86d503ff19`.

```terraform
# ProtectedApplication Resource Example
# Manages applications protected by Bot Defense in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtectedApplication configuration
resource "xcsh_protected_application" "example" {
  name      = "example-protected-application"
  namespace = "staging"
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/examples/)
- [xcsh_protected_application](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/)
