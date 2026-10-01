---
page_title: "xcsh_authentication landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_authentication landing."
---

# xcsh_authentication landing

<a id="canonical-1001133113031021-0312110233112201-1313013233101233-0220003302033333-3111302101121213-0313132200203312-1002120232112230-2123221000130210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2332301102210112-2321110232311222-1100323102121123-2212001121131123-0002333301122013-1012332103303001-3032300132313102-1030101123331311"></a>

## xcsh_authentication — xcsh_authentication / 232331300223 / 2

Breadcrumbs:

- xcsh_authentication

Manages a Authentication resource in F5 Distributed Cloud.

<a id="canonical-3011330320313231-1232310001120031-0212231313213311-1330123301032202-0201112120130031-2230120100222102-2103121203231212-0102100201311310"></a>

## Prerequisites — xcsh_authentication / 232331300223 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0001232010322120-2102213323201231-1120111312311213-0231200000320020-1121022302003013-1202210200020023-3122303313333130-2331222021132220"></a>

## Minimal configuration — xcsh_authentication / 232331300223 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Authentication Resource Example
# Manages a Authentication resource in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Authentication configuration
resource "xcsh_authentication" "example" {
  name      = "example-authentication"
  namespace = "staging"
}
```

<a id="canonical-0120010030102303-2101333012310200-1202001011023130-1303002300220032-0220013212023033-1110203332322300-1010203311230310-1310100132022311"></a>

## Root configuration — xcsh_authentication / 232331300223 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-2232103110320300-2311132203232113-2112122323010022-0001321232120313-3011332102022331-2011332222201113-2002023020013232-2202222011003222"></a>

## Next pages — xcsh_authentication / 232331300223 / 6

- [Property reference](../guides/resources--authentication--reference--group-001.md#canonical-3010313311100301-1131100201313113-0013112032033123-0023231203231232-1221330222323210-2213323003220101-1110202313031123-3231300013023003)
- [Examples](../guides/resources--authentication--examples--group-001.md#canonical-2222002030332030-2211020200323212-2310130011032302-0013021012021022-2211200022113033-1213312013012133-0022330111032222-3012301012201312)
- [Import](../guides/resources--authentication--lifecycle--group-001.md#canonical-3232010031111121-3102203230033011-2100313102221210-3023201111100300-2023031020230000-2220221213123332-1211111121312203-1213230233000232)
- [Timeouts](../guides/resources--authentication--lifecycle--group-001.md#canonical-1013021012023113-1221323220312003-1323002201130022-0003321112122000-0330333130100220-3322221032103000-3313321110010022-2333221023221301)
