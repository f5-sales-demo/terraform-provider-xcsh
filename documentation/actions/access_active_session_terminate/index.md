---
page_title: "xcsh_access_active_session_terminate"
subcategory: ""
description: "xcsh_access_active_session_terminate for xcsh_access_active_session_terminate."
xcsh_docs: {"aliases": [], "body_bytes": 1397, "body_sha256": "sha256:241e52a228834e9a85f07538c480113306f5ce68900d5c83460eb745fb8e78de", "child_ids": ["xcsh-docs:actions:access_active_session_terminate:reference", "xcsh-docs:actions:access_active_session_terminate:examples", "xcsh-docs:actions:access_active_session_terminate:lifecycle"], "collection_id": "xcsh-docs:actions:access_active_session_terminate:collection", "completeness": "complete", "id": "xcsh-docs:actions:access_active_session_terminate:fundamentals", "parent_id": null, "path": "documentation/actions/access_active_session_terminate/index.md", "provider_name": "access_active_session_terminate", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/access_active_session_terminate/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_access_active_session_terminate for xcsh_access_active_session_terminate.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_access_active_session_terminate

Breadcrumbs:

- xcsh_access_active_session_terminate

Resource deletion operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AccessActiveSessionTerminate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_access_active_session_terminate" "example" {
  config {
    id        = "example-value"
    namespace = "example-value"
  }
}
```

## Root configuration

Required root properties: `id`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_session_terminate/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_session_terminate/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_session_terminate/lifecycle/)
