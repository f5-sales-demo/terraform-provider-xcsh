---
page_title: "xcsh_network_connector"
subcategory: "Networking"
description: "xcsh_network_connector for xcsh_network_connector."
xcsh_docs: {"aliases": [], "body_bytes": 1551, "body_sha256": "sha256:fcc2617da9924d93d9941453e309dfe264258690c1c58f0c738b3c722c897da4", "child_ids": ["xcsh-docs:data-sources:network_connector:reference", "xcsh-docs:data-sources:network_connector:examples"], "collection_id": "xcsh-docs:data-sources:network_connector:collection", "completeness": "complete", "id": "xcsh-docs:data-sources:network_connector:fundamentals", "parent_id": null, "path": "documentation/data-sources/network_connector/index.md", "provider_name": "network_connector", "provider_schema_digest": "sha256:63e4fbb3e2007c78e36dc243aa3840076ce8a32eff8e7120cacece30bcdd2cd6", "provider_type": "data-sources", "role": "fundamentals", "schema_path": [], "schema_version": 1, "source_url": "https://f5-sales-demo.github.io/terraform-provider-xcsh/_data/pages/data-sources/network_connector/index.txt", "spec_pin_digest": "sha256:5236cddd67bd603b0ce9bb70d82e356a161a3be971d90fc60eb8119dfb833774", "summary": "xcsh_network_connector for xcsh_network_connector.", "upstream_identity": {"release_tag": "v9.0.0", "schema_components": ["network_connectorCreateRequest"], "target_commit": "f95183046197e6547a3ea123e324f53126eae440"}}
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

## Next pages

- [Property reference](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/properties/)
- [Examples](https://f5-sales-demo.github.io/terraform-provider-xcsh/data-sources/network_connector/examples/)
