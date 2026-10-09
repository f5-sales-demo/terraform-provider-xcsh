---
page_title: "xcsh_network_connector"
subcategory: "Networking"
description: "Reads a Network Connector created in the system namespace."
xcsh_docs: {"aliases": ["network connector"], "body_bytes": 1484, "body_sha256": "sha256:360cccdd6ee3a3c1d1fbe179188346b05512e89f5b67ff098fd9e8706d843dbb", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:data-sources:network_connector:reference", "xcsh-docs:data-sources:network_connector:examples"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:fundamentals", "parent_id": "xcsh-docs:data-sources:xcsh:navigation", "path": "documentation/data-sources/network_connector/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:3781a6379f2e577c345b0b96b045909a7e613252151aa7230ef45f5aa7bad429", "provider_type": "data-sources", "registry_anchor": "canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301", "registry_path": "docs/data-sources/network_connector.md", "relationships": [{"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:required", "target_id": "xcsh-docs:resources:virtual_network:fundamentals", "type": "advisory"}], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/index.txt", "spec_pin_digest": "sha256:2276c84e7ee95ed330198915b02d51b561c3c6ffa847cc94557c7c69ba2b4833", "summary": "Reads a Network Connector created in the system namespace.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.3", "schema_components": ["network_connectorCreateRequest"], "target_commit": "6e75ef52298b89a53124977b4ae265f8020a04a8"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_connector

Breadcrumbs:

- xcsh_network_connector

Reads a Network Connector created in the system namespace.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `virtual_network`.

- virtual_network: Network to connect

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkConnector Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkConnector by name
data "xcsh_network_connector" "example" {
  name      = "example-network-connector"
  namespace = "staging"
}

output "network_connector_id" {
  value = data.xcsh_network_connector.example.id
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/examples/)
