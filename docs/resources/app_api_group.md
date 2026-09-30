---
page_title: "xcsh_app_api_group"
subcategory: ""
description: "xcsh_app_api_group for xcsh_app_api_group."
xcsh_docs: {"aliases": [], "body_bytes": 1333, "body_sha256": "sha256:da575af80ede06e2108153f6f5a44d03cfa4859a46f07706230fc5d61e56d34a", "canonical_id": "xcsh-docs:resources:app_api_group:fundamentals", "child_ids": ["xcsh-docs:resources:app_api_group:reference", "xcsh-docs:resources:app_api_group:examples", "xcsh-docs:resources:app_api_group:import", "xcsh-docs:resources:app_api_group:timeouts"], "collection_id": "xcsh-docs:resources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_api_group:fundamentals", "parent_id": null, "path": "docs/resources/app_api_group.md", "provider_name": "app_api_group", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_api_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_app_api_group for xcsh_app_api_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_app_api_group

Breadcrumbs:

- xcsh_app_api_group

Manages app\_api\_group creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppAPIGroup Resource Example
# Manages app_api_group creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppAPIGroup configuration
resource "xcsh_app_api_group" "example" {
  name      = "example-app-api-group"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--app_api_group--reference.md)
- [Examples](../guides/resources--app_api_group--examples.md)
- [Import](../guides/resources--app_api_group--import.md)
- [Timeouts](../guides/resources--app_api_group--timeouts.md)
