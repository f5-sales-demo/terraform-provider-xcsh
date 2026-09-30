---
page_title: "xcsh_data_type"
subcategory: ""
description: "xcsh_data_type for xcsh_data_type."
xcsh_docs: {"aliases": [], "body_bytes": 1475, "body_sha256": "sha256:f2b565be12b47c77b23a3979828a4f1b47ebf228f4db41f755ac12afeb974610", "child_ids": ["xcsh-docs:resources:data_type:reference", "xcsh-docs:resources:data_type:examples", "xcsh-docs:resources:data_type:import", "xcsh-docs:resources:data_type:timeouts"], "collection_id": "xcsh-docs:resources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_type:fundamentals", "parent_id": null, "path": "documentation/resources/data_type/index.md", "provider_name": "data_type", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "resources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_data_type for xcsh_data_type.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_data_type

Breadcrumbs:

- xcsh_data_type

Manages data\_type creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DataType Resource Example
# Manages data_type creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataType configuration
resource "xcsh_data_type" "example" {
  name      = "example-data-type"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/data_type/lifecycle/timeouts/)
