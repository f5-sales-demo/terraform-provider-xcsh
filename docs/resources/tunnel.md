---
page_title: "xcsh_tunnel"
subcategory: ""
description: "xcsh_tunnel for xcsh_tunnel."
xcsh_docs: {"aliases": [], "body_bytes": 1277, "body_sha256": "sha256:f5eb03758382a76a32d681042fd665d76a147b727ee731c5b429abaf4588173d", "canonical_id": "xcsh-docs:resources:tunnel:fundamentals", "child_ids": ["xcsh-docs:resources:tunnel:reference", "xcsh-docs:resources:tunnel:examples", "xcsh-docs:resources:tunnel:import", "xcsh-docs:resources:tunnel:timeouts"], "collection_id": "xcsh-docs:resources:tunnel:collection", "completeness": "complete", "id": "xcsh-docs:resources:tunnel:fundamentals", "parent_id": null, "path": "docs/resources/tunnel.md", "provider_name": "tunnel", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/tunnel/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_tunnel for xcsh_tunnel.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tunnelCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_tunnel

Breadcrumbs:

- xcsh_tunnel

Manages tunnel in a given namespace. If one already exist it will give a error in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Tunnel Resource Example
# Manages tunnel in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Tunnel configuration
resource "xcsh_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--tunnel--reference.md)
- [Examples](../guides/resources--tunnel--examples.md)
- [Import](../guides/resources--tunnel--import.md)
- [Timeouts](../guides/resources--tunnel--timeouts.md)
