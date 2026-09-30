---
page_title: "Data source"
subcategory: "Networking"
description: "Data source for xcsh_endpoint."
xcsh_docs: {"aliases": [], "body_bytes": 1130, "body_sha256": "sha256:d2c739b600f7a038aa19d89b383828a305bd71c31ae81b51599772b4a6aaf1c9", "child_ids": [], "collection_id": "xcsh-docs:data-sources:endpoint:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:995586c12b63c50200cc6d03f85d5a9e36c9682dfdd230f380379061443ed02b", "source_path": "examples/data-sources/xcsh_endpoint/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:endpoint:example:data-source", "parent_id": "xcsh-docs:data-sources:endpoint:examples", "path": "documentation/data-sources/endpoint/examples/data-source/index.md", "provider_name": "endpoint", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/endpoint/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_endpoint.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["endpointCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

# Data source

Breadcrumbs:

- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/examples/)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_endpoint/data-source.tf`; digest `sha256:995586c12b63c50200cc6d03f85d5a9e36c9682dfdd230f380379061443ed02b`.

```terraform
# Endpoint Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Endpoint by name
data "xcsh_endpoint" "example" {
  name      = "example-endpoint"
  namespace = "staging"
}

output "endpoint_id" {
  value = data.xcsh_endpoint.example.id
}
```

## Next pages

- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/examples/)
- [xcsh_endpoint](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/endpoint/)
