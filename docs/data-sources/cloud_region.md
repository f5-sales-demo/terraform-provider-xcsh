---
page_title: "xcsh_cloud_region"
subcategory: ""
description: "xcsh_cloud_region for xcsh_cloud_region."
xcsh_docs: {"aliases": [], "body_bytes": 1185, "body_sha256": "sha256:32f291faa0164f7b35b8657415d603b2de78094f8251a3957c2a3297953411a9", "canonical_id": "xcsh-docs:data-sources:cloud_region:fundamentals", "child_ids": ["xcsh-docs:data-sources:cloud_region:reference", "xcsh-docs:data-sources:cloud_region:examples"], "collection_id": "xcsh-docs:data-sources:cloud_region:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:cloud_region:fundamentals", "parent_id": null, "path": "docs/data-sources/cloud_region.md", "provider_name": "cloud_region", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_region/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_cloud_region for xcsh_cloud_region.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": [], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_cloud_region

Breadcrumbs:

- xcsh_cloud_region

Manages a Cloud Region resource in F5 Distributed Cloud for cloud re specification. configuration.
(read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CloudRegion Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudRegion by name
data "xcsh_cloud_region" "example" {
  name      = "example-cloud-region"
  namespace = "staging"
}

output "cloud_region_id" {
  value = data.xcsh_cloud_region.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--cloud_region--reference.md)
- [Examples](../guides/data-sources--cloud_region--examples.md)
