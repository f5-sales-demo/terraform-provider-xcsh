---
page_title: "xcsh_geo_location_set examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_geo_location_set examples."
---

# xcsh_geo_location_set examples

<a id="canonical-0223332013113021-0001332130320020-0311010023022210-2023020203101321-1102220131101130-1210211322100322-0210031310303333-2200320122323013"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-3103011222330300-2233200322023210-0200023031210213-3231332213130232-1013021130033033-2013330101030012-3301030030300101-3321313130311100)
- Examples

<a id="canonical-1313101012331033-1131122031333310-2203222223302302-2300111223320010-2120303102033212-2333201210032310-1301210230210030-1223013132102000"></a>

### Complete configurations for `xcsh_geo_location_set`

- [Resource](resources--geo_location_set--examples--group-001.md#canonical-0330213121232311-1231200102323110-2202312321201323-3121322002131301-2333220120310312-0213030330132220-3230200133202333-1333002123232002): valid configuration.

<a id="canonical-0330213121232311-1231200102323110-2202312321201323-3121322002131301-2333220120310312-0213030330132220-3230200133202333-1333002123232002"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Resource example

Breadcrumbs:

- [xcsh_geo_location_set](../resources/geo_location_set.md#canonical-3103011222330300-2233200322023210-0200023031210213-3231332213130232-1013021130033033-2013330101030012-3301030030300101-3321313130311100)
- [Examples](resources--geo_location_set--examples--group-001.md#canonical-0223332013113021-0001332130320020-0311010023022210-2023020203101321-1102220131101130-1210211322100322-0210031310303333-2200320122323013)
- Resource

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/resources/xcsh_geo_location_set/resource.tf`; digest `sha256:0e7cc4dbb2f138697448312f9d3840d2e253e947097376d3ac1a000ac5c54714`.

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
