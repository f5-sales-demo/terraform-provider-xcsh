---
page_title: "xcsh_address_allocator landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_address_allocator landing."
---

# xcsh_address_allocator landing

<a id="canonical-2221003231222131-0012132022313001-1331101002320300-2210032101320132-3133132121300222-1310012132000200-0313122021123003-3221012203222322"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0313122112313200-0130031230012202-2010121222033221-0300201222121333-1233123131101331-2300111033202013-3201302322011123-2201110000003300"></a>

## xcsh_address_allocator — xcsh_address_allocator / 203330311201 / 2

Breadcrumbs:

- xcsh_address_allocator

Manages Address Allocator will create an address allocator object in 'system' namespace of the user
in F5 Distributed Cloud.

<a id="canonical-3130230212030122-3331201210031303-3333000320323121-1013321311023002-2331022033222320-2331222113001320-0301313333213323-2303022023322011"></a>

## Prerequisites — xcsh_address_allocator / 203330311201 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2010321210310101-1123201003030300-3022022311313210-1302212023031211-1312223312031221-0220333223201013-1023230130321220-3003223202120211"></a>

## Minimal configuration — xcsh_address_allocator / 203330311201 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AddressAllocator Resource Example
# Manages Address Allocator will create an address allocator object in 'system' namespace of the user in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AddressAllocator configuration
resource "xcsh_address_allocator" "example" {
  name      = "example-address-allocator"
  namespace = "staging"

  address_pool = ["example-value"]
}
```

<a id="canonical-1023133312211201-0211231220130102-3012131230232112-0213210000210221-1011303031121110-1321111211111330-3131300131213203-0330032122031110"></a>

## Root configuration — xcsh_address_allocator / 203330311201 / 5

Required root properties: `address_pool`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2121323312332033-1311033000332010-1021112120222023-3203013132021233-2322212211322312-2213010222213113-0010303330121023-0223223202211322"></a>

## Next pages — xcsh_address_allocator / 203330311201 / 6

- [Property reference](../guides/resources--address_allocator--reference--group-001.md#canonical-1230110323000212-2131122201211332-1030122332223221-1013221331022031-2010023222021300-0120130133003011-1001323132213111-1010131113330213)
- [Examples](../guides/resources--address_allocator--examples--group-001.md#canonical-2122233302132220-3232230300022102-0223303330221223-1202332103132120-1130001332322133-0300033120330001-1310232003003020-1303302301211102)
- [Import](../guides/resources--address_allocator--lifecycle--group-001.md#canonical-1332013322130313-3321202323023222-0103320210300230-0132011011032201-2120120331120201-0031203123201023-2232033111102320-2020121130133332)
- [Timeouts](../guides/resources--address_allocator--lifecycle--group-001.md#canonical-1200322321130011-3000003213030032-2201323033322102-1203011211332223-0221330202331312-2203023211322311-3311302222201121-1222331311121202)
