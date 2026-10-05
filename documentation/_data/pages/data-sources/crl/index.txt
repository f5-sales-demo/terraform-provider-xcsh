---
page_title: "xcsh_crl"
subcategory: ""
description: "Manages a CRL resource in F5 Distributed Cloud for api to create crl object. configuration."
xcsh_docs: {"aliases": ["crl"], "body_bytes": 1250, "body_sha256": "sha256:d967e46ea0f0df3afb8f4e58aeb4889b7b81ec980243c512f0bc87dec2c7143d", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:crl:reference", "xcsh-docs:data-sources:crl:examples"], "classification": {"rules_sha256": "sha256:4f920e935e3e9e01c1ed47fa429843ff2503dc1cd21c48cc88a4ef7a6b8b0d6a", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:crl:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:crl:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/crl/index.md", "product": "distributed-cloud", "provider_name": "crl", "provider_schema_digest": "sha256:46b91e0dfead6561ff76cfb5d9ccf48f9937eb084797761eb9aa6a2759342733", "provider_type": "data-sources", "registry_anchor": "canonical-1303332111131203-2002232212001132-0030110222210023-0210330330101332-1012330233220233-3212220211023233-3303111232113130-2302300032211003", "registry_path": "docs/data-sources/crl.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/crl/index.txt", "spec_pin_digest": "sha256:92d57ca4044c441eb33576dc672fc743258240803e45b1b9796649248e2b5dff", "summary": "Manages a CRL resource in F5 Distributed Cloud for api to create crl object. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v11.0.0", "schema_components": ["crlCreateRequest"], "target_commit": "93e701d8415eb6d501534c03289fc2180b1ad3d2"}}
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
