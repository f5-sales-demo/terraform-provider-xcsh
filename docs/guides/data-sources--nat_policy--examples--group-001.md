---
page_title: "xcsh_nat_policy examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nat_policy examples."
---

# xcsh_nat_policy examples

<a id="canonical-0221111221232231-3310223012110303-2210213320100101-0123203303130130-2302220210310111-0301033300311232-3220032012320122-0310201001003012"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- Examples

<a id="canonical-3111113000132013-2033111322011303-1032231123003002-3100330302331131-3133230201111000-1201231221203021-3213001313102121-0002121101212332"></a>

### Complete configurations for `xcsh_nat_policy`

- [Data source](data-sources--nat_policy--examples--group-001.md#canonical-2100301132120231-2110131320221302-3302122201310033-1132133200033301-3300120310310021-1200333213323021-1231103221312301-1032111132023320): valid configuration.

<a id="canonical-2100301132120231-2110131320221302-3302122201310033-1132133200033301-3300120310310021-1200333213323021-1231103221312301-1032111132023320"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_nat_policy](../data-sources/nat_policy.md#canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121)
- [Examples](data-sources--nat_policy--examples--group-001.md#canonical-0221111221232231-3310223012110303-2210213320100101-0123203303130130-2302220210310111-0301033300311232-3220032012320122-0310201001003012)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_nat_policy/data-source.tf`; digest `sha256:d086b598bab70cc64257bc21568a65534a135b8fdb72d6f6952421f88e8adafe`.

```terraform
# NATPolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing NATPolicy by name
data "xcsh_nat_policy" "example" {
  name      = "example-nat-policy"
  namespace = "staging"
}

output "nat_policy_id" {
  value = data.xcsh_nat_policy.example.id
}
```
