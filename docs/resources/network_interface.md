---
page_title: "xcsh_network_interface"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_interface."
---

# xcsh_network_interface

<a id="canonical-3303002200232320-3331233131000231-2111111223023013-2313310022210202-2333223030322133-2311103000000220-1221101012132133-2022200333323310"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_network_interface

Manages a Network Interface resource in F5 Distributed Cloud for network interface represents
configuration of a network device. it is created by users in system namespace. configuration.

<a id="canonical-0323222203220321-1100111301200103-0001013200323302-1030023330203022-0101032001112302-1331210013032123-0120223323022022-0211001223113232"></a>

### Prerequisites for `xcsh_network_interface`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2023121031220312-1101210112332110-1302333022221300-3003232303131210-0302200330123123-3120330232111321-1302230300200202-1323101203300111"></a>

### Minimal configuration for `xcsh_network_interface`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-0220332000322311-3233101220331030-3302123030212222-2311223031320330-3012011121010122-3303211221130302-1232300001132002-3301322232200302"></a>

### Root configuration for `xcsh_network_interface`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2330311100333332-3330011110032121-1332222320030321-3021200033120120-3231031010032232-0303022201222023-2300032023303320-1030132333132320"></a>

### Explore this collection for `xcsh_network_interface`

- [Property reference](../guides/resources--network_interface--reference--group-001.md#canonical-2322022123030133-3000312103003333-0310202220121330-0310123030213020-3022333010321131-3222303201013302-0200010223022132-2000311123301212)
- [Examples](../guides/resources--network_interface--examples--group-001.md#canonical-3132230203221132-3031330331333000-0222013331322211-1113211033302102-3103023331112012-3211203130301301-3301101022012211-2303213321321031)
- [Import](../guides/resources--network_interface--lifecycle--group-001.md#canonical-3202322311332023-0100223231033130-2213001230333101-2302222313120021-0323233203000131-3310112033131001-0021311030221121-1111303211301112)
- [Timeouts](../guides/resources--network_interface--lifecycle--group-001.md#canonical-2033302012022012-1312232211112223-1021312110301132-0310112010313230-3023131210132302-3311122130200000-0213222102103133-0320121212031221)
