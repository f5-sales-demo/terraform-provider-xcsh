---
page_title: "xcsh_authorization_server landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authorization_server landing."
---

# xcsh_authorization_server landing

<a id="canonical-3123011103301332-2020213322103110-2220312130201323-0120021121102210-3010302120211132-3100321223002321-1020323131322210-3122032302302331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2300101310032120-0101130202012323-3120112313001312-1311231100131120-3010211031013030-0022320211131222-1003123330200302-3120211320223021"></a>

## xcsh_authorization_server — xcsh_authorization_server / 321213003321 / 2

Breadcrumbs:

- xcsh_authorization_server

Manages authorization\_server creates a new object in the storage backend for metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-3303011232122310-0020131131233123-3123100330112320-3323330102122310-2202021223211020-2113210121231221-0003013231213212-1212003222213211"></a>

## Prerequisites — xcsh_authorization_server / 321213003321 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1011133013010210-2222231231011031-1003001133112230-0320303011321331-1310032113132133-3032201113012323-0221223200123103-3332010223033211"></a>

## Minimal configuration — xcsh_authorization_server / 321213003321 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AuthorizationServer Resource Example
# Manages authorization_server creates a new object in the storage backend for metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AuthorizationServer configuration
resource "xcsh_authorization_server" "example" {
  name      = "example-authorization-server"
  namespace = "staging"

  jwks_uri = "example-value"
}
```

<a id="canonical-0103033033003001-2120033133020030-0201120133120333-3113012110130123-0210201103012111-3030113221332101-1310312103301000-0200131013132023"></a>

## Root configuration — xcsh_authorization_server / 321213003321 / 5

Required root properties: `jwks_uri`, `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0212311113031022-0301121330013103-1323023223303020-0002110233333230-2120112012111220-1313210211001232-2032322012323010-1231300012311211"></a>

## Next pages — xcsh_authorization_server / 321213003321 / 6

- [Property reference](../guides/resources--authorization_server--reference--group-001.md#canonical-0211111022313012-0333312311133222-3201102131123002-1000213313220112-0313210110001033-0311120333133101-3203103013101012-3320032021010331)
- [Examples](../guides/resources--authorization_server--examples--group-001.md#canonical-0123232121303011-2013012103001211-1021110113133132-3012133011121032-2011112332012332-0013012201121022-2123332313132320-2212100302331330)
- [Import](../guides/resources--authorization_server--lifecycle--group-001.md#canonical-3311130103213132-0210202300312100-3032323010000201-2311222100103221-0100311121130303-1011113132100033-2320012131020013-1013321200310133)
- [Timeouts](../guides/resources--authorization_server--lifecycle--group-001.md#canonical-1003013211013031-0300003020203232-0023330321023111-0133132212301212-3221102101310230-0223210311133310-2322020230312202-3031122230011231)
