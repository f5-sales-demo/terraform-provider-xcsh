---
page_title: "xcsh_securemesh_site"
subcategory: ""
description: "xcsh_securemesh_site for xcsh_securemesh_site."
xcsh_docs: {"aliases": [], "body_bytes": 1300, "body_sha256": "sha256:814389fe697e6be3d45d97286389dd2b788f0e22f5be8665153e90cf79a652f2", "child_ids": ["xcsh-docs:data-sources:securemesh_site:reference", "xcsh-docs:data-sources:securemesh_site:examples"], "collection_id": "xcsh-docs:data-sources:securemesh_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:securemesh_site:fundamentals", "parent_id": null, "path": "documentation/data-sources/securemesh_site/index.md", "provider_name": "securemesh_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/securemesh_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_securemesh_site for xcsh_securemesh_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["securemesh_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_securemesh_site

Breadcrumbs:

- xcsh_securemesh_site

Manages a Securemesh Site resource in F5 Distributed Cloud for deploying secure mesh edge sites with
distributed security.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecuremeshSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecuremeshSite by name
data "xcsh_securemesh_site" "example" {
  name      = "example-securemesh-site"
  namespace = "staging"
}

output "securemesh_site_id" {
  value = data.xcsh_securemesh_site.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/securemesh_site/examples/)
