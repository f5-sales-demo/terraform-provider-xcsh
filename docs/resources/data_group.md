---
page_title: "xcsh_data_group"
subcategory: ""
description: "xcsh_data_group for xcsh_data_group."
xcsh_docs: {"aliases": [], "body_bytes": 1325, "body_sha256": "sha256:42299bf4545b8c5abf14c30bd3b8993051729e2f229f07b7a1b8b487b4c400a5", "canonical_id": "xcsh-docs:resources:data_group:fundamentals", "child_ids": ["xcsh-docs:resources:data_group:reference", "xcsh-docs:resources:data_group:examples", "xcsh-docs:resources:data_group:import", "xcsh-docs:resources:data_group:timeouts"], "collection_id": "xcsh-docs:resources:data_group:collection", "completeness": "complete", "id": "xcsh-docs:resources:data_group:fundamentals", "parent_id": null, "path": "docs/resources/data_group.md", "provider_name": "data_group", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "resources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/data_group/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_data_group for xcsh_data_group.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_groupCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_data_group

Breadcrumbs:

- xcsh_data_group

Manages data group in a given namespace. If one already exists it will give an error in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DataGroup Resource Example
# Manages data group in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic DataGroup configuration
resource "xcsh_data_group" "example" {
  name      = "example-data-group"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/resources--data_group--reference.md)
- [Examples](../guides/resources--data_group--examples.md)
- [Import](../guides/resources--data_group--import.md)
- [Timeouts](../guides/resources--data_group--timeouts.md)
