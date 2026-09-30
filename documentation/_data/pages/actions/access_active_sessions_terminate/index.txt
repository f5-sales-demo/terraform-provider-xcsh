---
page_title: "xcsh_access_active_sessions_terminate"
subcategory: ""
description: "xcsh_access_active_sessions_terminate for xcsh_access_active_sessions_terminate."
xcsh_docs: {"aliases": [], "body_bytes": 1267, "body_sha256": "sha256:98f2fadb438551c565b0d5d45447ac3db02660164e3339813bfada5bf7f058ba", "child_ids": ["xcsh-docs:actions:access_active_sessions_terminate:reference", "xcsh-docs:actions:access_active_sessions_terminate:examples", "xcsh-docs:actions:access_active_sessions_terminate:lifecycle"], "collection_id": "xcsh-docs:actions:access_active_sessions_terminate:collection", "completeness": "complete", "id": "xcsh-docs:actions:access_active_sessions_terminate:fundamentals", "parent_id": null, "path": "documentation/actions/access_active_sessions_terminate/index.md", "provider_name": "access_active_sessions_terminate", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "actions", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/actions/access_active_sessions_terminate/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_access_active_sessions_terminate for xcsh_access_active_sessions_terminate.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_access_active_sessions_terminate

Breadcrumbs:

- xcsh_access_active_sessions_terminate

Resource creation operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AccessActiveSessionsTerminate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_access_active_sessions_terminate" "example" {
  config {
    namespace = "example-value"
  }
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/examples/)
- [Lifecycle](https://f5-sales-demo.github.io/terraform-provider-xcsh/actions/access_active_sessions_terminate/lifecycle/)
