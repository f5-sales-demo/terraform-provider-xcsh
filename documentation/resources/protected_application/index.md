---
page_title: "xcsh_protected_application"
subcategory: ""
description: "xcsh_protected_application for xcsh_protected_application."
xcsh_docs: {"aliases": [], "body_bytes": 1516, "body_sha256": "sha256:7452221eaa7c1adb108e3bbf20ca49982c4c63792b4a11cb6a829ce28e58f6f2", "child_ids": ["xcsh-docs:resources:protected_application:reference", "xcsh-docs:resources:protected_application:examples", "xcsh-docs:resources:protected_application:import", "xcsh-docs:resources:protected_application:timeouts"], "collection_id": "xcsh-docs:resources:protected_application:collection", "completeness": "complete", "id": "xcsh-docs:resources:protected_application:fundamentals", "parent_id": null, "path": "documentation/resources/protected_application/index.md", "provider_name": "protected_application", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/protected_application/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_protected_application for xcsh_protected_application.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["protected_applicationCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_protected_application

Breadcrumbs:

- xcsh_protected_application

Manages applications protected by Bot Defense in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/protected_application/lifecycle/timeouts/)
