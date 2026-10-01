---
page_title: "xcsh_namespace"
subcategory: ""
description: "xcsh_namespace for xcsh_namespace."
xcsh_docs: {"aliases": [], "body_bytes": 1226, "body_sha256": "sha256:0dd28f0a0e0249e9fe3f35de2ff351c9f148a17781309da650e451c117f17980", "canonical_id": "xcsh-docs:data-sources:namespace:fundamentals", "child_ids": ["xcsh-docs:data-sources:namespace:reference", "xcsh-docs:data-sources:namespace:examples"], "collection_id": "xcsh-docs:data-sources:namespace:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:namespace:fundamentals", "parent_id": null, "path": "docs/data-sources/namespace.md", "provider_name": "namespace", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "publishing_destination": "registry", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/namespace/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_namespace for xcsh_namespace.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["namespaceCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_namespace

Breadcrumbs:

- xcsh_namespace

Manages new namespace. Name of the object is name of the name space in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](../guides/data-sources--namespace--reference.md)
- [Examples](../guides/data-sources--namespace--examples.md)
