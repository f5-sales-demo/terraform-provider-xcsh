---
page_title: "xcsh_segment_connection examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_segment_connection examples."
---

# xcsh_segment_connection examples

<a id="canonical-0210323013221130-1201212220333012-3110212033212013-0023101300000031-2133232101000320-2013333303022112-0013110032021210-1030002313233133"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_segment_connection](../data-sources/segment_connection.md#canonical-0030011210101331-2323210032230002-1133223312021002-2000100321331332-0221323131333131-3323103302112300-2012332113200223-2222011321202323)
- Examples

<a id="canonical-0330123100120312-1211321100223130-3133313311313113-1111202113310231-2121232212103202-0022201301311303-0311221231101130-0220030333110131"></a>

### Complete configurations for `xcsh_segment_connection`

- [Data source](data-sources--segment_connection--examples--group-001.md#canonical-0013320203022301-1021301010103003-1011302323121122-3120003012320313-0222211130122322-0322210203003023-0032230222331211-2213301133113003): valid configuration.

<a id="canonical-0013320203022301-1021301010103003-1011302323121122-3120003012320313-0222211130122322-0322210203003023-0032230222331211-2213301133113003"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_segment_connection](../data-sources/segment_connection.md#canonical-0030011210101331-2323210032230002-1133223312021002-2000100321331332-0221323131333131-3323103302112300-2012332113200223-2222011321202323)
- [Examples](data-sources--segment_connection--examples--group-001.md#canonical-0210323013221130-1201212220333012-3110212033212013-0023101300000031-2133232101000320-2013333303022112-0013110032021210-1030002313233133)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_segment_connection/data-source.tf`; digest `sha256:d15a7947bb7cd84484433c713ef71c626c3d8bc137b6e6bda822fbc845d28228`.

```terraform
# SegmentConnection Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SegmentConnection by name
data "xcsh_segment_connection" "example" {
  name      = "example-segment-connection"
  namespace = "staging"
}

output "segment_connection_id" {
  value = data.xcsh_segment_connection.example.id
}
```
