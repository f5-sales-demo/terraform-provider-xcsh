---
page_title: "xcsh_cminstance"
subcategory: ""
description: "xcsh_cminstance for xcsh_cminstance."
xcsh_docs: {"aliases": [], "body_bytes": 1350, "body_sha256": "sha256:dde4b1d2d9f19e3248715e619419667d5e5d5a4e4b074a52e8b7831117c86a56", "canonical_id": "xcsh-docs:resources:cminstance:fundamentals", "child_ids": ["xcsh-docs:resources:cminstance:reference", "xcsh-docs:resources:cminstance:examples", "xcsh-docs:resources:cminstance:import", "xcsh-docs:resources:cminstance:timeouts"], "collection_id": "xcsh-docs:resources:cminstance:collection", "completeness": "complete", "id": "xcsh-docs:resources:cminstance:fundamentals", "parent_id": null, "path": "docs/resources/cminstance.md", "provider_name": "cminstance", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/cminstance/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_cminstance for xcsh_cminstance.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cminstanceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_cminstance

Breadcrumbs:

- xcsh_cminstance

Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Cminstance Resource Example
# Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Cminstance configuration
resource "xcsh_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "staging"

  port     = 1
  username = "example-value"
}
```

## Root configuration

Required root properties: `name`, `namespace`, `port`, `username`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--cminstance--reference.md)
- [Examples](../guides/resources--cminstance--examples.md)
- [Import](../guides/resources--cminstance--import.md)
- [Timeouts](../guides/resources--cminstance--timeouts.md)
