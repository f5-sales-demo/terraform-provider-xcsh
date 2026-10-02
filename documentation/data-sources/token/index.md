---
page_title: "xcsh_token"
subcategory: "Identity"
description: "Manages new token. Token object is used to manage site admission. User must generate token before provisioning and pass this token to site during it's registration in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["token"], "body_bytes": 1387, "body_sha256": "sha256:345afe43eb20dad1449af766a3c31246c94e70ecfef36acf7b2be77ee48fd6a4", "capabilities": ["identity"], "category": "identity", "child_ids": ["xcsh-docs:data-sources:token:reference", "xcsh-docs:data-sources:token:examples"], "classification": {"rules_sha256": "sha256:6be810e90af34359481d31eabd71cd76265065a170da3c5cb985e3c6e6971951", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:token:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:token:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/token/index.md", "product": "distributed-cloud", "provider_name": "token", "provider_schema_digest": "sha256:e63a07e98b4893c041a6f79be7e19c64babfe543fc6847c17641037e084e1c7c", "provider_type": "data-sources", "registry_anchor": "canonical-2000130020133110-1113012332320202-2223121002312130-2001210311231111-3323332201023323-2311120322030210-1131312201010133-0322033200102202", "registry_path": "docs/data-sources/token.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/token/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "Manages new token. Token object is used to manage site admission. User must generate token before provisioning and pass this token to site during it's registration in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["tokenCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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
