---
page_title: "xcsh_network_interface"
subcategory: ""
description: "Reads network device interface configuration in the system namespace."
xcsh_docs: {"aliases": ["network interface"], "body_bytes": 1379, "body_sha256": "sha256:43beb29647e770483c90c903bd3ca11a5f1d2b38e69eacfa2c435bde72e5a1b8", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_interface:reference", "xcsh-docs:data-sources:network_interface:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-domain"}, "collection_id": "xcsh-docs:data-sources:network_interface:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_interface:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_interface/index.md", "product": "distributed-cloud", "provider_name": "network_interface", "provider_schema_digest": "sha256:a43a21ec82c017926d35acc6efc3b0a8740fd394357a09aa0e63e65503eb7aff", "provider_type": "data-sources", "registry_anchor": "canonical-0223201022310111-1331122211222221-0131301133002101-2112333331010103-2323032232200011-1030201303202000-3333303002001111-1111003100202132", "registry_path": "docs/data-sources/network_interface.md", "relationships": [], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_interface/index.txt", "spec_pin_digest": "sha256:d5e8f282045e20f8be00960c9e0ed5287ab3532b8da7d4b559f1cea17fd51bca", "summary": "Reads network device interface configuration in the system namespace.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.9", "schema_components": ["network_interfaceCreateRequest"], "target_commit": "e04c5d45f7a4e103e2dc84bf52b767182b4de3bc"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_interface

Breadcrumbs:

- xcsh_network_interface

Reads network device interface configuration in the system namespace.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkInterface Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkInterface by name
data "xcsh_network_interface" "example" {
  name      = "example-network-interface"
  namespace = "staging"
}

output "network_interface_id" {
  value = data.xcsh_network_interface.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_interface/examples/)
