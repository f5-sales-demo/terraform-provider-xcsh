---
page_title: "xcsh_geo_location_set"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_geo_location_set."
---

# xcsh_geo_location_set

<a id="canonical-1001100000212311-3233013232003113-2232222223002331-1133222232220002-1101232133012211-1121021022203120-3302330011301312-1023310223132122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Overview

Breadcrumbs:

- xcsh_geo_location_set

Reads Geo Location Set information from F5 Distributed Cloud.

<a id="canonical-3132313023330011-1233122312201003-1201310303213103-1123222202122231-3311012031210101-3211022112130133-3213313222333113-3133012013001333"></a>

### Prerequisites for `xcsh_geo_location_set`

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1123201221132302-2032333211033313-1203203132133013-0302111102200123-3032023322012312-1212321030230113-0213223030333022-1130102102301200"></a>

### Minimal configuration for `xcsh_geo_location_set`

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# GeoLocationSet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing GeoLocationSet by name
data "xcsh_geo_location_set" "example" {
  name      = "example-geo-location-set"
  namespace = "system"
}

output "geo_location_set_id" {
  value = data.xcsh_geo_location_set.example.id
}
```

<a id="canonical-2200203223200103-1222133020121221-1123121333110102-2113331333021013-1113311121121213-3203323000200200-3220220221330223-1100001101032200"></a>

### Root configuration for `xcsh_geo_location_set`

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-1321131200020113-2011311200232232-0310010303031113-2233023221222031-0333003123131301-3323220310010021-0311210010000133-2130223121002121"></a>

### Explore this collection for `xcsh_geo_location_set`

- [Property reference](../guides/data-sources--geo_location_set--reference--group-001.md#canonical-1123331121000122-0300310122302103-1323002103311113-2222120103331101-2033133122011030-3102030113320122-1333103020013031-1112201210133330)
- [Examples](../guides/data-sources--geo_location_set--examples--group-001.md#canonical-2031131031200303-2230312001031111-3111020122302131-2121021112302021-2301112222112312-1300332102121221-0000201012031202-2220313030203111)
