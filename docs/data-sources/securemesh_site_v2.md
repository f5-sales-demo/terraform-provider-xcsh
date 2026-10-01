---
page_title: "xcsh_securemesh_site_v2 landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_securemesh_site_v2 landing."
---

# xcsh_securemesh_site_v2 landing

<a id="canonical-3033010310220323-3110311010223102-2010022012201213-2220103303200101-2000302233012310-3113200002022231-1123012300200020-0003333330101001"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3213203332132233-3312022013013023-1200121310020213-3110333231311223-3233313230010111-2020300203121022-1132330222020101-2022211211320103"></a>

## xcsh_securemesh_site_v2 — xcsh_securemesh_site_v2 / 232002200110 / 2

Breadcrumbs:

- xcsh_securemesh_site_v2

Manages a Securemesh Site V2 resource in F5 Distributed Cloud for deploying secure mesh edge sites
with security and networking controls.

<a id="canonical-0022310220023221-1201022012130220-3220320231001011-1322010332301230-3131113203311030-0020001321211120-0002221213230220-0130003200111111"></a>

## Prerequisites — xcsh_securemesh_site_v2 / 232002200110 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3022101202101120-3102021313030013-2211000202003313-1101123333313133-1023303000233011-2121121330333312-0023213313303010-1011220102222212"></a>

## Minimal configuration — xcsh_securemesh_site_v2 / 232002200110 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecuremeshSiteV2 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecuremeshSiteV2 by name
data "xcsh_securemesh_site_v2" "example" {
  name      = "example-securemesh-site-v2"
  namespace = "system"
}

output "securemesh_site_v2_id" {
  value = data.xcsh_securemesh_site_v2.example.id
}
```

<a id="canonical-3103020203120230-1210330002133121-1212131230331132-1001012333203230-1332211011221233-2203210222333321-2121202322213212-1212133230310013"></a>

## Root configuration — xcsh_securemesh_site_v2 / 232002200110 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-2223011103111123-3120202330321311-0030111221302201-0323202130223201-0101332333203012-0121310302011032-3321022113112033-0321020133233132"></a>

## Next pages — xcsh_securemesh_site_v2 / 232002200110 / 6

- [Property reference](../guides/data-sources--securemesh_site_v2--reference--group-001.md#canonical-1012223010032010-2203302200231202-3120001311232130-2102103100333033-3102211113031103-3332130330000122-0033200130112230-3010031233202002)
- [Examples](../guides/data-sources--securemesh_site_v2--examples--group-001.md#canonical-3131012331030122-2111321213311303-0210231321111013-3130330201023003-3230320203231321-3301312023320210-0310032311033233-3332203200033313)
