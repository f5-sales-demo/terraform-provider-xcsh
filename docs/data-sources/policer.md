---
page_title: "xcsh_policer"
subcategory: ""
description: "xcsh_policer for xcsh_policer."
xcsh_docs: {"aliases": [], "body_bytes": 1183, "body_sha256": "sha256:9dc04013c9be0bc040702dce37f619d2f6621402cb3064de596af9deea1da999", "canonical_id": "xcsh-docs:data-sources:policer:fundamentals", "child_ids": ["xcsh-docs:data-sources:policer:reference", "xcsh-docs:data-sources:policer:examples"], "collection_id": "xcsh-docs:data-sources:policer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policer:fundamentals", "parent_id": null, "path": "docs/data-sources/policer.md", "provider_name": "policer", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policer/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_policer for xcsh_policer.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["policerCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_policer

Breadcrumbs:

- xcsh_policer

Manages new policer with traffic rate limits in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Policer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Policer by name
data "xcsh_policer" "example" {
  name      = "example-policer"
  namespace = "staging"
}

output "policer_id" {
  value = data.xcsh_policer.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--policer--reference.md)
- [Examples](../guides/data-sources--policer--examples.md)
