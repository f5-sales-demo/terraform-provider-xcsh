---
page_title: "xcsh_subnet"
subcategory: ""
description: "Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an interface of a vm/pod. it is created in user or shared namespace. configuration."
xcsh_docs: {"aliases": ["subnet"], "body_bytes": 1367, "body_sha256": "sha256:bd2cbe19e9fde15954868961001bfbed52dc6c8bf4c9100c5762de7a56667d58", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:subnet:reference", "xcsh-docs:data-sources:subnet:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:subnet:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/subnet/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210", "registry_path": "docs/data-sources/subnet.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/subnet/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an interface of a vm/pod. it is created in user or shared namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["subnetCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_subnet

Breadcrumbs:

- xcsh_subnet

Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an
interface of a vm/pod. it is created in user or shared namespace. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Subnet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Subnet by name
data "xcsh_subnet" "example" {
  name      = "example-subnet"
  namespace = "staging"
}

output "subnet_id" {
  value = data.xcsh_subnet.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/subnet/examples/)
