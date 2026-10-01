---
page_title: "xcsh_protected_domain landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protected_domain landing."
---

# xcsh_protected_domain landing

<a id="canonical-3133012330313310-3233013213001311-1212311200333023-2121112332020022-2232211202113213-1220123021122333-3111220011333221-3312301231323110"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3313232121112303-2023210332201000-2010200310313010-3203002223212032-0030230222121210-0200101012310100-3122221110210112-2223231121000300"></a>

## xcsh_protected_domain — xcsh_protected_domain / 022020222030 / 2

Breadcrumbs:

- xcsh_protected_domain

Manages Domain to protect in F5 Distributed Cloud.

<a id="canonical-3121303103021021-0120202132220333-0102101131232221-1333212212332110-2332233333012003-3213202330331223-2031231133101120-2321200113031213"></a>

## Prerequisites — xcsh_protected_domain / 022020222030 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2223312132100301-0202331032010002-3333312231013211-1120210012131311-2221231132112210-2230230213210231-2133100130320021-3100132131033132"></a>

## Minimal configuration — xcsh_protected_domain / 022020222030 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ProtectedDomain Resource Example
# Manages Domain to protect in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ProtectedDomain configuration
resource "xcsh_protected_domain" "example" {
  name      = "example-protected-domain"
  namespace = "staging"

  protected_domain = "example.com"
}
```

<a id="canonical-0200031002210030-1002310030201100-3123301231123321-0311202323210000-3330020101122121-3302322030132322-0202113310112020-1111020303233231"></a>

## Root configuration — xcsh_protected_domain / 022020222030 / 5

Required root properties: `name`, `namespace`, `protected_domain`. Full root flags and choices appear in the property reference.

<a id="canonical-1230000201121311-1111033000313203-1332222201110121-2322300122000223-1023332313032100-3031031331330310-2133032232033130-1020230103330202"></a>

## Next pages — xcsh_protected_domain / 022020222030 / 6

- [Property reference](../guides/resources--protected_domain--reference--group-001.md#canonical-0031202300233031-2101003333102132-1231000303030120-1020000033013310-3100323321110030-2313323021230021-3110000211021111-1302003213131230)
- [Examples](../guides/resources--protected_domain--examples--group-001.md#canonical-1031221233100201-2313210011123002-1020031220131123-3322003310200131-1011111321033102-1321001001033030-1200120331033032-2002012333003231)
- [Import](../guides/resources--protected_domain--lifecycle--group-001.md#canonical-2211232130123103-2210200122202320-2113010003020111-0211222301122123-1102323331123033-3301031230231020-3212210310321333-3033023211113211)
- [Timeouts](../guides/resources--protected_domain--lifecycle--group-001.md#canonical-1300010200211132-0321020211312200-0123330210300102-0202231301013231-1221303311213302-0130213322203001-3111233321230133-0312212321130212)
