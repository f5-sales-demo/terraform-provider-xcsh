---
page_title: "xcsh_network_connector examples"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_network_connector examples."
---

# xcsh_network_connector examples

<a id="canonical-3122113101222231-3312000230223303-1133221210021201-1323201211021320-0131303133312321-3010230001001110-3202321232320131-0201123233233333"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- Examples

<a id="canonical-1330013101022221-3213232003102002-2102201130221103-3103311013023100-0300300000121030-1220230303101323-3321011233331302-3302300102121132"></a>

### Complete configurations for `xcsh_network_connector`

- [Resource](resources--network_connector--examples--group-001.md#canonical-1030210113010330-3111121101310302-1130211303231232-3030211332110311-0232111301310120-1033231001203113-2133001020133213-3320010021030021): valid configuration.

<a id="canonical-1030210113010330-3111121101310302-1130211303231232-3030211332110311-0232111301310120-1033231001203113-2133001020133213-3320010021030021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_network_connector](../resources/network_connector.md#canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100)
- [Examples](resources--network_connector--examples--group-001.md#canonical-3122113101222231-3312000230223303-1133221210021201-1323201211021320-0131303133312321-3010230001001110-3202321232320131-0201123233233333)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_network_connector/resource.tf`; digest `sha256:79950822068b0bbf291a51dad0cff745ae6eacc14d286990d529310031a03379`.

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
