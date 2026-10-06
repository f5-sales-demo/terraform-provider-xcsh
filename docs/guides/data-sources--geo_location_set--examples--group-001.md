---
page_title: "xcsh_geo_location_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_geo_location_set examples."
---

# xcsh_geo_location_set examples

<a id="canonical-2031131031200303-2230312001031111-3111020122302131-2121021112302021-2301112222112312-1300332102121221-0000201012031202-2220313030203111"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-1001100000212311-3233013232003113-2232222223002331-1133222232220002-1101232133012211-1121021022203120-3302330011301312-1023310223132122)
- Examples

<a id="canonical-3213011202123310-2300200001300131-2332120320022133-3213231110233113-3122100132311331-2230221230112210-1131022230001231-3000011200321231"></a>

### Complete configurations for `xcsh_geo_location_set`

- [Data source](data-sources--geo_location_set--examples--group-001.md#canonical-0230123203110203-1010212223323020-0222020203013212-0311132300332123-1011213212232302-2130110222302011-0010223032211111-1121123111310321): valid configuration.

<a id="canonical-0230123203110203-1010212223323020-0222020203013212-0311132300332123-1011213212232302-2130110222302011-0010223032211111-1121123111310321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_geo_location_set](../data-sources/geo_location_set.md#canonical-1001100000212311-3233013232003113-2232222223002331-1133222232220002-1101232133012211-1121021022203120-3302330011301312-1023310223132122)
- [Examples](data-sources--geo_location_set--examples--group-001.md#canonical-2031131031200303-2230312001031111-3111020122302131-2121021112302021-2301112222112312-1300332102121221-0000201012031202-2220313030203111)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_geo_location_set/data-source.tf`; digest `sha256:fe5ac2dcce0034284edb5717f6d3da18db76cd12233f57cc1a3e9be26885c3ea`.

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
