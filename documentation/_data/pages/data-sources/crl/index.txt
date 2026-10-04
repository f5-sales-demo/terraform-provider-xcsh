---
page_title: "xcsh_crl"
subcategory: ""
description: "Manages a CRL resource in F5 Distributed Cloud for api to create crl object. configuration."
xcsh_docs: {"aliases": ["crl"], "body_bytes": 1250, "body_sha256": "sha256:d967e46ea0f0df3afb8f4e58aeb4889b7b81ec980243c512f0bc87dec2c7143d", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:crl:reference", "xcsh-docs:data-sources:crl:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:crl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:crl:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/crl/index.md", "product": "distributed-cloud", "provider_name": "crl", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1303332111131203-2002232212001132-0030110222210023-0210330330101332-1012330233220233-3212220211023233-3303111232113130-2302300032211003", "registry_path": "docs/data-sources/crl.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/crl/index.txt", "spec_pin_digest": "sha256:be3d0ac7e39778bb5070994fc48800318715ca8ca0ee228725aebf23a5cb76fd", "summary": "Manages a CRL resource in F5 Distributed Cloud for api to create crl object. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v10.0.1", "schema_components": ["crlCreateRequest"], "target_commit": "ec431fb7909aae6a211468542c1e74c04122a56a"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_crl

Breadcrumbs:

- xcsh_crl

Manages a CRL resource in F5 Distributed Cloud for api to create crl object. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# CRL Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing CRL by name
data "xcsh_crl" "example" {
  name      = "example-crl"
  namespace = "staging"
}

output "crl_id" {
  value = data.xcsh_crl.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/crl/examples/)
