---
page_title: "xcsh_cloud_link"
subcategory: ""
description: "xcsh_cloud_link for xcsh_cloud_link."
xcsh_docs: {"aliases": [], "body_bytes": 1300, "body_sha256": "sha256:6a42a98c07cd0ea23f745f45186a5c66b59e1b396b466bf1027a107685c36b0c", "child_ids": ["xcsh-docs:data-sources:cloud_link:reference", "xcsh-docs:data-sources:cloud_link:examples"], "collection_id": "xcsh-docs:data-sources:cloud_link:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_link:fundamentals", "parent_id": null, "path": "documentation/data-sources/cloud_link/index.md", "provider_name": "cloud_link", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_link/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_cloud_link for xcsh_cloud_link.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_cloud_link

Breadcrumbs:

- xcsh_cloud_link

Manages new CloudLink with configured parameters in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudLink Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudLink by name
data "xcsh_cloud_link" "example" {
  name      = "example-cloud-link"
  namespace = "staging"
}

output "cloud_link_id" {
  value = data.xcsh_cloud_link.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/examples/)
