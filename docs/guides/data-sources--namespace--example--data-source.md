---
page_title: "Data source"
subcategory: ""
description: "Data source for xcsh_namespace."
xcsh_docs: {"aliases": [], "body_bytes": 1036, "body_sha256": "sha256:946597b92a445b8ced340b4670e3503d1047885d34aff416a32c641a401abffa", "canonical_id": "xcsh-docs:data-sources:namespace:example:data-source", "child_ids": [], "collection_id": "xcsh-docs:data-sources:namespace:collection", "completeness": "complete", "evidence": {"attribution": "Schema-derived minimal configuration validated with the checked-out provider.", "outcome": "valid configuration", "sha256": "sha256:25f104065790b7b905d53c41ec3034115fb2bd1d74237a7b3ed84110191752a1", "source_path": "examples/data-sources/xcsh_namespace/data-source.tf", "validation": "terraform validate"}, "id": "xcsh-docs:data-sources:namespace:example:data-source", "parent_id": "xcsh-docs:data-sources:namespace:examples", "path": "docs/guides/data-sources--namespace--example--data-source.md", "provider_name": "namespace", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "example", "schema_path": ["data-source"], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/namespace/examples/data-source/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Data source for xcsh_namespace.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["namespaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

Source: `examples/data-sources/xcsh_namespace/data-source.tf`; digest `sha256:25f104065790b7b905d53c41ec3034115fb2bd1d74237a7b3ed84110191752a1`.

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

# Look up an existing Namespace by name
data "xcsh_namespace" "example" {
  name      = "example-namespace"
  namespace = "staging"
}

output "namespace_id" {
  value = data.xcsh_namespace.example.id
}
```

## Next pages

- [Examples](data-sources--namespace--examples.md)
- [xcsh_namespace](../data-sources/namespace.md)
