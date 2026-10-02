---
page_title: "xcsh_app_api_group"
subcategory: ""
description: "Manages app_api_group creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["app api group", "backend servers", "origin servers", "upstream servers"], "body_bytes": 1621, "body_sha256": "sha256:02eea4f2c35f9cffaff88a1c77c593cdaa45eb7dff592d117719de3d03017d8e", "capabilities": ["api-management"], "category": "api-management", "child_ids": ["xcsh-docs:resources:app_api_group:reference", "xcsh-docs:resources:app_api_group:examples", "xcsh-docs:resources:app_api_group:import", "xcsh-docs:resources:app_api_group:timeouts"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:resources:app_api_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_api_group:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/app_api_group/index.md", "product": "distributed-cloud", "provider_name": "app_api_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "registry_anchor": "canonical-0211313313012010-0123200302333333-1031223302112311-0123132320001012-2111233310130311-1003020201113303-0221003031030120-2022002111202013", "registry_path": "docs/resources/app_api_group.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_api_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages app_api_group creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_api_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

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

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_api_group/lifecycle/timeouts/)
