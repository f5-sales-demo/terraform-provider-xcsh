---
page_title: "xcsh_securemesh_site_v2"
subcategory: ""
description: "xcsh_securemesh_site_v2 for xcsh_securemesh_site_v2."
xcsh_docs: {"aliases": [], "body_bytes": 1343, "body_sha256": "sha256:18c190794d01aa4de12f32ced450e92ee4ab2a1ca8f28a0af5c9549c6092dc40", "canonical_id": "xcsh-docs:data-sources:securemesh_site_v2:fundamentals", "child_ids": ["xcsh-docs:data-sources:securemesh_site_v2:reference", "xcsh-docs:data-sources:securemesh_site_v2:examples"], "collection_id": "xcsh-docs:data-sources:securemesh_site_v2:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site_v2:fundamentals", "parent_id": null, "path": "docs/data-sources/securemesh_site_v2.md", "provider_name": "securemesh_site_v2", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site_v2/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_securemesh_site_v2 for xcsh_securemesh_site_v2.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_site_v2CreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_securemesh_site_v2

Breadcrumbs:

- xcsh_securemesh_site_v2

Manages a Securemesh Site V2 resource in F5 Distributed Cloud for deploying secure mesh edge sites
with security and networking controls.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecuremeshSiteV2 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecuremeshSiteV2 by name
data "xcsh_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2"
  namespace = "system"
}

output "securemesh_site_v2_id" {
  value = data.xcsh_securemesh_site_v2.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--securemesh_site_v2--reference.md)
- [Examples](../guides/data-sources--securemesh_site_v2--examples.md)
