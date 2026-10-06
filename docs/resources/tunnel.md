---
page_title: "xcsh_tunnel"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_tunnel."
---

# xcsh_tunnel

<a id="canonical-0203021123133200-3211200030212033-1303111112231210-3132112211210223-1120012223230031-3120003313300332-1222213113111002-3111033220000122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_tunnel

Manages tunnel in a given namespace. If one already exist it will give a error in F5 Distributed
Cloud.

<a id="canonical-2031221213301312-3111123132220303-1231032313022321-3101222331032013-3002111212320203-0210023101232223-3033220331201130-0201320210330002"></a>

### Prerequisites for `xcsh_tunnel`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1023112322003022-0230120000113303-0010321033232020-2233102102000202-2313010233021311-3103020033131112-1023310021021303-0300012022210013"></a>

### Minimal configuration for `xcsh_tunnel`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Tunnel Resource Example
# Manages tunnel in a given namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Tunnel configuration
resource "xcsh_tunnel" "example" {
  name      = "example-tunnel"
  namespace = "staging"
}
```

<a id="canonical-2321221031301210-3233303311321322-3222011030002230-3011002021023100-0300322320233213-1033333020010210-3002200321030201-3333121303333103"></a>

### Root configuration for `xcsh_tunnel`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0021221000031003-1113220222121033-1013311213001130-2213321330203030-3222023200300123-3333312022030221-0121311001030323-1303033023011310"></a>

### Explore this collection for `xcsh_tunnel`

- [Property reference](../guides/resources--tunnel--reference--group-001.md#canonical-0200201231331322-0303330032131211-1023201331203011-0231303320312231-3020103211200302-0321300301031220-0123132320301220-2202331313303213)
- [Examples](../guides/resources--tunnel--examples--group-001.md#canonical-2313003312031200-0221021012003310-1331120231133012-0101333030122330-2020203002311013-1011200200333123-1332322203132332-2103103231221131)
- [Import](../guides/resources--tunnel--lifecycle--group-001.md#canonical-1010131233201320-3203100203112233-3100302112033223-0210030201322003-2330312321130221-1233203330313332-3012322103303101-1330312132300010)
- [Timeouts](../guides/resources--tunnel--lifecycle--group-001.md#canonical-1230003013201103-3320232012332123-0002112120313013-2100112000101101-2332120012312230-1233133320102023-0200313112202101-3301003112311131)
