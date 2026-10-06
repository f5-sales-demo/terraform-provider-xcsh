---
page_title: "xcsh_network_interface examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface examples."
---

# xcsh_network_interface examples

<a id="canonical-3132230203221132-3031330331333000-0222013331322211-1113211033302102-3103023331112012-3211203130301301-3301101022012211-2303213321321031"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- Examples

<a id="canonical-0211232211033233-2021122021200133-0103333110322210-1020013110103332-3013032313300223-0133233002302201-0231200003320300-0300113312300000"></a>

### Complete configurations for `xcsh_network_interface`

- [Resource](resources--network_interface--examples--group-001.md#canonical-2003300332031011-0310201123023101-2001130020212130-0133100130023231-0122000312012030-0213003300322000-3313221322331232-0303200100032021): valid configuration.

<a id="canonical-2003300332031011-0310201123023101-2001130020212130-0133100130023231-0122000312012030-0213003300322000-3313221322331232-0303200100032021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_network_interface](../resources/network_interface.md#canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310)
- [Examples](resources--network_interface--examples--group-001.md#canonical-3132230203221132-3031330331333000-0222013331322211-1113211033302102-3103023331112012-3211203130301301-3301101022012211-2303213321321031)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_interface/resource.tf`; digest `sha256:7844384dc87eb9b41829fe1bcb33636a61a1469e9b6c183c89c2bb605fd224c5`.

```terraform
# NetworkInterface Resource Example
# Manages a Network Interface resource in F5 Distributed Cloud for network interface represents configuration of a network device.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic NetworkInterface configuration
resource "xcsh_network_interface" "example" {
  name      = "example-network-interface"
  namespace = "staging"
}
```
