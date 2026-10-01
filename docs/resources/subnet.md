---
page_title: "xcsh_subnet landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_subnet landing."
---

# xcsh_subnet landing

<a id="canonical-0220113200023203-0330223333302123-0030302310210031-0031203130131223-0021111020100010-1302112031031001-1210101022122232-3300033203221001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3330021032132033-0130223033203213-0223010012222121-2210301201233011-3330001000321110-3211312203313133-0001211133012010-2001121013033323"></a>

## xcsh_subnet — xcsh_subnet / 121311322023 / 2

Breadcrumbs:

- xcsh_subnet

Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an
interface of a vm/pod. it is created in user or shared namespace. configuration.

<a id="canonical-3332023332332023-0021112001112031-3011100311320131-0003003020321200-3001222312132102-0100230131130330-2231311312113132-2030132330010133"></a>

## Prerequisites — xcsh_subnet / 121311322023 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0330303122013223-2101312011222133-0231032133022332-2032101321021310-1233013231001322-1103021231132110-1223023202033303-0022322102131023"></a>

## Minimal configuration — xcsh_subnet / 121311322023 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Subnet Resource Example
# Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an interface of a vm/pod.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Subnet configuration
resource "xcsh_subnet" "example" {
  name      = "example-subnet"
  namespace = "staging"
}
```

<a id="canonical-1101130022200023-1221303132102112-1303203013132130-1111230010011310-0131100202132310-0313101133311132-0201231122002322-0210203021021310"></a>

## Root configuration — xcsh_subnet / 121311322023 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1002002312302333-1020013321301031-1021101032210023-2002031100003320-2313232011012032-3210320212002020-2301233213201331-3301110122312011"></a>

## Next pages — xcsh_subnet / 121311322023 / 6

- [Property reference](../guides/resources--subnet--reference--group-001.md#canonical-0022331312302002-0010132211323310-2132031032011313-2230112310011301-1232302230021100-3102222330303123-1102202113312211-0131303030030331)
- [Examples](../guides/resources--subnet--examples--group-001.md#canonical-1213033300321211-3212302313120023-3023002000313103-0301300230331231-2033020103022203-2213121013303233-1020023113123100-3013010000012213)
- [Import](../guides/resources--subnet--lifecycle--group-001.md#canonical-2133302013023312-0313221302123301-1313011330200313-3112332131333123-1023233100311221-1011221302300011-1222312112300210-1212333213022020)
- [Timeouts](../guides/resources--subnet--lifecycle--group-001.md#canonical-0233312333123303-0022003333112323-3303010221022022-2300300001020113-1103321330131021-3311013103100300-3211011133123020-3331331311213323)
