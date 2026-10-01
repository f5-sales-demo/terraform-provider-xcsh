---
page_title: "xcsh_api_discovery"
subcategory: ""
description: "xcsh_api_discovery for xcsh_api_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1432, "body_sha256": "sha256:3c9404ec859675e187b19bef8a6b14bd984476534b8914c8125931ad0b263548", "canonical_id": "xcsh-docs:resources:api_discovery:fundamentals", "child_ids": ["xcsh-docs:resources:api_discovery:reference", "xcsh-docs:resources:api_discovery:examples", "xcsh-docs:resources:api_discovery:import", "xcsh-docs:resources:api_discovery:timeouts"], "collection_id": "xcsh-docs:resources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:resources:api_discovery:fundamentals", "parent_id": null, "path": "docs/resources/api_discovery.md", "provider_name": "api_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/api_discovery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_api_discovery for xcsh_api_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_api_discovery

Breadcrumbs:

- xcsh_api_discovery

Manages API discovery creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# APIDiscovery Resource Example
# Manages API discovery creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic APIDiscovery configuration
resource "xcsh_api_discovery" "example" {
  name      = "example-api-discovery"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--api_discovery--reference.md)
- [Examples](../guides/resources--api_discovery--examples.md)
- [Import](../guides/resources--api_discovery--import.md)
- [Timeouts](../guides/resources--api_discovery--timeouts.md)
