---
page_title: "xcsh_network_connector examples"
subcategory: "Networking"
description: "Complete grouped canonical reference for xcsh_network_connector examples."
---

# xcsh_network_connector examples

<a id="canonical-3111123311200233-3221113313323111-1030301032222222-2233132330113212-2012130321031310-0022330033310111-0021302301033202-2220033120232122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- Examples

<a id="canonical-1101122201113232-2230021020233133-3333133003301331-0113331123100002-0021232303110232-1001311020003230-1222002001313221-2001211332331332"></a>

### Complete configurations for `xcsh_network_connector`

- [Data source](data-sources--network_connector--examples--group-001.md#canonical-0023000011301211-0311132300203303-1010013212013123-0311323010132310-3232010101200001-0331121220320220-2333103333313230-0012312220023213): valid configuration.

<a id="canonical-0023000011301211-0311132300203303-1010013212013123-0311323010132310-3232010101200001-0331121220320220-2333103333313230-0012312220023213"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_network_connector](../data-sources/network_connector.md#canonical-3100100033113131-1013100333112132-1220212103012023-0211023132210233-2323121020123322-2323213203113121-1203301130023023-3202032133312301)
- [Examples](data-sources--network_connector--examples--group-001.md#canonical-3111123311200233-3221113313323111-1030301032222222-2233132330113212-2012130321031310-0022330033310111-0021302301033202-2220033120232122)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_connector/data-source.tf`; digest `sha256:842679078281a092ba5ea7fed221d6b2c452c242a5d7cc3a8d19f42931b95be3`.

```terraform
# NetworkConnector Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NetworkConnector by name
data "xcsh_network_connector" "example" {
  name      = "example-network-connector"
  namespace = "staging"
}

output "network_connector_id" {
  value = data.xcsh_network_connector.example.id
}
```
