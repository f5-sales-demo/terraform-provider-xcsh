---
page_title: "xcsh_site_signatures_update"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_site_signatures_update."
---

# xcsh_site_signatures_update

<a id="canonical-2121133033301232-0011100131212313-0013203021322030-2021023332100333-1321010310023012-3332003100213031-1333010220333010-1320230310323133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_site_signatures_update

Requests an update of site signatures.

<a id="canonical-1003302010010213-1013321310211332-3331033200113330-2020110203130320-2232132222021132-2000232203200230-2133122002021232-3023023010312331"></a>

### Prerequisites for `xcsh_site_signatures_update`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3012232011132101-2032132333030323-2203013221201220-2130201122111233-3013331003013122-0022101120010203-3023130132310132-3023112131313020"></a>

### Minimal configuration for `xcsh_site_signatures_update`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SiteSignaturesUpdate Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_site_signatures_update" "example" {
  config {
    namespace = "example-value"
  }
}
```

<a id="canonical-2210323133302233-1320000111031023-2111021103032200-0213110001230031-0230022121231030-2310133312233313-1310320203001321-3201312220001033"></a>

### Root configuration for `xcsh_site_signatures_update`

Required root properties: `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1303011103203113-1111301301002211-3032010311012002-0110130121110220-2322210022310103-2330020030223313-3102012201311320-1302212231332112"></a>

### Explore this collection for `xcsh_site_signatures_update`

- [Property reference](../guides/actions--site_signatures_update--reference--group-001.md#canonical-2103311202213331-3322033031022121-3121001201130323-0111123021202301-2313002222033000-2333220012121133-2230132001320330-3233213301303202)
- [Examples](../guides/actions--site_signatures_update--examples--group-001.md#canonical-1133211000213003-1211333131332233-3010121333101232-3320302222201103-1303011220210213-1330031203320202-3303112331111132-1010002322013121)
- [Lifecycle](../guides/actions--site_signatures_update--lifecycle--group-001.md#canonical-0130221213333233-0201210102022213-2002000020332232-0233003221023000-0021230302030211-0332232321113122-3112012120210323-3302021201323030)
