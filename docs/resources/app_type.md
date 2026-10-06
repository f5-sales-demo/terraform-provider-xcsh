---
page_title: "xcsh_app_type"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_type."
---

# xcsh_app_type

<a id="canonical-2220100013322301-3233230230220131-2230301203210321-0110213003232211-2331220333323130-0302001023303001-2320103333130202-1122210103122102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_app_type

Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed
Cloud.

<a id="canonical-1022110003213123-2323220310233101-2113122021311122-3101231001310322-0332023320331222-2031103122121021-3133033320313000-1121220000120232"></a>

### Prerequisites for `xcsh_app_type`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0021022001213113-3023001001112001-0032012233113002-0303120120132211-0102103330331222-1001002122310021-0003010111323113-2331031011033013"></a>

### Minimal configuration for `xcsh_app_type`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppType Resource Example
# Manages App type will create the configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppType configuration
resource "xcsh_app_type" "example" {
  name      = "example-app-type"
  namespace = "staging"
}
```

<a id="canonical-3002001311123323-1213230202331213-2211202132322302-1020020113231132-3023033223130310-0212122202111030-0233331010223203-2020022120232302"></a>

### Root configuration for `xcsh_app_type`

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3323123013001132-1133122220331111-1321103023010222-1323110310320200-3311003003311011-3121001001323320-3333311333333220-2110213321032122"></a>

### Explore this collection for `xcsh_app_type`

- [Property reference](../guides/resources--app_type--reference--group-001.md#canonical-2010013132232302-1121312303131032-1120000202313211-0303313333201213-1313211102301122-0021132003012310-2330313333002232-1320001121301320)
- [Examples](../guides/resources--app_type--examples--group-001.md#canonical-3313232102212003-2110322003232032-1022002110331123-2100110311031302-2201131131312212-0233003322122301-1332112323100030-3033122312231231)
- [Import](../guides/resources--app_type--lifecycle--group-001.md#canonical-2223021120323011-2221033232013133-1301012102321203-0223132113303030-1023102303002333-2323020123010323-2130001213202332-1202220322220011)
- [Timeouts](../guides/resources--app_type--lifecycle--group-001.md#canonical-1001331200113010-0301000303301222-3132012201020122-2310301301132012-0113330133132231-3230011201021022-1031221332211012-2320322020213221)
