---
page_title: "xcsh_network_connector landing"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_network_connector landing."
---

# xcsh_network_connector landing

<a id="canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-5dd81a10ab30d81c3f8fc6860fdba41d96392dc505ce4a4c992200f18226d13d"></a>

## xcsh_network_connector — xcsh_network_connector / 73c0179c57c4 / 2

Breadcrumbs:

- xcsh_network_connector

Manages a Network Connector resource in F5 Distributed Cloud for network connector is created by
users in system namespace. configuration.

<a id="canonical-a32d51bf19a1a2413bf0141d4f25756bb0f7b00e45e56aaea923e1ff0e763c7b"></a>

## Prerequisites — xcsh_network_connector / 73c0179c57c4 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `virtual_network`.

- virtual_network: Network to connect

<a id="canonical-6d12c00ba578c8f390ee2e0347669f46550a1937f32c66cf30ff889102b749c8"></a>

## Minimal configuration — xcsh_network_connector / 73c0179c57c4 / 4

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

<a id="canonical-af214b6a5a69f0ee5f515cf1819da49f7924fe45f9c666a6c0ce0f8218fb3207"></a>

## Root configuration — xcsh_network_connector / 73c0179c57c4 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2fa4884f9e18d2d1feaa13838b1b9fa84cb40eacdb60b5b857c280c09d9fa440"></a>

## Next pages — xcsh_network_connector / 73c0179c57c4 / 6

- [Property reference](../guides/data-sources--network_connector--reference--group-001.md#canonical-c06a37a78cc7aecb727957f232ae170c33a6bff489e5e7a0d08405ec75af5138)
- [Examples](../guides/data-sources--network_connector--examples--group-001.md#canonical-d56f582fe95f7ed54cc4eaaaaf7bc5e6867393740af0fd1509cb13e2a83d8b9a)
