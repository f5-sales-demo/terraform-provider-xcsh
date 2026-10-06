---
page_title: "xcsh_public_ip"
subcategory: ""
description: "Reads an existing public IP object from the requested namespace."
xcsh_docs: {"aliases": ["public ip"], "body_bytes": 1294, "body_sha256": "sha256:505a5b611e852e4ea154b28343107088c787fc239a7b79fa7bd4da2bf6ad41ea", "capabilities": [], "category": null, "child_ids": ["xcsh-docs:data-sources:public_ip:reference", "xcsh-docs:data-sources:public_ip:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": [], "status": "unresolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:public_ip:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:public_ip:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/public_ip/index.md", "product": "distributed-cloud", "provider_name": "public_ip", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-1031112030301010-2311310020031031-2211013213012022-1133000123132320-1321011231003203-2132100011003132-1013002332020322-0201311232021222", "registry_path": "docs/data-sources/public_ip.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/public_ip/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reads an existing public IP object from the requested namespace.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": [], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_public_ip

Breadcrumbs:

- xcsh_public_ip

Reads an existing public IP object from the requested namespace.

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

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/public_ip/examples/)
