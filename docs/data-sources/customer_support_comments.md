---
page_title: "xcsh_customer_support_comments landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_customer_support_comments landing."
---

# xcsh_customer_support_comments landing

<a id="canonical-3233331010300203-0002001013222330-0201130211303202-2232020123310031-2211320023331130-0033233023122313-1203101322331330-3321120232032311"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2321231020201211-2121300201302321-3312110310200301-2131031100201011-0330231032200032-3333331302013302-0032022321232112-1212011031213102"></a>

## xcsh_customer_support_comments — xcsh_customer_support_comments / 123130333011 / 2

Breadcrumbs:

- xcsh_customer_support_comments

Resource retrieval operation.

<a id="canonical-2200122213000130-0312233332131231-0121300111311121-0133102223332321-1002220122011332-2031222030212332-2231331011230223-0023320301001323"></a>

## Prerequisites — xcsh_customer_support_comments / 123130333011 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1303303312112110-2223302322002000-0301122020313220-2203331111310220-0130231201010313-2011132000033133-3110202122220120-2003202312203131"></a>

## Minimal configuration — xcsh_customer_support_comments / 123130333011 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-3302312300211120-0310131130213230-0330303213013312-3212031031031102-2300330121012223-2113130022322320-2320122001212321-1033133321002012"></a>

## Root configuration — xcsh_customer_support_comments / 123130333011 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-2012132312302213-2001322032030001-3122000010110012-3103300103330020-2003103023113102-3110011102032233-3101002232302011-2131330223302202"></a>

## Next pages — xcsh_customer_support_comments / 123130333011 / 6

- [Property reference](../guides/data-sources--customer_support_comments--reference--group-001.md#canonical-0102322120201321-0003320313130322-1013202100231000-2332022030311102-0212020231220133-3100100201113001-0211003001111232-3222003310331013)
- [Examples](../guides/data-sources--customer_support_comments--examples--group-001.md#canonical-0130002102312203-2202333001113020-2200101033221310-3221123000133030-3033002210130230-2230333111013011-2203112221121131-2021133013032331)
