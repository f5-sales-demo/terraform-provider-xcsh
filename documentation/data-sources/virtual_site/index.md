---
page_title: "xcsh_virtual_site"
subcategory: "Infrastructure"
description: "xcsh_virtual_site for xcsh_virtual_site."
xcsh_docs: {"aliases": [], "body_bytes": 1253, "body_sha256": "sha256:f5173cc83ac5bfde15ff5e651f201e089d1698b94938118b4c91fffb2d6bedcf", "child_ids": ["xcsh-docs:data-sources:virtual_site:reference", "xcsh-docs:data-sources:virtual_site:examples"], "collection_id": "xcsh-docs:data-sources:virtual_site:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:virtual_site:fundamentals", "parent_id": null, "path": "documentation/data-sources/virtual_site/index.md", "provider_name": "virtual_site", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/virtual_site/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_virtual_site for xcsh_virtual_site.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["virtual_siteCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# xcsh_virtual_site

Breadcrumbs:

- xcsh_virtual_site

Manages virtual site object in given namespace in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualSite Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing VirtualSite by name
data "xcsh_virtual_site" "example" {
  name      = "example-virtual-site"
  namespace = "staging"
}

output "virtual_site_id" {
  value = data.xcsh_virtual_site.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_site/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/virtual_site/examples/)
