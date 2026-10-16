---
page_title: "xcsh_trusted_ca_list"
subcategory: ""
description: "Reads Trusted Ca List information from F5 Distributed Cloud."
xcsh_docs: {"aliases": ["trusted ca list"], "body_bytes": 1348, "body_sha256": "sha256:5ad1899c8f175d57d82a0365e1d1ebe03bcf9db91d98c99a87d1adab4b80c282", "capabilities": ["security"], "category": "security", "child_ids": ["xcsh-docs:data-sources:trusted_ca_list:reference", "xcsh-docs:data-sources:trusted_ca_list:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:trusted_ca_list:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:trusted_ca_list:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/trusted_ca_list/index.md", "product": "distributed-cloud", "provider_name": "trusted_ca_list", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-3231112200302210-0000022221223030-0200201111213211-3101031133123233-3133122221310202-1301130023031133-3210320322301322-1320111311102122", "registry_path": "docs/data-sources/trusted_ca_list.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/trusted_ca_list/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Reads Trusted Ca List information from F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["trusted_ca_listCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_trusted_ca_list

Breadcrumbs:

- xcsh_trusted_ca_list

Reads Trusted Ca List information from F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# TrustedCAList Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TrustedCAList by name
data "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}

output "trusted_ca_list_id" {
  value = data.xcsh_trusted_ca_list.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/trusted_ca_list/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/trusted_ca_list/examples/)
