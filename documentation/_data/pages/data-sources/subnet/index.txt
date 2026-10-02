---
page_title: "xcsh_subnet"
subcategory: ""
description: "Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an interface of a vm/pod. it is created in user or shared namespace. configuration."
xcsh_docs: {"aliases": ["subnet"], "body_bytes": 1367, "body_sha256": "sha256:bd2cbe19e9fde15954868961001bfbed52dc6c8bf4c9100c5762de7a56667d58", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:subnet:reference", "xcsh-docs:data-sources:subnet:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:subnet:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:subnet:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/subnet/index.md", "product": "distributed-cloud", "provider_name": "subnet", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2030233332213333-3213211200222203-3230133323032230-1012203021120130-0010113332012110-2331021111131003-0132121132331301-3312323301120210", "registry_path": "docs/data-sources/subnet.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/subnet/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an interface of a vm/pod. it is created in user or shared namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["subnetCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
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
