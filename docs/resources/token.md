---
page_title: "xcsh_token landing"
subcategory: "Identity"
description: "Complete grouped canonical reference for xcsh_token landing."
---

# xcsh_token landing

<a id="canonical-1033011013002132-3200233320301222-0133100331111013-2201110031233021-3032032333322002-0033103111032332-0232210133220103-2232103023033020"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223333312222101-1301011022113200-3322010033321123-0022212212213011-0030112032301232-1122110013202020-1302221102122202-3021002230230331"></a>

## xcsh_token — xcsh_token / 331013331032 / 2

Breadcrumbs:

- xcsh_token

Manages new token. Token object is used to manage site admission. User must generate token before
provisioning and pass this token to site during it's registration in F5 Distributed Cloud.

<a id="canonical-0100012320131300-0331311120223232-1010133321300131-2213211232020230-0130210121203110-0200300033333213-1223333103011323-3031233331322301"></a>

## Prerequisites — xcsh_token / 331013331032 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Standard.

<a id="canonical-1013023310001300-0021133103031230-1200032200033103-2001123132123222-2303000021332000-1030132013300302-3202332232220210-3132331101313020"></a>

## Minimal configuration — xcsh_token / 331013331032 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Token Resource Example
# Manages new token.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Token configuration
resource "xcsh_token" "example" {
  name      = "example-token"
  namespace = "system"
  type      = 1
  site_name = "example-securemesh-site"
}
```

<a id="canonical-2222030323033012-1300112110110131-0203120303311200-0030021133232101-0110302131033102-1301031031300001-1113302313210310-0320100200323233"></a>

## Root configuration — xcsh_token / 331013331032 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-3330313011322311-3120210022131023-1210313010120210-2002023102330011-3221321120010201-2320301321210302-1033011233333312-0033301030031300"></a>

## Next pages — xcsh_token / 331013331032 / 6

- [Property reference](../guides/resources--token--reference--group-001.md#canonical-0011220303302322-2103033031012202-3113213332012123-0021203310111012-1122320333101333-3213031012021212-0131120002113213-3022102111320110)
- [Examples](../guides/resources--token--examples--group-001.md#canonical-1231001331010212-1313303001022020-1211202200302331-3202013321101030-2030333303033133-2200310331003311-3031023232120123-0203113200222102)
- [Import](../guides/resources--token--lifecycle--group-001.md#canonical-3131222232031130-1102111221001201-1310323230133202-1310100223300000-2132132202210232-0122021232113323-2003233122130213-1021331220333332)
- [Timeouts](../guides/resources--token--lifecycle--group-001.md#canonical-1130320132012022-0212013111310131-2031311231301312-0021130002321003-2133300131331222-1311133312221113-2332000230210302-3320300300313032)
