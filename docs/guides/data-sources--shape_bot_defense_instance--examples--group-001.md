---
page_title: "xcsh_shape_bot_defense_instance examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_shape_bot_defense_instance examples."
---

# xcsh_shape_bot_defense_instance examples

<a id="canonical-1100231133123201-2223300212332101-0310101230300323-0221011123303230-1201321212303220-0102322330223331-1111031331202111-0000011131120220"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_shape_bot_defense_instance](../data-sources/shape_bot_defense_instance.md#canonical-3000123112021030-0302103133011033-2210121132011122-2132220211231033-3233312310101331-1222032010313211-1333211301300023-0231320200212323)
- Examples

<a id="canonical-1301013311033022-1200212101313231-0331220210301330-1232020303220302-0000210212213322-0223031222303100-2131202332033111-2123021023302003"></a>

### Complete configurations for `xcsh_shape_bot_defense_instance`

- [Data source](data-sources--shape_bot_defense_instance--examples--group-001.md#canonical-3322212013131311-1111032230113003-1110202122221132-1033320110201221-1020230113231001-0232332121302032-2023013211302101-3233112231001022): valid configuration.

<a id="canonical-3322212013131311-1111032230113003-1110202122221132-1033320110201221-1020230113231001-0232332121302032-2023013211302101-3233112231001022"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_shape_bot_defense_instance](../data-sources/shape_bot_defense_instance.md#canonical-3000123112021030-0302103133011033-2210121132011122-2132220211231033-3233312310101331-1222032010313211-1333211301300023-0231320200212323)
- [Examples](data-sources--shape_bot_defense_instance--examples--group-001.md#canonical-1100231133123201-2223300212332101-0310101230300323-0221011123303230-1201321212303220-0102322330223331-1111031331202111-0000011131120220)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_shape_bot_defense_instance/data-source.tf`; digest `sha256:559cb2ebbea61576b9b4c66aad8123cbb2b2816c6fc4d4e4939f3b3e386387b6`.

```terraform
# ShapeBotDefenseInstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ShapeBotDefenseInstance by name
data "xcsh_shape_bot_defense_instance" "example" {
  name      = "example-shape-bot-defense-instance"
  namespace = "staging"
}

output "shape_bot_defense_instance_id" {
  value = data.xcsh_shape_bot_defense_instance.example.id
}
```
