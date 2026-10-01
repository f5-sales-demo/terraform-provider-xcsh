---
page_title: "xcsh_access_active_session"
subcategory: ""
description: "xcsh_access_active_session for xcsh_access_active_session."
xcsh_docs: {"aliases": [], "body_bytes": 1218, "body_sha256": "sha256:806d48edc3d9bda7e18ab9e001c920f8acecff0a85a84765db8607c096070b34", "canonical_id": "xcsh-docs:data-sources:access_active_session:fundamentals", "child_ids": ["xcsh-docs:data-sources:access_active_session:reference", "xcsh-docs:data-sources:access_active_session:examples"], "collection_id": "xcsh-docs:data-sources:access_active_session:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:access_active_session:fundamentals", "parent_id": null, "path": "docs/data-sources/access_active_session.md", "provider_name": "access_active_session", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/access_active_session/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_access_active_session for xcsh_access_active_session.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_access_active_session

Breadcrumbs:

- xcsh_access_active_session

Resource retrieval operation.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AccessActiveSession DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_access_active_session" "example" {
  id        = "example-value"
  namespace = "example-value"
}

output "access_active_session_result" {
  value = data.xcsh_access_active_session.example
}
```

## Root configuration

Required root properties: `id`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--access_active_session--reference.md)
- [Examples](../guides/data-sources--access_active_session--examples.md)
