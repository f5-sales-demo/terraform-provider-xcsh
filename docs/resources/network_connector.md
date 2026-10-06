---
page_title: "xcsh_network_connector"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_network_connector."
---

# xcsh_network_connector

<a id="canonical-0030231112200120-2310101313112331-0112221320311330-1010030020110100-0320030301321031-0201122000110313-2302331223321203-3130322032000100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_connector

Manages a Network Connector resource in F5 Distributed Cloud for network connector is created by
users in system namespace. configuration.

<a id="canonical-3020130232122003-0202303201232011-2132013322303323-1002121230023001-1301322123203030-0111310322303031-2220132010233110-2333332100321031"></a>

### Prerequisites for `xcsh_network_connector`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

Required dependencies: `virtual_network`.

- virtual_network: Network to connect

<a id="canonical-3110331301210311-0101112032210112-3221001130233112-1121301302321002-2030323103031131-1320100231300102-2322033101111000-2300023211320013"></a>

### Minimal configuration for `xcsh_network_connector`

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

<a id="canonical-2311313310000332-0123000220000131-3103301020011322-1331103002123030-1002022221323201-2220030213220320-1230003103330221-3131222311211202"></a>

### Root configuration for `xcsh_network_connector`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3123233202100211-2312120132321321-2132303012113321-2130021210121013-0011121322320232-1230313301320332-1203112321001202-1213121031103302"></a>

### Explore this collection for `xcsh_network_connector`

- [Property reference](../guides/resources--network_connector--reference--group-001.md#canonical-3230201320313300-3100103222213301-3222301220332000-2102123311113022-2103213133210102-1103110232333032-2123011010231323-2310223103003331)
- [Examples](../guides/resources--network_connector--examples--group-001.md#canonical-3122113101222231-3312000230223303-1133221210021201-1323201211021320-0131303133312321-3010230001001110-3202321232320131-0201123233233333)
- [Import](../guides/resources--network_connector--lifecycle--group-001.md#canonical-2121021212010001-3132121011031220-0300312003111232-0122100212021110-1132301221013132-1122202020030030-1100001331331123-1313000032130213)
- [Timeouts](../guides/resources--network_connector--lifecycle--group-001.md#canonical-3012322222301233-3323333130133033-3203123210013320-3110211003112222-2110122023112211-1131010121210100-3300002133130022-2001031132210112)
