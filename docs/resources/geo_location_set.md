---
page_title: "xcsh_geo_location_set"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_geo_location_set."
---

# xcsh_geo_location_set

<a id="canonical-3103011222330300-2233200322023210-0200023031210213-3231332213130232-1013021130033033-2013330101030012-3301030030300101-3321313130311100"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_geo_location_set

Manages Geolocation Set in F5 Distributed Cloud.

<a id="canonical-3202023000013110-0033022132212111-1233233322203313-2212123220131333-1333121200102013-2210022121221302-2330023220202121-2020223023031301"></a>

### Prerequisites for `xcsh_geo_location_set`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3321303202231120-0311011130020012-1213320010000231-0111302221232121-2203223120022203-1210230200300211-1132121132133312-2211003022123321"></a>

### Minimal configuration for `xcsh_geo_location_set`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# GeoLocationSet Resource Example
# Manages Geolocation Set in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic GeoLocationSet configuration
resource "xcsh_geo_location_set" "example" {
  name      = "example-geo-location-set"
  namespace = "system"
}
```

<a id="canonical-2222221023022101-2301122110311010-1130310010111211-0103333203120111-1030202120110100-3231012332131010-3001032031113102-1211320210002013"></a>

### Root configuration for `xcsh_geo_location_set`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-0022231302032333-0300023102002010-1133102122332013-0332003000311332-1310132302012031-0233010233010221-0113333123031013-3121201313211101"></a>

### Explore this collection for `xcsh_geo_location_set`

- [Property reference](../guides/resources--geo_location_set--reference--group-001.md#canonical-1123112132012331-2320021312120003-1120132113110332-0331323012021123-1320320333300013-2001331312211030-0311313333232210-2132322132112012)
- [Examples](../guides/resources--geo_location_set--examples--group-001.md#canonical-0223332013113021-0001332130320020-0311010023022210-2023020203101321-1102220131101130-1210211322100322-0210031310303333-2200320122323013)
- [Import](../guides/resources--geo_location_set--lifecycle--group-001.md#canonical-3301032132000130-1032133211221023-2031110121122332-0010222332232302-0130021133030310-3123101011102310-3211122100122121-3301113232230002)
- [Timeouts](../guides/resources--geo_location_set--lifecycle--group-001.md#canonical-3123230222023010-3201310023201202-3031220130310010-3332020111133300-1210132122103321-1303233103132322-0031202021023202-2021102233122032)
