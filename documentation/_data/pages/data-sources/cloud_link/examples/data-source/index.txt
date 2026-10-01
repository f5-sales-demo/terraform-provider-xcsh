---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cloud_link."
xcsh_docs: {"aliases": [], "body_bytes": 1253, "body_sha256": "sha256:006aacc54d24e411f99de5a535ab6e7a769ee9351f0609f509a37670f2e062ab", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cloud_link:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a7eb053580e156c686886504937855a78a126659f9e2d4b1412c03b42464b698", "source_path": "examples/data-sources/xcsh_cloud_link/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cloud_link:example:data-source", "parent_id": "xcsh-docs:data-sources:cloud_link:examples", "path": "documentation/data-sources/cloud_link/examples/data-source/index.md", "provider_name": "cloud_link", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_link/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_cloud_link.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_linkCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_link/data-source.tf`; digest `sha256:a7eb053580e156c686886504937855a78a126659f9e2d4b1412c03b42464b698`.

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

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/examples/)
- [xcsh_cloud_link](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/cloud_link/)
