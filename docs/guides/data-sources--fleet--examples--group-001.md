---
page_title: "xcsh_fleet examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fleet examples."
---

# xcsh_fleet examples

<a id="canonical-1223103321010323-0231100030013302-1233302032200110-0210031301212220-2100122330323123-0310212333032022-2010022133220121-0100330002231213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- Examples

<a id="canonical-0031013332123100-0131030231112230-2223331322001011-0330213122133232-2101212310023002-3211202322120033-0303022310000010-1000332300111100"></a>

### Complete configurations for `xcsh_fleet`

- [Data source](data-sources--fleet--examples--group-001.md#canonical-1323010013013231-3003101211233233-3010120120301000-0230121132123333-1331011333002031-1211102213323202-2302012333020033-2033013320121233): valid configuration.

<a id="canonical-1323010013013231-3003101211233233-3010120120301000-0230121132123333-1331011333002031-1211102213323202-2302012333020033-2033013320121233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_fleet](../data-sources/fleet.md#canonical-3310302123210330-1013130130202112-0123313311102020-0303231032231002-2231010331133332-0313232311211100-2132313003203002-0320222110330222)
- [Examples](data-sources--fleet--examples--group-001.md#canonical-1223103321010323-0231100030013302-1233302032200110-0210031301212220-2100122330323123-0310212333032022-2010022133220121-0100330002231213)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_fleet/data-source.tf`; digest `sha256:6d30c4d1a598a1f7f950e3ba739fc0d0cef6e7cc4d4b114700efeb50ec416c17`.

```terraform
# Fleet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Fleet by name
data "xcsh_fleet" "example" {
  name      = "example-fleet"
  namespace = "staging"
}

output "fleet_id" {
  value = data.xcsh_fleet.example.id
}
```
