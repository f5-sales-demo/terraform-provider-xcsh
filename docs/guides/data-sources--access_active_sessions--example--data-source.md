---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_access_active_sessions."
xcsh_docs: {"aliases": [], "body_bytes": 1110, "body_sha256": "sha256:3f5bf794e2d663b196c81bc5868d27ce809757fe3d35cd65321871fad9004ddf", "canonical_id": "xcsh-docs:data-sources:access_active_sessions:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:access_active_sessions:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:9f2701c9ce37e0c0584f554dedc7733344fe85678c9dadc11b3997db86cac357", "source_path": "examples/data-sources/xcsh_access_active_sessions/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:access_active_sessions:example:data-source", "parent_id": "xcsh-docs:data-sources:access_active_sessions:examples", "path": "docs/guides/data-sources--access_active_sessions--example--data-source.md", "provider_name": "access_active_sessions", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/access_active_sessions/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_access_active_sessions.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_access_active_sessions](../data-sources/access_active_sessions.md)
- [Examples](data-sources--access_active_sessions--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_access_active_sessions/data-source.tf`; digest `sha256:9f2701c9ce37e0c0584f554dedc7733344fe85678c9dadc11b3997db86cac357`.

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

## Next pages

- [Examples](data-sources--access_active_sessions--examples.md)
- [xcsh_access_active_sessions](../data-sources/access_active_sessions.md)
