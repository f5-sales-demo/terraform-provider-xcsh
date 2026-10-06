---
page_title: "xcsh_segment examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_segment examples."
---

# xcsh_segment examples

<a id="canonical-0210203003102130-2132223030332301-3223100302300300-3131312102123022-1311222120010123-0220311130321330-1300000003331333-1121130121113311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_segment](../data-sources/segment.md#canonical-1322202313211133-3232200032023212-3030131113213211-1232130001132103-2232301233230330-2022002332021123-3031321121233232-3101311003121100)
- Examples

<a id="canonical-2312011333300100-3213130201202313-1200311201320213-0200301131223202-3213202123210123-0121211321303010-0012103221013311-3110213311233222"></a>

### Complete configurations for `xcsh_segment`

- [Data source](data-sources--segment--examples--group-001.md#canonical-2232010232021203-2221031113210023-0131212110111310-0003221321320031-2023113000120320-0311333313333233-1012011003000102-3021220000112222): valid configuration.

<a id="canonical-2232010232021203-2221031113210023-0131212110111310-0003221321320031-2023113000120320-0311333313333233-1012011003000102-3021220000112222"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_segment](../data-sources/segment.md#canonical-1322202313211133-3232200032023212-3030131113213211-1232130001132103-2232301233230330-2022002332021123-3031321121233232-3101311003121100)
- [Examples](data-sources--segment--examples--group-001.md#canonical-0210203003102130-2132223030332301-3223100302300300-3131312102123022-1311222120010123-0220311130321330-1300000003331333-1121130121113311)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_segment/data-source.tf`; digest `sha256:2e9bfc7c69e50200df48c3eca01ce203db4e0516ed11b6bf987fdf485d490ac4`.

```terraform
# Segment Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Segment by name
data "xcsh_segment" "example" {
  name      = "example-segment"
  namespace = "system"
}

output "segment_id" {
  value = data.xcsh_segment.example.id
}
```
