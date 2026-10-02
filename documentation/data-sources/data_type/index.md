---
page_title: "xcsh_data_type"
subcategory: ""
description: "Manages data_type creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["backend servers", "data type", "origin servers", "upstream servers"], "body_bytes": 1327, "body_sha256": "sha256:cde62d053f5cb1e5b0136b943fdf8338bcd6536cd1f91cb247fcbb0cb520257c", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:data_type:reference", "xcsh-docs:data-sources:data_type:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:data_type:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:data_type:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/data_type/index.md", "product": "distributed-cloud", "provider_name": "data_type", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2002131311032121-2101212221230130-1301301323201121-2222100003011122-3120101301100102-0212113321001310-3010331111332022-0002332000000133", "registry_path": "docs/data-sources/data_type.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/data_type/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages data_type creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["data_typeCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_data_type

Breadcrumbs:

- xcsh_data_type

Manages data\_type creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DataType Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DataType by name
data "xcsh_data_type" "example" {
  name      = "example-data-type"
  namespace = "staging"
}

output "data_type_id" {
  value = data.xcsh_data_type.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/data_type/examples/)
