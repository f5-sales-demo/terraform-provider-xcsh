---
page_title: "xcsh_ike1 examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_ike1 examples."
---

# xcsh_ike1 examples

<a id="canonical-3132332313110102-2301103110101323-3013231113332213-2101031223111032-3332331103131323-1221301031101331-1200200320022232-2002033021333210"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)
- Examples

<a id="canonical-1323003332301100-0023023233023210-0233031100033110-0030221002310223-1221123100030112-3001300230103010-2312320122102211-1012002220021033"></a>

### Complete configurations for `xcsh_ike1`

- [Data source](data-sources--ike1--examples--group-001.md#canonical-3022001323333201-2001223231333011-3210311031311210-3033232202212002-0021323023202230-2002021111130230-3113133000002103-0213310311101300): valid configuration.

<a id="canonical-3022001323333201-2001223231333011-3210311031311210-3033232202212002-0021323023202230-2002021111130230-3113133000002103-0213310311101300"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_ike1](../data-sources/ike1.md#canonical-2112011122002111-1121301111312122-3021231213103111-3131302213110331-0301300202011011-1200320222013102-3321012020321333-1000030102103313)
- [Examples](data-sources--ike1--examples--group-001.md#canonical-3132332313110102-2301103110101323-3013231113332213-2101031223111032-3332331103131323-1221301031101331-1200200320022232-2002033021333210)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_ike1/data-source.tf`; digest `sha256:1f5aa2a482784f4287f2518ad47fd88a6918d0ecdd01b6d56a7fd34d407e402f`.

```terraform
# Ike1 Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Ike1 by name
data "xcsh_ike1" "example" {
  name      = "example-ike1"
  namespace = "staging"
}

output "ike1_id" {
  value = data.xcsh_ike1.example.id
}
```
