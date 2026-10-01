---
page_title: "xcsh_voltstack_site landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_voltstack_site landing."
---

# xcsh_voltstack_site landing

<a id="canonical-2321122020020113-3300101223021223-1330033011121200-0102203210311033-2301300201121233-3102120211011323-1113113111112000-2313011002230113"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3223001322322300-2032212221112202-0132332333223130-1103222213003232-2031011321321100-3021303213300102-0303203130103121-0230110022101022"></a>

## xcsh_voltstack_site — xcsh_voltstack_site / 031101031002 / 2

Breadcrumbs:

- xcsh_voltstack_site

Manages a Voltstack Site resource in F5 Distributed Cloud for deploying App Stack edge computing
sites.

<a id="canonical-1123012233332122-2031312303000231-2013332312323301-1323120330030202-2333333222300212-2221013001122100-1023222320311231-0132103213332203"></a>

## Prerequisites — xcsh_voltstack_site / 031101031002 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3112000030030001-3221132213120211-0320330111133121-1103022122101331-3012302331112322-3230220123033110-0323232110132011-2110211031001303"></a>

## Minimal configuration — xcsh_voltstack_site / 031101031002 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# VoltstackSite Resource Example
# Manages a Voltstack Site resource in F5 Distributed Cloud for deploying App Stack edge computing sites.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic VoltstackSite configuration
resource "xcsh_voltstack_site" "example" {
  name      = "example-voltstack-site"
  namespace = "staging"

  volterra_certified_hw = "example-value"
}
```

<a id="canonical-1011332112321102-3110311110030033-3133113120122023-3101231102223103-3311223200122101-0302332232030211-1230330320232231-3132101112031101"></a>

## Root configuration — xcsh_voltstack_site / 031101031002 / 5

Required root properties: `name`, `namespace`, `volterra_certified_hw`. Full root flags and choices appear in the property reference.

<a id="canonical-1010130333200130-0332100133033013-1121331233033321-3101101120122312-3222110212223130-1302331033311201-1103333101113113-1211331002221331"></a>

## Next pages — xcsh_voltstack_site / 031101031002 / 6

- [Property reference](../guides/resources--voltstack_site--reference--group-001.md#canonical-2303313130303000-1321333232312030-2110222133201202-1121003303213312-0221012033213113-0321333012001121-3323103113212202-1012112312220111)
- [Examples](../guides/resources--voltstack_site--examples--group-001.md#canonical-2000120202312222-2211323300223201-1032311002123102-1200033311312000-1331122302231313-3032010120220020-0133200201003101-1110323012110322)
- [Import](../guides/resources--voltstack_site--lifecycle--group-001.md#canonical-0001301210130033-2313220120202202-0122322100020213-3322311011313122-1331323103212023-0130101233313000-1130211013231321-2032020313003323)
- [Timeouts](../guides/resources--voltstack_site--lifecycle--group-001.md#canonical-3001231222312003-2312011002231211-2132333103221211-0000132323313321-3332000300200311-1200011220111021-1203123203102220-1033213231320111)
