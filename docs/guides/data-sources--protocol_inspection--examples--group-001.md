---
page_title: "xcsh_protocol_inspection examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_protocol_inspection examples."
---

# xcsh_protocol_inspection examples

<a id="canonical-2222200023121001-2120130320002212-3002302310233020-3021231113212123-3120033122221302-2013033122021310-1313023233113113-1023111323332123"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)
- Examples

<a id="canonical-3100033332110322-0102132320322030-3213102110023231-2003222101302223-1103202232212301-0300133212213222-1033130220122321-2313331220013231"></a>

### Complete configurations for `xcsh_protocol_inspection`

- [Data source](data-sources--protocol_inspection--examples--group-001.md#canonical-2321033120121020-3101302133330101-2131210100112013-1223322021203030-3211213013202023-3112213231210012-2132113032320231-1203012130102321): valid configuration.

<a id="canonical-2321033120121020-3101302133330101-2131210100112013-1223322021203030-3211213013202023-3112213231210012-2132113032320231-1203012130102321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_protocol_inspection](../data-sources/protocol_inspection.md#canonical-1313301302113001-0102313332203012-2202123331012332-2332033323132031-2332220030232012-3021230003332301-2212111100302303-0210333132120310)
- [Examples](data-sources--protocol_inspection--examples--group-001.md#canonical-2222200023121001-2120130320002212-3002302310233020-3021231113212123-3120033122221302-2013033122021310-1313023233113113-1023111323332123)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_protocol_inspection/data-source.tf`; digest `sha256:8e357f2e82e2e660e9b2335aec0a0bf6fb5efafdd6746da23a0a982fc6492237`.

```terraform
# ProtocolInspection Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ProtocolInspection by name
data "xcsh_protocol_inspection" "example" {
  name      = "example-protocol-inspection"
  namespace = "staging"
}

output "protocol_inspection_id" {
  value = data.xcsh_protocol_inspection.example.id
}
```
