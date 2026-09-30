---
page_title: "xcsh_access_active_sessions"
subcategory: ""
description: "xcsh_access_active_sessions for xcsh_access_active_sessions."
xcsh_docs: {"aliases": [], "body_bytes": 1176, "body_sha256": "sha256:aab1f5958676eaacb956094bdcf728519b471d62a32963308b7851af8b75137c", "child_ids": ["xcsh-docs:data-sources:access_active_sessions:reference", "xcsh-docs:data-sources:access_active_sessions:examples"], "collection_id": "xcsh-docs:data-sources:access_active_sessions:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:access_active_sessions:fundamentals", "parent_id": null, "path": "documentation/data-sources/access_active_sessions/index.md", "provider_name": "access_active_sessions", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/access_active_sessions/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_access_active_sessions for xcsh_access_active_sessions.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_access_active_sessions

Breadcrumbs:

- xcsh_access_active_sessions

Resource retrieval operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AccessActiveSessions DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_access_active_sessions" "example" {
  namespace = "example-value"
}

output "access_active_sessions_result" {
  value = data.xcsh_access_active_sessions.example
}
```

## Root configuration

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/access_active_sessions/examples/)
