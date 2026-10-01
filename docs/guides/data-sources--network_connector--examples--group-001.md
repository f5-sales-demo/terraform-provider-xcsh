---
page_title: "xcsh_network_connector examples"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_network_connector examples."
---

# xcsh_network_connector examples

<a id="canonical-d56f582fe95f7ed54cc4eaaaaf7bc5e6867393740af0fd1509cb13e2a83d8b9a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-516a15eeac248bdfff7c3c7d17f5b40209bb352e41d480ec6a081de98197ef7e"></a>

## Examples — Examples / b7c46e201168 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- Examples

<a id="canonical-ad2cab86721009f779b882674538587cbbcbcef6c74420952fc1d2e648b727c9"></a>

## Complete configurations — Examples / b7c46e201168 / 3

- [Data source](data-sources--network_connector--examples--group-001.md#canonical-0b005c65357b08f3441e61db35ec47b4ee1118013d668e28bf4ffdec06da82e7): valid configuration.

<a id="canonical-afc06cb75ea1292601841889605011d4f9bd651076c816b014ffd11ab84df0dc"></a>

## Next pages — Examples / b7c46e201168 / 4

- [Data source](data-sources--network_connector--examples--group-001.md#canonical-0b005c65357b08f3441e61db35ec47b4ee1118013d668e28bf4ffdec06da82e7)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)

<a id="canonical-0b005c65357b08f3441e61db35ec47b4ee1118013d668e28bf4ffdec06da82e7"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b3c78f44dce3f9e649903bc2c6a5317215ef82ce760374138f6898051cdbaa71"></a>

## Data source — Data source / 9c26027d1cc0 / 2

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
- [Examples](data-sources--network_connector--examples--group-001.md#canonical-d56f582fe95f7ed54cc4eaaaaf7bc5e6867393740af0fd1509cb13e2a83d8b9a)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_connector/data-source.tf`; digest `sha256:842679078281a092ba5ea7fed221d6b2c452c242a5d7cc3a8d19f42931b95be3`.

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

<a id="canonical-ad47ad8abfabe57b6f38bc113a41f943375158780662cd155eb2a9a308ec788f"></a>

## Next pages — Data source / 9c26027d1cc0 / 3

- [Examples](data-sources--network_connector--examples--group-001.md#canonical-d56f582fe95f7ed54cc4eaaaaf7bc5e6867393740af0fd1509cb13e2a83d8b9a)
- [xcsh_network_connector](../data-sources/network_connector.md#canonical-d040f5dd4743f59e6899318b252de92fbb6486fabb9e35d963c5c2cbe239fdb1)
