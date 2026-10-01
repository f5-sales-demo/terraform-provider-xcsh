---
page_title: "xcsh_external_connector landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_external_connector landing."
---

# xcsh_external_connector landing

<a id="canonical-3122312021220130-3300301303031130-2302131220123022-3100221331011030-0221003203020132-3312113011210302-1031003011331110-0113230133233222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3331211320002300-1233210313003033-3002032203310110-3320122212023003-3301120020221023-0110032330323120-2201332010233311-1303132311022331"></a>

## xcsh_external_connector — xcsh_external_connector / 131022323002 / 2

Breadcrumbs:

- xcsh_external_connector

Manages a External Connector resource in F5 Distributed Cloud for external\_connector configuration
specification. configuration.

<a id="canonical-3022111333121103-0323223003010332-3002030301332122-3223210120321202-3201212122012210-2111312101022001-2032122011102121-3220310211321221"></a>

## Prerequisites — xcsh_external_connector / 131022323002 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0222111213231300-2201200130132331-1130321120333033-0131011330022013-2022212220210210-1310023312202322-0123200030212121-1000211232000120"></a>

## Minimal configuration — xcsh_external_connector / 131022323002 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ExternalConnector Resource Example
# Manages a External Connector resource in F5 Distributed Cloud for external_connector configuration specification.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic ExternalConnector configuration
resource "xcsh_external_connector" "example" {
  name      = "example-external-connector"
  namespace = "staging"
}
```

<a id="canonical-3030123011011322-1030312020020010-3000010021322303-1100333200130232-1020023033120121-1313313132323003-0222030021231310-3102013113222123"></a>

## Root configuration — xcsh_external_connector / 131022323002 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1003310030130333-0122310201312020-2233300313121200-1033111131132113-1133312313001212-2322331332320131-2203032021331100-2201012322021300"></a>

## Next pages — xcsh_external_connector / 131022323002 / 6

- [Property reference](../guides/resources--external_connector--reference--group-001.md#canonical-3213311113001013-2210200203110132-0213333230312322-2323001310333022-3213131211021012-0222230300311031-1133032330331310-3032202302013232)
- [Examples](../guides/resources--external_connector--examples--group-001.md#canonical-1023020021113213-0322331233001322-0012100022302130-1013123302233230-0232032201001213-3203001323100231-3330210233012201-1001111112222333)
- [Import](../guides/resources--external_connector--lifecycle--group-001.md#canonical-3101203302032013-3223302322323222-2012123333131213-0111113112232132-1221222111101032-1332122322200000-2330022313230303-1321030113210320)
- [Timeouts](../guides/resources--external_connector--lifecycle--group-001.md#canonical-0311311020002201-0111320220201230-1221201123312323-3002302331212201-0120213000231022-0023222201210022-2121313201110202-2131311122230310)
