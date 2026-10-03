---
page_title: "xcsh_policer"
subcategory: ""
description: "Manages new policer with traffic rate limits in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["policer"], "body_bytes": 1268, "body_sha256": "sha256:18ed580f0cd4f680c4abd9b405aaf00114082df9be799788c43176a709a9e964", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:policer:reference", "xcsh-docs:data-sources:policer:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:policer:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:policer:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/policer/index.md", "product": "distributed-cloud", "provider_name": "policer", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-3330032030013021-1312303333232323-1331010000013021-1132031333330232-2302330331211311-0310120132213301-2033021113022132-2232102121103000", "registry_path": "docs/data-sources/policer.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/policer/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages new policer with traffic rate limits in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["policerCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_policer

Breadcrumbs:

- xcsh_policer

Manages new policer with traffic rate limits in F5 Distributed Cloud.

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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policer/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/policer/examples/)
