---
page_title: "xcsh_endpoint"
subcategory: "Networking"
description: "xcsh_endpoint for xcsh_endpoint."
xcsh_docs: {"aliases": [], "body_bytes": 1333, "body_sha256": "sha256:f2d747826b78a88eaa583e90780d82fecd0f0da62d4f5f6d9339727087567326", "canonical_id": "xcsh-docs:resources:endpoint:fundamentals", "child_ids": ["xcsh-docs:resources:endpoint:reference", "xcsh-docs:resources:endpoint:examples", "xcsh-docs:resources:endpoint:import", "xcsh-docs:resources:endpoint:timeouts"], "collection_id": "xcsh-docs:resources:endpoint:collection", "completeness": "complete", "id": "xcsh-docs:resources:endpoint:fundamentals", "parent_id": null, "path": "docs/resources/endpoint.md", "provider_name": "endpoint", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/endpoint/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_endpoint for xcsh_endpoint.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_endpoint

Breadcrumbs:

- xcsh_endpoint

Manages endpoint will create the object in the storage backend for namespace metadata.namespace in
F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Endpoint Resource Example
# Manages endpoint will create the object in the storage backend for namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Endpoint configuration
resource "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--endpoint--reference.md)
- [Examples](../guides/resources--endpoint--examples.md)
- [Import](../guides/resources--endpoint--import.md)
- [Timeouts](../guides/resources--endpoint--timeouts.md)
