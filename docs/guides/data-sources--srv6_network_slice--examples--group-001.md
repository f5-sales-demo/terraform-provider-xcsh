---
page_title: "xcsh_srv6_network_slice examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_srv6_network_slice examples."
---

# xcsh_srv6_network_slice examples

<a id="canonical-2302231300112313-0022112123132022-3101033201332331-3132003300223310-2230202223132110-1302301203313020-2001202320332301-2031322030010122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_srv6_network_slice](../data-sources/srv6_network_slice.md#canonical-0001322321013132-3303233021210322-1322230110201110-3330310302210111-3123030332122130-1230012021313000-1121212210223211-3123021213301223)
- Examples

<a id="canonical-1300300300330211-1113312201311112-2210331230323100-0103331122022110-3332202333133302-0210222230131221-1201302121331001-2310322332312012"></a>

### Complete configurations for `xcsh_srv6_network_slice`

- [Data source](data-sources--srv6_network_slice--examples--group-001.md#canonical-3202233013313201-3100332002310112-3103033200121200-3031101001110300-1210103233130211-1122300322000211-2132011222121230-1212022312211011): valid configuration.

<a id="canonical-3202233013313201-3100332002310112-3103033200121200-3031101001110300-1210103233130211-1122300322000211-2132011222121230-1212022312211011"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_srv6_network_slice](../data-sources/srv6_network_slice.md#canonical-0001322321013132-3303233021210322-1322230110201110-3330310302210111-3123030332122130-1230012021313000-1121212210223211-3123021213301223)
- [Examples](data-sources--srv6_network_slice--examples--group-001.md#canonical-2302231300112313-0022112123132022-3101033201332331-3132003300223310-2230202223132110-1302301203313020-2001202320332301-2031322030010122)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_srv6_network_slice/data-source.tf`; digest `sha256:5f3b6cfc9cb152dc2c604228803d1e4acdb80bbc9c71bb42003bba14ed6bbdd1`.

```terraform
# Srv6NetworkSlice Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Srv6NetworkSlice by name
data "xcsh_srv6_network_slice" "example" {
  name      = "example-srv6-network-slice"
  namespace = "system"
}

output "srv6_network_slice_id" {
  value = data.xcsh_srv6_network_slice.example.id
}
```
