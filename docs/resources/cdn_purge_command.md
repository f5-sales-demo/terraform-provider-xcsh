---
page_title: "xcsh_cdn_purge_command"
subcategory: ""
description: "xcsh_cdn_purge_command for xcsh_cdn_purge_command."
xcsh_docs: {"aliases": [], "body_bytes": 1354, "body_sha256": "sha256:ea6473f12b8c07678586e180501a67f31f4168e848ff660bfe82b5d990a971c3", "canonical_id": "xcsh-docs:resources:cdn_purge_command:fundamentals", "child_ids": ["xcsh-docs:resources:cdn_purge_command:reference", "xcsh-docs:resources:cdn_purge_command:examples", "xcsh-docs:resources:cdn_purge_command:import", "xcsh-docs:resources:cdn_purge_command:timeouts"], "collection_id": "xcsh-docs:resources:cdn_purge_command:collection", "completeness": "complete", "id": "xcsh-docs:resources:cdn_purge_command:fundamentals", "parent_id": null, "path": "docs/resources/cdn_purge_command.md", "provider_name": "cdn_purge_command", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cdn_purge_command/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_cdn_purge_command for xcsh_cdn_purge_command.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cdn_purge_commandCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

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
# CDNPurgeCommand Resource Example
# Manages a CDN Purge Command resource in F5 Distributed Cloud for cdn purge command specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic CDNPurgeCommand configuration
resource "xcsh_cdn_purge_command" "example" {
  name      = "example-cdn-purge-command"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--cdn_purge_command--reference.md)
- [Examples](../guides/resources--cdn_purge_command--examples.md)
- [Import](../guides/resources--cdn_purge_command--import.md)
- [Timeouts](../guides/resources--cdn_purge_command--timeouts.md)
