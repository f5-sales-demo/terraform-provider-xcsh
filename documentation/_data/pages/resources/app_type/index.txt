---
page_title: "xcsh_app_type"
subcategory: ""
description: "xcsh_app_type for xcsh_app_type."
xcsh_docs: {"aliases": [], "body_bytes": 1551, "body_sha256": "sha256:a856611fabf91907e4f65cffc4b5c13c00273fa2229e2644b9bacf7e29a09e74", "child_ids": ["xcsh-docs:resources:app_type:reference", "xcsh-docs:resources:app_type:examples", "xcsh-docs:resources:app_type:import", "xcsh-docs:resources:app_type:timeouts"], "collection_id": "xcsh-docs:resources:app_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:app_type:fundamentals", "parent_id": null, "path": "documentation/resources/app_type/index.md", "provider_name": "app_type", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/app_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_app_type for xcsh_app_type.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["app_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_app_type

Breadcrumbs:

- xcsh_app_type

Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppType Resource Example
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

# Basic AppType configuration
resource "xcsh_app_type" "example" {
  name      = "example-app-type"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/app_type/lifecycle/timeouts/)
