---
page_title: "xcsh_cdn_purge_command"
subcategory: ""
description: "xcsh_cdn_purge_command for xcsh_cdn_purge_command."
xcsh_docs: {"aliases": [], "body_bytes": 1407, "body_sha256": "sha256:855776324132689f550d511e90253edfbbe1eab152f0155b1e1e987b68cf5c5c", "child_ids": ["xcsh-docs:data-sources:cdn_purge_command:reference", "xcsh-docs:data-sources:cdn_purge_command:examples"], "collection_id": "xcsh-docs:data-sources:cdn_purge_command:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cdn_purge_command:fundamentals", "parent_id": null, "path": "documentation/data-sources/cdn_purge_command/index.md", "provider_name": "cdn_purge_command", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cdn_purge_command/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_cdn_purge_command for xcsh_cdn_purge_command.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_purge_commandCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_cdn_purge_command

Breadcrumbs:

- xcsh_cdn_purge_command

Manages a CDN Purge Command resource in F5 Distributed Cloud for cdn purge command specification.
configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_purge_command/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cdn_purge_command/examples/)
