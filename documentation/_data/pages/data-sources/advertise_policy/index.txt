---
page_title: "xcsh_advertise_policy"
subcategory: ""
description: "Manages a Advertise Policy resource in F5 Distributed Cloud for advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers. configuration."
xcsh_docs: {"aliases": ["advertise policy"], "body_bytes": 1495, "body_sha256": "sha256:da455f193619b6490d0bd43e80bc5bd4af7ac3fb21b6b514e4282e6dbb509754", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:advertise_policy:reference", "xcsh-docs:data-sources:advertise_policy:examples"], "classification": {"rules_sha256": "sha256:789b2828af1d6a9f69d7c06f4cdc4bbc58a3b1dbac4a676da8d43bc0f312b2a0", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:advertise_policy:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:advertise_policy:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/advertise_policy/index.md", "product": "distributed-cloud", "provider_name": "advertise_policy", "provider_schema_digest": "sha256:e93f7e38d36f867af35587962cdc3c5282fd19efd6d50da0ec22d86812ee056f", "provider_type": "data-sources", "registry_anchor": "canonical-1313323203033333-1300230112313332-2100102121033233-0222102011113211-1021300002113302-1302230112320222-2132203330001002-0321121130220222", "registry_path": "docs/data-sources/advertise_policy.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/advertise_policy/index.txt", "spec_pin_digest": "sha256:442a6f7ed6e6f9010cd38e0a636e997c70358deccd3ce493d233d9ed86c49d27", "summary": "Manages a Advertise Policy resource in F5 Distributed Cloud for advertise_policy object controls how and where a service represented by a given virtual_host object is advertised to consumers. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v9.0.1", "schema_components": ["advertise_policyCreateRequest"], "target_commit": "158db014109f2a838b95bccd8eb1870a39f8ca71"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_advertise_policy

Breadcrumbs:

- xcsh_advertise_policy

Manages a Advertise Policy resource in F5 Distributed Cloud for advertise\_policy object controls
how and where a service represented by a given virtual\_host object is advertised to consumers.
configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AdvertisePolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AdvertisePolicy by name
data "xcsh_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "staging"
}

output "advertise_policy_id" {
  value = data.xcsh_advertise_policy.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/advertise_policy/examples/)
