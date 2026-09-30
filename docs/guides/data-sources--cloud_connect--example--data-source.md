---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_cloud_connect."
xcsh_docs: {"aliases": [], "body_bytes": 987, "body_sha256": "sha256:800069f703c55018a76cfa33826eb847245bd797fa2cb082e379e2f025eb86b6", "canonical_id": "xcsh-docs:data-sources:cloud_connect:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:cloud_connect:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:eed109fcd2f0afc5835242d94d0d30137f2b13d899b31f9e138f00b05427dcd6", "source_path": "examples/data-sources/xcsh_cloud_connect/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:cloud_connect:example:data-source", "parent_id": "xcsh-docs:data-sources:cloud_connect:examples", "path": "docs/guides/data-sources--cloud_connect--example--data-source.md", "provider_name": "cloud_connect", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/cloud_connect/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_cloud_connect.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["cloud_connectCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
- [Examples](data-sources--cloud_connect--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_cloud_connect/data-source.tf`; digest `sha256:eed109fcd2f0afc5835242d94d0d30137f2b13d899b31f9e138f00b05427dcd6`.

```terraform
# CloudConnect Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CloudConnect by name
data "xcsh_cloud_connect" "example" {
  name      = "example-cloud-connect"
  namespace = "staging"
}

output "cloud_connect_id" {
  value = data.xcsh_cloud_connect.example.id
}
```

## Next pages

- [Examples](data-sources--cloud_connect--examples.md)
- [xcsh_cloud_connect](../data-sources/cloud_connect.md)
