---
page_title: "xcsh_public_ip"
subcategory: ""
description: "Manages a Public IP resource in F5 Distributed Cloud for get public_ip will get the object from the storage backend for namespace metadata.namespace. configuration. (read-only data source)"
xcsh_docs: {"aliases": ["public ip"], "body_bytes": 1406, "body_sha256": "sha256:79211e99c1f1aa0867a406e7209e9095a5da36263e4f3cd3516252f0a096f015", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:public_ip:reference", "xcsh-docs:data-sources:public_ip:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:public_ip:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:public_ip:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/public_ip/index.md", "product": "distributed-cloud", "provider_name": "public_ip", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1031112030301010-2311310020031031-2211013213012022-1133000123132320-1321011231003203-2132100011003132-1013002332020322-0201311232021222", "registry_path": "docs/data-sources/public_ip.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/public_ip/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Public IP resource in F5 Distributed Cloud for get public_ip will get the object from the storage backend for namespace metadata.namespace. configuration. (read-only data source)", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": [], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_public_ip

Breadcrumbs:

- xcsh_public_ip

Manages a Public IP resource in F5 Distributed Cloud for get public\_ip will get the object from the
storage backend for namespace metadata.namespace. configuration. (read-only data source)

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# PublicIP Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing PublicIP by name
data "xcsh_public_ip" "example" {
  name      = "example-public-ip"
  namespace = "staging"
}

output "public_ip_id" {
  value = data.xcsh_public_ip.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/examples/)
