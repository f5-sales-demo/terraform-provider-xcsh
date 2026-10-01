---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_namespace."
xcsh_docs: {"aliases": [], "body_bytes": 1066, "body_sha256": "sha256:c67d9be9b670763f5d125d9696c13ae05f9c1a39463e1535682945315e050ee7", "canonical_id": "xcsh-docs:data-sources:namespace:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:namespace:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:a2418f75dbce2847a8b4c2579a5eb124cd8b027cba5de6019e63eeb934c35492", "source_path": "examples/data-sources/xcsh_namespace/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:namespace:example:data-source", "parent_id": "xcsh-docs:data-sources:namespace:examples", "path": "docs/guides/data-sources--namespace--example--data-source.md", "provider_name": "namespace", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/namespace/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_namespace.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["namespaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# Data source

Breadcrumbs:

- [xcsh_namespace](../data-sources/namespace.md)
- [Examples](data-sources--namespace--examples.md)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_namespace/data-source.tf`; digest `sha256:a2418f75dbce2847a8b4c2579a5eb124cd8b027cba5de6019e63eeb934c35492`.

```terraform
# Namespace Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Credentials are supplied externally.
provider "xcsh" {}

# Look up an existing Namespace by name
data "xcsh_namespace" "example" {
  name = "example-namespace"
}

output "namespace_id" {
  value = data.xcsh_namespace.example.id
}
```

## Next pages

- [Examples](data-sources--namespace--examples.md)
- [xcsh_namespace](../data-sources/namespace.md)
