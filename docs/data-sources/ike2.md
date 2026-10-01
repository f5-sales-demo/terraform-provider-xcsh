---
page_title: "xcsh_ike2"
subcategory: ""
description: "xcsh_ike2 for xcsh_ike2."
xcsh_docs: {"aliases": [], "body_bytes": 1184, "body_sha256": "sha256:9d3494c3a1ce5df0af265d7b7f8e2a75ff22167744e1b70a65d013f7a655ea45", "canonical_id": "xcsh-docs:data-sources:ike2:fundamentals", "child_ids": ["xcsh-docs:data-sources:ike2:reference", "xcsh-docs:data-sources:ike2:examples"], "collection_id": "xcsh-docs:data-sources:ike2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:ike2:fundamentals", "parent_id": null, "path": "docs/data-sources/ike2.md", "provider_name": "ike2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/ike2/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_ike2 for xcsh_ike2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["ike2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_ike2

Breadcrumbs:

- xcsh_ike2

Manages a Ike2 resource in F5 Distributed Cloud for ike phase2 profile specification. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Ike2 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Ike2 by name
data "xcsh_ike2" "example" {
  name      = "example-ike2"
  namespace = "staging"
}

output "ike2_id" {
  value = data.xcsh_ike2.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--ike2--reference.md)
- [Examples](../guides/data-sources--ike2--examples.md)
