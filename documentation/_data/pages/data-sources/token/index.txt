---
page_title: "xcsh_token"
subcategory: "Identity"
description: "Manages new token. Token object is used to manage site admission. User must generate token before provisioning and pass this token to site during it's registration in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["token"], "body_bytes": 1387, "body_sha256": "sha256:345afe43eb20dad1449af766a3c31246c94e70ecfef36acf7b2be77ee48fd6a4", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:token:reference", "xcsh-docs:data-sources:token:examples"], "classification": {"rules_sha256": "sha256:9636a66231d73f64c001eb778c187ea542d4b4c342eb731c651d4b3b89fd2764", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:token:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:token:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/token/index.md", "product": "distributed-cloud", "provider_name": "token", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-2000130020133110-1113012332320202-2223121002312130-2001210311231111-3323332201023323-2311120322030210-1131312201010133-0322033200102202", "registry_path": "docs/data-sources/token.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/token/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages new token. Token object is used to manage site admission. User must generate token before provisioning and pass this token to site during it's registration in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["tokenCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_token

Breadcrumbs:

- xcsh_token

Manages new token. Token object is used to manage site admission. User must generate token before
provisioning and pass this token to site during it's registration in F5 Distributed Cloud.

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/token/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/token/examples/)
