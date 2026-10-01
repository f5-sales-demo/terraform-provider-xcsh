---
page_title: "xcsh_registration landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_registration landing."
---

# xcsh_registration landing

<a id="canonical-0213101222032121-3202213230020032-1300311133322023-3102010212010010-0302023033222033-2230311223131211-1221110133201103-1300301230303301"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1023323303000132-0020222112121032-0021030032323311-3112022010002322-1000332303013121-1022101233220321-3302321203120003-3200000013013001"></a>

## xcsh_registration — xcsh_registration / 123103030122 / 2

Breadcrumbs:

- xcsh_registration

Manages a Registration resource in F5 Distributed Cloud for vpm creates registration using this
message, never used by users. configuration.

<a id="canonical-1103210001321322-3001332022213222-2110312221220111-3023223201301332-1330020331023232-2230221322212300-2321112031233121-2310320210323132"></a>

## Prerequisites — xcsh_registration / 123103030122 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3022323332312210-1022113103301202-0220302002221003-2320022332021003-1223012010132210-2131331121122100-1210301023211103-2123312301301030"></a>

## Minimal configuration — xcsh_registration / 123103030122 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Registration Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Registration by name
data "xcsh_registration" "example" {
  name      = "example-registration"
  namespace = "staging"
}

output "registration_id" {
  value = data.xcsh_registration.example.id
}
```

<a id="canonical-3001112130011331-0333000010021112-0133123330212221-0110122022310311-2031122103303331-0231110221013232-3321033222102223-0123221010033010"></a>

## Root configuration — xcsh_registration / 123103030122 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0331330330230231-3200130213121220-2200311333222002-1130122223312001-1230130133310000-1320213203233003-1213303111120210-0012321131333213"></a>

## Next pages — xcsh_registration / 123103030122 / 6

- [Property reference](../guides/data-sources--registration--reference--group-001.md#canonical-1101112222230233-3310100112220233-0122221010111210-1032100101220210-0000320131100022-3121310301031310-1210101132122023-1010212333103012)
- [Examples](../guides/data-sources--registration--examples--group-001.md#canonical-0300203131021300-1230011003310011-0211010322121322-1122310000121231-2202100310321000-2020031132033031-3220322213011231-2031002333213313)
