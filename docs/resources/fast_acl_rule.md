---
page_title: "xcsh_fast_acl_rule landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl_rule landing."
---

# xcsh_fast_acl_rule landing

<a id="canonical-3211210133012011-3013232222012110-0223303130203332-3111320233310102-0302012020022011-2023213303223030-1010212231200312-0211302221203302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3003201310222011-3303111213330203-0312011113033301-0313331221010012-1030331231310202-3031333331321113-2023122032333330-1321320323230021"></a>

## xcsh_fast_acl_rule — xcsh_fast_acl_rule / 011301003110 / 2

Breadcrumbs:

- xcsh_fast_acl_rule

Manages new Fast ACL rule, has specification to match source IP, source port and action to apply in
F5 Distributed Cloud.

<a id="canonical-1013022003311300-2031023300133013-2221313211212203-0020122300001323-1221302202020230-2322231323112123-3333233121323201-0111221031301303"></a>

## Prerequisites — xcsh_fast_acl_rule / 011301003110 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-0212301333201320-3322120203213203-2203211103111021-3312220323100033-3210333100121012-3212001021230202-1123101100330021-2223222030021022"></a>

## Minimal configuration — xcsh_fast_acl_rule / 011301003110 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FastACLRule Resource Example
# Manages new Fast ACL rule, has specification to match source IP, source port and action to apply in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FastACLRule configuration
resource "xcsh_fast_acl_rule" "example" {
  name      = "example-fast-acl-rule"
  namespace = "staging"
}
```

<a id="canonical-3013301311212122-3231200002302103-2032201210322201-2332221130030230-0320003032201100-3102302210331203-2321011312020201-1220132301232131"></a>

## Root configuration — xcsh_fast_acl_rule / 011301003110 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-3233022203112131-2103333033120111-2010100123001232-3033312030232313-0033302021033211-1022021331013032-2010010330133110-3110101023223223"></a>

## Next pages — xcsh_fast_acl_rule / 011301003110 / 6

- [Property reference](../guides/resources--fast_acl_rule--reference--group-001.md#canonical-0011333330012010-1222121310111312-3121011031301112-1122322123011322-2031013302131013-2030220000333201-3233233332230221-2203000100020232)
- [Examples](../guides/resources--fast_acl_rule--examples--group-001.md#canonical-2332303311010220-2331002222330022-3222333230321022-3200311020123320-1001311302002013-0210110221030102-3230030311031021-3002010323233311)
- [Import](../guides/resources--fast_acl_rule--lifecycle--group-001.md#canonical-3102010323103301-2013303330301003-0033012121311313-2330203230320023-1310322012120331-2310300012202032-3000222200323132-3232022211230200)
- [Timeouts](../guides/resources--fast_acl_rule--lifecycle--group-001.md#canonical-1221202020011101-3021301202112232-0231031001121230-2013120320112213-3100113102131011-2132333311232321-1003121303030011-2123110032233313)
