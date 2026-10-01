---
page_title: "xcsh_network_interface examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface examples."
---

# xcsh_network_interface examples

<a id="canonical-9fa7fa1fc3b9dbc53b75c7440e35e8a3e37ba004e065e2deff3541442c1ec5cc"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-b721a2682346866941518fb69e5eb21f37722e5dbc085a006a66da862586ba72"></a>

## Examples — Examples / 7832e6a49d15 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- Examples

<a id="canonical-346d9c285419015e7e737cc3dffd054c5d23298bacb655c94cb18edd4d2ac08d"></a>

## Complete configurations — Examples / 7832e6a49d15 / 3

- [Data source](data-sources--network_interface--examples--group-001.md#canonical-4cfc2bec4ccb1c343cddbb1507d832ce2c1e5028ca4a83498704bee9208a164b): valid configuration.

<a id="canonical-5beafeefff9cfb3d326cd4bee8226651bdebec644b6bd5e4f42baa173064436e"></a>

## Next pages — Examples / 7832e6a49d15 / 4

- [Data source](data-sources--network_interface--examples--group-001.md#canonical-4cfc2bec4ccb1c343cddbb1507d832ce2c1e5028ca4a83498704bee9208a164b)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)

<a id="canonical-4cfc2bec4ccb1c343cddbb1507d832ce2c1e5028ca4a83498704bee9208a164b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-86e1e94a79017275454fd4ea5f1ed631b7a844169022683650657b7d43e78356"></a>

## Data source — Data source / 726a3b2b8c61 / 2

Breadcrumbs:

- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
- [Examples](data-sources--network_interface--examples--group-001.md#canonical-9fa7fa1fc3b9dbc53b75c7440e35e8a3e37ba004e065e2deff3541442c1ec5cc)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_interface/data-source.tf`; digest `sha256:2b23ed57ede50484e6ffd1c4ce6fdf9ba70d1bbaddf2b7e8ceb68afb488d14e2`.

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

<a id="canonical-24520d8a60085c339a37c0ff671b65ae5277eb0362896e99ad1c5b30824d7285"></a>

## Next pages — Data source / 726a3b2b8c61 / 3

- [Examples](data-sources--network_interface--examples--group-001.md#canonical-9fa7fa1fc3b9dbc53b75c7440e35e8a3e37ba004e065e2deff3541442c1ec5cc)
- [xcsh_network_interface](../data-sources/network_interface.md#canonical-2b84ad157d6a5aa91dc5f09196ffd113bb3ae8054c873880ffcc2055550d089e)
