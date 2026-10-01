---
page_title: "xcsh_fast_acl landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_fast_acl landing."
---

# xcsh_fast_acl landing

<a id="canonical-2322123002122313-2031220232103020-2333111231320022-3123222111033001-0102222022231222-0302321110322312-0220310012220300-1313311111202102"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2302320133101021-0112223222133222-3220002113333130-3302102022323121-1231110103120000-0000213002023033-2302020131031020-2021012210213303"></a>

## xcsh_fast_acl — xcsh_fast_acl / 130303022212 / 2

Breadcrumbs:

- xcsh_fast_acl

Manages object, object contains rules to protect site from denial of service It has
destination\{destination IP, destination port) and references to in F5 Distributed Cloud.

<a id="canonical-2233011301231130-3013220102132023-0020210210123213-2312300300332302-2122220112010010-1011331133320320-2001310300303110-1312313120231001"></a>

## Prerequisites — xcsh_fast_acl / 130303022212 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1200200332330023-3002003233012100-0130200003232120-3230021232011301-0011231030033120-1222123101311211-1012011212020012-3303303031223310"></a>

## Minimal configuration — xcsh_fast_acl / 130303022212 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# FastACL Resource Example
# Manages object, object contains rules to protect site from denial of service It has destination{destination IP, destination port) and references to in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic FastACL configuration
resource "xcsh_fast_acl" "example" {
  name      = "example-fast-acl"
  namespace = "system"
}
```

<a id="canonical-2013113312033230-0111323311101232-2122233002321020-1221000032012312-0001323122100201-0223023333000010-1331321032031131-3010123133230301"></a>

## Root configuration — xcsh_fast_acl / 130303022212 / 5

Required root properties: `name`. Full root flags and choices appear in the property reference.

<a id="canonical-2303313010112130-0120121312310002-1303103223123123-0030032100331211-1221012220330220-1103230313102120-0321102323133322-2210022202101122"></a>

## Next pages — xcsh_fast_acl / 130303022212 / 6

- [Property reference](../guides/resources--fast_acl--reference--group-001.md#canonical-3201133232131201-2322231132320321-1033320031233233-0020231123220112-2312100033212003-1110113101032211-2231230210300212-1132230330212312)
- [Examples](../guides/resources--fast_acl--examples--group-001.md#canonical-3012301000111300-0023322013033311-2330022311021130-0100221220023103-3210201301300332-3132321111312202-0301331330030113-3010010023330031)
- [Import](../guides/resources--fast_acl--lifecycle--group-001.md#canonical-3202231102020100-1220100310032113-0310322231321101-3212212333303121-0332030121023130-2203131113212030-3031212221311203-3301030331301112)
- [Timeouts](../guides/resources--fast_acl--lifecycle--group-001.md#canonical-3111122012233222-3221212212230230-1021103013222132-3200033313302311-1213213023313033-0131302010313103-3323101321110320-3123211132021112)
