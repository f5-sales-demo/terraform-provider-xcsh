---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cdn_purge_command."
xcsh_docs: {"aliases": [], "body_bytes": 1136, "body_sha256": "sha256:2d744ef40ca2ee64ff30bf5e3eb150e4150fa7210c1f62ddcdc20e9c649ca2d4", "canonical_id": "xcsh-docs:data-sources:cdn_purge_command:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cdn_purge_command:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:25cb8095a143a10866ff868fcae4396aaa9eca15b734c77fdd0d7054bddadd8a", "source_path": "examples/data-sources/xcsh_cdn_purge_command/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cdn_purge_command:example:data-source", "parent_id": "xcsh-docs:data-sources:cdn_purge_command:examples", "path": "docs/guides/data-sources--cdn_purge_command--example--data-source.md", "provider_name": "cdn_purge_command", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_purge_command/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_cdn_purge_command.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_purge_commandCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md)
- [Examples](data-sources--cdn_purge_command--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cdn_purge_command/data-source.tf`; digest `sha256:25cb8095a143a10866ff868fcae4396aaa9eca15b734c77fdd0d7054bddadd8a`.

```terraform
# CDNPurgeCommand Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CDNPurgeCommand by name
data "xcsh_cdn_purge_command" "example" {
  name      = "example-cdn-purge-command"
  namespace = "staging"
}

output "cdn_purge_command_id" {
  value = data.xcsh_cdn_purge_command.example.id
}
```

## Next pages

- [Examples](data-sources--cdn_purge_command--examples.md)
- [xcsh_cdn_purge_command](../data-sources/cdn_purge_command.md)
