---
page_title: "xcsh_network_connector"
subcategory: "Networking"
description: "Manages a Network Connector resource in F5 Distributed Cloud for network connector is created by users in system namespace. configuration."
xcsh_docs: {"aliases": ["network connector"], "body_bytes": 1825, "body_sha256": "sha256:5007827c0063ebbd2baa3bd4949de081c6de85fe4f7c73349e6ffbe05784d03b", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:network_connector:reference", "xcsh-docs:resources:network_connector:examples", "xcsh-docs:resources:network_connector:import", "xcsh-docs:resources:network_connector:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:resources:network_connector:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/network_connector/index.md", "product": "distributed-cloud", "provider_name": "network_connector", "provider_schema_digest": "sha256:2a40709de0091cd22196a6618edea0e31e7e6ff85d9e40925107cbffcb22ebe5", "provider_type": "resources", "registry_anchor": "canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100", "registry_path": "docs/resources/network_connector.md", "relationships": [{"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:required", "target_id": "xcsh-docs:resources:virtual_network:fundamentals", "type": "advisory"}], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/network_connector/index.txt", "spec_pin_digest": "sha256:3be1648c5e375b30aafb719c70acb8365d8b7ebacfbddc9faa09e91b89e848fd", "summary": "Manages a Network Connector resource in F5 Distributed Cloud for network connector is created by users in system namespace. configuration.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.8", "schema_components": ["network_connectorCreateRequest"], "target_commit": "8e60c07c1d028cdbb49608452f0838e05471be84"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_network_connector

Breadcrumbs:

- xcsh_network_connector

Manages a Network Connector resource in F5 Distributed Cloud for network connector is created by
users in system namespace. configuration.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `virtual_network`.

- virtual_network: Network to connect

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NetworkConnector Resource Example
# Manages a Network Connector resource in F5 Distributed Cloud for network connector is created by users in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkConnector configuration
resource "xcsh_network_connector" "example" {
  name      = "example-network-connector"
  namespace = "staging"
}
```

## Root configuration

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/network_connector/lifecycle/timeouts/)
