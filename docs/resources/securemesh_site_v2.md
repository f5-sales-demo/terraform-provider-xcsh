---
page_title: "xcsh_securemesh_site_v2 landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 landing."
---

# xcsh_securemesh_site_v2 landing

<a id="canonical-3000200222223112-2121123210332232-3331131023113311-0101321102313201-1230101233301033-1102230312220221-1020100000200232-0100112322233120"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2001210223032030-0300133113222102-0211310012123220-2011032302310031-3213321221002303-3010001301230111-2132202303110322-2313022330012133"></a>

## xcsh_securemesh_site_v2 — xcsh_securemesh_site_v2 / 302012320203 / 2

Breadcrumbs:

- xcsh_securemesh_site_v2

Manages a Securemesh Site V2 resource in F5 Distributed Cloud for deploying secure mesh edge sites
with security and networking controls.

<a id="canonical-1101002023300233-3002213112211233-1310111110000321-3100331321103032-2111233012330310-0230031312210301-3320330233122203-1320333300330020"></a>

## Prerequisites — xcsh_securemesh_site_v2 / 302012320203 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3120331221003103-0123020213300201-2320302312030311-2330023300300323-0113001332210120-2100202121221333-2111130230113322-1311021021332102"></a>

## Minimal configuration — xcsh_securemesh_site_v2 / 302012320203 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecuremeshSiteV2 Resource Example
# Manages a Securemesh Site V2 resource in F5 Distributed Cloud for deploying secure mesh edge sites with security and networking controls.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic SecuremeshSiteV2 configuration
resource "xcsh_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2"
  namespace = "system"
}
```

<a id="canonical-1232030232312330-2310310333112011-0103112013332200-0220033201331232-3113223231021221-0022221331012311-2301110213230013-2200200111203003"></a>

## Root configuration — xcsh_securemesh_site_v2 / 302012320203 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1133013331132113-1123323302003323-1333001202103022-3313212323301223-1031131321101002-0020000200332310-3300231032311111-2112102021110011"></a>

## Next pages — xcsh_securemesh_site_v2 / 302012320203 / 6

- [Property reference](../guides/resources--securemesh_site_v2--reference--group-001.md#canonical-2301010111101201-2220210123001232-1330030231123103-3121113232112322-0131203200320021-3021122122123123-2102102121133210-3220303330323302)
- [Examples](../guides/resources--securemesh_site_v2--examples--group-001.md#canonical-1303300012312320-3201110132301331-1220203112220033-2121022103321322-3322310212312003-1133112300000111-2012112001331231-3033021103030133)
- [Import](../guides/resources--securemesh_site_v2--lifecycle--group-001.md#canonical-2012320112212100-1230102332111212-0301020010123231-1001330200011202-1033220320310110-0320120311130230-3303103302101302-0301220212312231)
- [Timeouts](../guides/resources--securemesh_site_v2--lifecycle--group-001.md#canonical-0301003223322322-1310123212001232-2312121203032200-2001301230000000-1232213110102020-2230223321333303-2312103003021013-3210220232322313)
