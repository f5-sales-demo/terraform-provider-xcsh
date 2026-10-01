---
page_title: "xcsh_nfv_service landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nfv_service landing."
---

# xcsh_nfv_service landing

<a id="canonical-2212320103310132-2221300222221032-1233031123200120-3022110322311231-3301003023302233-3133232120201022-0322021211220213-1332311000003300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0023030310312301-1010123012033032-0201000022221330-1132212323201102-1002012231233111-1111020111201320-0011100213012031-3020323130221311"></a>

## xcsh_nfv_service — xcsh_nfv_service / 200112130300 / 2

Breadcrumbs:

- xcsh_nfv_service

Manages new NFV service with configured parameters in F5 Distributed Cloud.

<a id="canonical-0223003020021100-2333003001110011-1323033111021000-3130233303010301-2213120303333121-3322221331122113-1231033211001123-1020222213231210"></a>

## Prerequisites — xcsh_nfv_service / 200112130300 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3020113102121110-0331133020322333-2110331110111100-1110032232122021-1123032130300332-0012121233122302-1311033133311131-0021203301133232"></a>

## Minimal configuration — xcsh_nfv_service / 200112130300 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# NfvService Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NfvService by name
data "xcsh_nfv_service" "example" {
  name      = "example-nfv-service"
  namespace = "staging"
}

output "nfv_service_id" {
  value = data.xcsh_nfv_service.example.id
}
```

<a id="canonical-2301013311211011-0013322303003312-1023220031311310-3132123100313310-0312122210033310-1132111331133212-2322131112002110-1220331032333303"></a>

## Root configuration — xcsh_nfv_service / 200112130300 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2031032233011032-1002230233202323-0013300200323030-1102110002212031-0031333233012302-1333221111222230-0032103130321322-3200111023200203"></a>

## Next pages — xcsh_nfv_service / 200112130300 / 6

- [Property reference](../guides/data-sources--nfv_service--reference--group-001.md#canonical-2313010333323131-3010030231311121-2301323120300012-0122321020201331-3020223211220022-0100121100023011-2323321311300323-0122100132003020)
- [Examples](../guides/data-sources--nfv_service--examples--group-001.md#canonical-0113201012132103-2303210110031133-1100201210101112-3112311201021033-0132302110030230-2000330102332211-0130111220020003-0013200013312112)
