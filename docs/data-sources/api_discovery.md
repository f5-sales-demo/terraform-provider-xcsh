---
page_title: "xcsh_api_discovery"
subcategory: ""
description: "xcsh_api_discovery for xcsh_api_discovery."
xcsh_docs: {"aliases": [], "body_bytes": 1285, "body_sha256": "sha256:4a3d693b3eadaa5e87823f87403ccb4b0b49da93e8090116670581844497a1a7", "canonical_id": "xcsh-docs:data-sources:api_discovery:fundamentals", "child_ids": ["xcsh-docs:data-sources:api_discovery:reference", "xcsh-docs:data-sources:api_discovery:examples"], "collection_id": "xcsh-docs:data-sources:api_discovery:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:api_discovery:fundamentals", "parent_id": null, "path": "docs/data-sources/api_discovery.md", "provider_name": "api_discovery", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/api_discovery/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_api_discovery for xcsh_api_discovery.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["api_discoveryCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
# APIDiscovery Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing APIDiscovery by name
data "xcsh_api_discovery" "example" {
  name      = "example-api-discovery"
  namespace = "staging"
}

output "api_discovery_id" {
  value = data.xcsh_api_discovery.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--api_discovery--reference.md)
- [Examples](../guides/data-sources--api_discovery--examples.md)
