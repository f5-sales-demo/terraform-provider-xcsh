---
page_title: "xcsh_cminstance"
subcategory: ""
description: "xcsh_cminstance for xcsh_cminstance."
xcsh_docs: {"aliases": [], "body_bytes": 1247, "body_sha256": "sha256:b1cac04016ec2b87d56f1fb84c22e5d98dba09f7c61bed774083f4ba5a7a0e56", "canonical_id": "xcsh-docs:data-sources:cminstance:fundamentals", "child_ids": ["xcsh-docs:data-sources:cminstance:reference", "xcsh-docs:data-sources:cminstance:examples"], "collection_id": "xcsh-docs:data-sources:cminstance:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cminstance:fundamentals", "parent_id": null, "path": "docs/data-sources/cminstance.md", "provider_name": "cminstance", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cminstance/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_cminstance for xcsh_cminstance.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cminstanceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_cminstance

Breadcrumbs:

- xcsh_cminstance

Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed
Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Cminstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Cminstance by name
data "xcsh_cminstance" "example" {
  name      = "example-cminstance"
  namespace = "staging"
}

output "cminstance_id" {
  value = data.xcsh_cminstance.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--cminstance--reference.md)
- [Examples](../guides/data-sources--cminstance--examples.md)
