---
page_title: "xcsh_customer_support_comments examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_customer_support_comments examples."
---

# xcsh_customer_support_comments examples

<a id="canonical-0130002102312203-2202333001113020-2200101033221310-3221123000133030-3033002210130230-2230333111013011-2203112221121131-2021133013032331"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md#canonical-3233331010300203-0002001013222330-0201130211303202-2232020123310031-2211320023331130-0033233023122313-1203101322331330-3321120232032311)
- Examples

<a id="canonical-2012003131302122-2330232321221332-2210203120302300-0030133032132211-3303133102112322-2201012231102311-0201222310212201-1110012203323111"></a>

### Complete configurations for `xcsh_customer_support_comments`

- [Data source](data-sources--customer_support_comments--examples--group-001.md#canonical-2332130320302213-2223033010100300-2011303322211302-0303123320302333-0320133332211322-2223010030013210-3130110200121011-1123323322213321): valid configuration.

<a id="canonical-2332130320302213-2223033010100300-2011303322211302-0303123320302333-0320133332211322-2223010030013210-3130110200121011-1123323322213321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_customer_support_comments](../data-sources/customer_support_comments.md#canonical-3233331010300203-0002001013222330-0201130211303202-2232020123310031-2211320023331130-0033233023122313-1203101322331330-3321120232032311)
- [Examples](data-sources--customer_support_comments--examples--group-001.md#canonical-0130002102312203-2202333001113020-2200101033221310-3221123000133030-3033002210130230-2230333111013011-2203112221121131-2021133013032331)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_customer_support_comments/data-source.tf`; digest `sha256:0d5599711e86330f779d8085568282cccbd3e192a937997854401c5d7fbdb6a3`.

```terraform
# CustomerSupportComments DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_customer_support_comments" "example" {
  name = "example-value"
}

output "customer_support_comments_result" {
  value = data.xcsh_customer_support_comments.example
}
```
