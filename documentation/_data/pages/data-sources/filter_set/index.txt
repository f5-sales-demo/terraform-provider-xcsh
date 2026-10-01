---
page_title: "xcsh_filter_set"
subcategory: ""
description: "xcsh_filter_set for xcsh_filter_set."
xcsh_docs: {"aliases": [], "body_bytes": 1273, "body_sha256": "sha256:dfe66cc0e7396097907cd35fea5f571c87a5d5fc3e7411651a50cc2784c94d1d", "child_ids": ["xcsh-docs:data-sources:filter_set:reference", "xcsh-docs:data-sources:filter_set:examples"], "collection_id": "xcsh-docs:data-sources:filter_set:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:filter_set:fundamentals", "parent_id": null, "path": "documentation/data-sources/filter_set/index.md", "provider_name": "filter_set", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/filter_set/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_filter_set for xcsh_filter_set.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["filter_setCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_filter_set

Breadcrumbs:

- xcsh_filter_set

Manages specification in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FilterSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FilterSet by name
data "xcsh_filter_set" "example" {
  name      = "example-filter-set"
  namespace = "staging"
}

output "filter_set_id" {
  value = data.xcsh_filter_set.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/filter_set/examples/)
