---
page_title: "xcsh_trusted_ca_list landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_trusted_ca_list landing."
---

# xcsh_trusted_ca_list landing

<a id="canonical-3231112200302210-0000022221223030-0200201111213211-3101031133123233-3133122221310202-1301130023031133-3210320322301322-1320111311102122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1120203123312200-1000301310332200-2211102232010131-1031332013000002-3120313220233011-2300113233012312-3023002110011331-1102211301323013"></a>

## xcsh_trusted_ca_list — xcsh_trusted_ca_list / 202300200201 / 2

Breadcrumbs:

- xcsh_trusted_ca_list

Manages a Trusted CA List resource in F5 Distributed Cloud for trusted certificate authority list
management.

<a id="canonical-1100213233320323-0012000310301111-1201323201322033-0223111232001113-0232213030330213-2322210120000202-3011113102122100-3322012000113031"></a>

## Prerequisites — xcsh_trusted_ca_list / 202300200201 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-2103112232102110-2102311221113121-1023313331223021-3022123011101331-3322203201232312-2111230000132333-1131000123212332-0000323232220203"></a>

## Minimal configuration — xcsh_trusted_ca_list / 202300200201 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# TrustedCAList Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing TrustedCAList by name
data "xcsh_trusted_ca_list" "example" {
  name      = "example-trusted-ca-list"
  namespace = "staging"
}

output "trusted_ca_list_id" {
  value = data.xcsh_trusted_ca_list.example.id
}
```

<a id="canonical-2330132313232133-3310333313231332-0332311113313012-0330323330221300-1003002032110213-3122323122002322-1201222130032232-2021020033222002"></a>

## Root configuration — xcsh_trusted_ca_list / 202300200201 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-0110331322221232-1212301003231001-0300200210120313-0011221013132133-0122103220310212-0231111032223330-3102002132001021-2003202022211322"></a>

## Next pages — xcsh_trusted_ca_list / 202300200201 / 6

- [Property reference](../guides/data-sources--trusted_ca_list--reference--group-001.md#canonical-3130232031111032-3303330311222223-3022110011301203-1122313200231131-2032233133020322-0030131110322322-1101000221212310-3022300311330213)
- [Examples](../guides/data-sources--trusted_ca_list--examples--group-001.md#canonical-3011330123323123-2011120213211131-0222331311023223-2312020212011321-2112233113021100-0211101033231202-2301121002211210-3303202233210020)
