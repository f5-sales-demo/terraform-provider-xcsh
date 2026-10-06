---
page_title: "xcsh_device_intelligence_summary examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_device_intelligence_summary examples."
---

# xcsh_device_intelligence_summary examples

<a id="canonical-3321132220332130-1310310320203213-2303202233130231-1232203011031011-2321031101113113-1210300031010010-3013101010002313-1310231302022030"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_device_intelligence_summary](../data-sources/device_intelligence_summary.md#canonical-1301033300112123-2331103013013311-3321132033101320-1310222002112032-1233113332311133-1310100202102222-2023210120211201-1120100233333321)
- Examples

<a id="canonical-1330220001110101-3113212330133220-3201121321032202-0331111233030322-3021321022201000-1130100222310000-2212003323203201-3213013210222312"></a>

### Complete configurations for `xcsh_device_intelligence_summary`

- [Data source](data-sources--device_intelligence_summary--examples--group-001.md#canonical-3202102222333320-1312230102022103-1231223221131112-3103012033022012-2020131322201223-0011302202233301-1302130221202113-1302013032011121): valid configuration.

<a id="canonical-3202102222333320-1312230102022103-1231223221131112-3103012033022012-2020131322201223-0011302202233301-1302130221202113-1302013032011121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_device_intelligence_summary](../data-sources/device_intelligence_summary.md#canonical-1301033300112123-2331103013013311-3321132033101320-1310222002112032-1233113332311133-1310100202102222-2023210120211201-1120100233333321)
- [Examples](data-sources--device_intelligence_summary--examples--group-001.md#canonical-3321132220332130-1310310320203213-2303202233130231-1232203011031011-2321031101113113-1210300031010010-3013101010002313-1310231302022030)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_device_intelligence_summary/data-source.tf`; digest `sha256:a3b5af58871245c4f3794bc2ed886a4421185806d4cb6b7c6eb5328aa94877da`.

```terraform
# DeviceIntelligenceSummary DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_device_intelligence_summary" "example" {
  namespace = "example-value"
}

output "device_intelligence_summary_result" {
  value = data.xcsh_device_intelligence_summary.example
}
```
