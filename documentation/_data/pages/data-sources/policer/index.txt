---
page_title: "xcsh_policer"
subcategory: ""
description: "Reads Policer information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["policer"], "body_bytes": 1264, "body_sha256": "sha256:d26d50475941808f7cac6656874090bc1c9a82bbf1e22a2b4a03de510ef68f5b", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:policer:reference", "xcsh-docs:data-sources:policer:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:policer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policer:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/policer/index.md", "product": "distributed-cloud", "provider_name": "policer", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-3330032030013021-1312303333232323-1331010000013021-1132031333330232-2302330331211311-0310120132213301-2033021113022132-2232102121103000", "registry_path": "docs/data-sources/policer.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policer/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reads Policer information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["policerCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_policer

Breadcrumbs:

- xcsh_policer

Reads Policer information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Policer Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Policer by name
data "xcsh_policer" "example" {
  name      = "example-policer"
  namespace = "staging"
}

output "policer_id" {
  value = data.xcsh_policer.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policer/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policer/examples/)
