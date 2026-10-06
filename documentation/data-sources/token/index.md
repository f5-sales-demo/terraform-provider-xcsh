---
page_title: "xcsh_token"
subcategory: "Identity"
description: "Reads a site admission token used during registration and provisioning."
xcsh_docs: {"aliases": ["token"], "body_bytes": 1283, "body_sha256": "sha256:bf10245749894e401f1c2c5651fd74c32f6adb2fdec600511646921b1e6c957b", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:token:reference", "xcsh-docs:data-sources:token:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:token:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:token:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/token/index.md", "product": "distributed-cloud", "provider_name": "token", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "data-sources", "registry_anchor": "canonical-2000130020133110-1113012332320202-2223121002312130-2001210311231111-3323332201023323-2311120322030210-1131312201010133-0322033200102202", "registry_path": "docs/data-sources/token.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/token/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Reads a site admission token used during registration and provisioning.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["tokenCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_token

Breadcrumbs:

- xcsh_token

Reads a site admission token used during registration and provisioning.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Token Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Token by name
data "xcsh_token" "example" {
  name      = "example-token"
  namespace = "system"
}

output "token_id" {
  value = data.xcsh_token.example.id
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/token/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/token/examples/)
