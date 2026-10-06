---
page_title: "xcsh_virtual_network"
subcategory: "Networking"
description: "Manages virtual network in given namespace in F5 Distributed Cloud."
xcsh_docs: {"aliases": ["virtual network"], "body_bytes": 1678, "body_sha256": "sha256:6181d884253f5c8233f70b27e964494b3544c4f787a877f027d4c08ffc870002", "capabilities": ["networking"], "category": "networking", "child_ids": ["xcsh-docs:resources:virtual_network:reference", "xcsh-docs:resources:virtual_network:examples", "xcsh-docs:resources:virtual_network:import", "xcsh-docs:resources:virtual_network:timeouts"], "classification": {"rules_sha256": "sha256:e07d3e14cffec3e6fb45302e5bd7308760c60c7baecaa18c6d687c95d51ece70", "sources": ["receipt-pinned-upstream", "reviewed-rule"], "status": "resolved", "upstream_category_source": "receipt-pinned-resource"}, "collection_id": "xcsh-docs:resources:virtual_network:collection", "completeness": "complete", "id": "xcsh-docs:resources:virtual_network:fundamentals", "parent_id": "xcsh-docs:resources:xcsh:navigation", "path": "documentation/resources/virtual_network/index.md", "product": "distributed-cloud", "provider_name": "virtual_network", "provider_schema_digest": "sha256:76b040ce5603716b60411a0b7b17f23487536172558be7ac592beb4163ee8dcd", "provider_type": "resources", "registry_anchor": "canonical-0201332030311102-3030202012131131-3202102323100303-1011312013312200-3120301022031310-3013233001332030-0303333100331323-1232233111012001", "registry_path": "docs/resources/virtual_network.md", "relationships": [{"anchor": "", "enforcement": "upstream-advisory", "source": "receipt-pinned-dependency:optional", "target_id": "xcsh-docs:resources:network_connector:fundamentals", "type": "advisory"}], "retrieval_version": 1, "role": "fundamentals", "schema_path": [], "schema_version": 1, "sections": [], "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/resources/virtual_network/index.txt", "spec_pin_digest": "sha256:d5d38f86a968695cc56903b906faab6a444f49e9fcc196f42ae694b1a3960be0", "summary": "Manages virtual network in given namespace in F5 Distributed Cloud.", "tasks": ["configuration"], "unresolved_relationships": [], "upstream_identity": {"release_tag": "v12.0.0", "schema_components": ["virtual_networkCreateRequest"], "target_commit": "a9be0360815e3fd3fa08ded845e03c0bd4afd6ec"}}
---

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

# xcsh_virtual_network

Breadcrumbs:

- xcsh_virtual_network

Manages virtual network in given namespace in F5 Distributed Cloud.

## Prerequisites

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

Optional integrations: `network_connector`.

- network_connector: Connect to external networks

## Minimal configuration

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VirtualNetwork Resource Example
# Manages virtual network in given namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VirtualNetwork configuration
resource "xcsh_virtual_network" "example" {
  name      = "example-virtual-network"
  namespace = "system"
}
```

## Root configuration

Required root properties: `name`. Full root flags and choices appear in the property reference.

## Explore this collection

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/examples/)
- [Import](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/lifecycle/import/)
- [Timeouts](https://f5-sales-demo.github.io/terraform-provider-xcsh/resources/virtual_network/lifecycle/timeouts/)
