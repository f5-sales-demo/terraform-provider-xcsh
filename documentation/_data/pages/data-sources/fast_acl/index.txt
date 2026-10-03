---
page_title: "xcsh_fast_acl"
subcategory: ""
description: "Manages object, object contains rules to protect site from denial of service It has destination{destination IP, destination port) and references to in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["fast acl"], "body_bytes": 1366, "body_sha256": "sha256:f19000995cafd08299d3207b69c7aeff3c41e7309cc73193c29686c0c2c4b26d", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:fast_acl:reference", "xcsh-docs:data-sources:fast_acl:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:fast_acl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:fast_acl:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/fast_acl/index.md", "product": "distributed-cloud", "provider_name": "fast_acl", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2020211210023323-3001311010111301-2222021113012213-3301233223211130-3313023120030023-1233001003230231-1320322003303313-3301101022102210", "registry_path": "docs/data-sources/fast_acl.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/fast_acl/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages object, object contains rules to protect site from denial of service It has destination{destination IP, destination port) and references to in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["fast_aclCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_fast_acl

Breadcrumbs:

- xcsh_fast_acl

Manages object, object contains rules to protect site from denial of service It has
destination\{destination IP, destination port) and references to in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FastACL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing FastACL by name
data "xcsh_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"
}

output "fast_acl_id" {
  value = data.xcsh_fast_acl.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/fast_acl/examples/)
