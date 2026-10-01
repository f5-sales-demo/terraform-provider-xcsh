---
page_title: "xcsh_bot_infrastructure landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_bot_infrastructure landing."
---

# xcsh_bot_infrastructure landing

<a id="canonical-2013132323231331-0303132312210331-1111131303032231-1230222213330000-3013332023002112-2232231312222310-2333212110130200-1220202223231210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201213300330213-0032111222322322-3013023312202112-2001102121201301-1330320221102101-0020330332233313-2322010202031300-3210302111213313"></a>

## xcsh_bot_infrastructure — xcsh_bot_infrastructure / 013031032303 / 2

Breadcrumbs:

- xcsh_bot_infrastructure

Manages Bot Infrastructure in F5 Distributed Cloud.

<a id="canonical-0212000311012323-2020223003331302-2203121203203321-2213203323010121-3020200130313313-2121331122001311-2322002233033221-0120113301232313"></a>

## Prerequisites — xcsh_bot_infrastructure / 013031032303 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0113100101112033-2232213103212100-3232111200231203-0031121211201110-2220212132001131-0213010300133230-2031112211211232-2221300210323213"></a>

## Minimal configuration — xcsh_bot_infrastructure / 013031032303 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# BotInfrastructure Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing BotInfrastructure by name
data "xcsh_bot_infrastructure" "example" {
  name      = "example-bot-infrastructure"
  namespace = "staging"
}

output "bot_infrastructure_id" {
  value = data.xcsh_bot_infrastructure.example.id
}
```

<a id="canonical-1122220011032001-1011013330230100-1120032332201331-3333023112221113-0320302330320300-0100211103233112-2230322220330321-3223112121223302"></a>

## Root configuration — xcsh_bot_infrastructure / 013031032303 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0122122302310112-2311120022021322-3122333211011331-0032311313310020-1023012330223021-0002011330223300-0220111131330023-0220230330232021"></a>

## Next pages — xcsh_bot_infrastructure / 013031032303 / 6

- [Property reference](../guides/data-sources--bot_infrastructure--reference--group-001.md#canonical-3203303203103230-0001213001331001-2201330032223212-1221312203121210-1202122312301302-0021121221223202-1033220220132003-1122303332333232)
- [Examples](../guides/data-sources--bot_infrastructure--examples--group-001.md#canonical-3320013121211030-2031323200002323-2132020301312100-1033303231202032-1322332132203020-3120200230130301-3023000303232123-0301210121302203)
