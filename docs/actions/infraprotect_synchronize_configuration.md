---
page_title: "xcsh_infraprotect_synchronize_configuration"
subcategory: ""
description: "xcsh_infraprotect_synchronize_configuration for xcsh_infraprotect_synchronize_configuration."
xcsh_docs: {"aliases": [], "body_bytes": 1183, "body_sha256": "sha256:ac47528c37255bfb6f75911e0bab9da399d8aac2fa2b0f18cdaa70810d7b455a", "canonical_id": "xcsh-docs:actions:infraprotect_synchronize_configuration:fundamentals", "child_ids": ["xcsh-docs:actions:infraprotect_synchronize_configuration:reference", "xcsh-docs:actions:infraprotect_synchronize_configuration:examples", "xcsh-docs:actions:infraprotect_synchronize_configuration:lifecycle"], "collection_id": "xcsh-docs:actions:infraprotect_synchronize_configuration:collection", "completeness": "complete", "id": "xcsh-docs:actions:infraprotect_synchronize_configuration:fundamentals", "parent_id": null, "path": "docs/actions/infraprotect_synchronize_configuration.md", "provider_name": "infraprotect_synchronize_configuration", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/infraprotect_synchronize_configuration/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_infraprotect_synchronize_configuration for xcsh_infraprotect_synchronize_configuration.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_infraprotect_synchronize_configuration

Breadcrumbs:

- xcsh_infraprotect_synchronize_configuration

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# InfraprotectSynchronizeConfiguration Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_infraprotect_synchronize_configuration" "example" {
  config {
    namespace = "example-value"
  }
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/actions--infraprotect_synchronize_configuration--reference.md)
- [Examples](../guides/actions--infraprotect_synchronize_configuration--examples.md)
- [Lifecycle](../guides/actions--infraprotect_synchronize_configuration--lifecycle.md)
