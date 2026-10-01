---
page_title: "xcsh_nat_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_nat_policy landing."
---

# xcsh_nat_policy landing

<a id="canonical-2211333230110221-1023302013201300-2020112103300231-1033030223223010-0020003220130132-3323021301233201-1030131302223322-2302033131310121"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1321233202103320-1303303300122030-3010301010313012-3000213322023120-0333232022202122-2033321112230221-0111213011020213-0132013312302310"></a>

## xcsh_nat_policy — xcsh_nat_policy / 033111302233 / 2

Breadcrumbs:

- xcsh_nat_policy

Manages a NAT Policy resource in F5 Distributed Cloud for nat policy create specification configures
nat policy with multiple rules,. configuration.

<a id="canonical-2122202110202000-2303101031131131-0333220313332203-1022230201211111-2110022131303111-2223220011232023-0310132123310221-3001032113303333"></a>

## Prerequisites — xcsh_nat_policy / 033111302233 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0300030000003030-0311011101313002-0303112000313310-3231032301202313-2113222110103033-3023210223211131-3221020232321322-3133031311300200"></a>

## Minimal configuration — xcsh_nat_policy / 033111302233 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-0122030132310230-0003132002121201-1211202110010023-1322223113231220-2101022030001300-1032203112313100-0010211320001201-2333211211223010"></a>

## Root configuration — xcsh_nat_policy / 033111302233 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1121202120232011-0331321123001321-2203023001320133-2311301321320200-2210113020203333-1230200130303032-3112013202230100-3300110220133110"></a>

## Next pages — xcsh_nat_policy / 033111302233 / 6

- [Property reference](../guides/data-sources--nat_policy--reference--group-001.md#canonical-0313000023303013-1312311300100203-2312121103033010-0011202012203121-2320110020132021-0013320012122120-1332300232020201-2300013013202121)
- [Examples](../guides/data-sources--nat_policy--examples--group-001.md#canonical-0221111221232231-3310223012110303-2210213320100101-0123203303130130-2302220210310111-0301033300311232-3220032012320122-0310201001003012)
