---
page_title: "xcsh_waf_attack_signatures landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_waf_attack_signatures landing."
---

# xcsh_waf_attack_signatures landing

<a id="canonical-1002120312121031-2231120213022032-3113323210101313-3123231202123002-0000220212231033-3012130333333223-0330122132233323-2033323020101321"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-1210101302311103-2022012033012301-0111012210100201-1123321003122122-1033001122303320-0210232302212022-2130313003113102-2320221200320301"></a>

## xcsh_waf_attack_signatures — xcsh_waf_attack_signatures / 012221023110 / 2

Breadcrumbs:

- xcsh_waf_attack_signatures

Resource retrieval operation.

<a id="canonical-3131130323011100-3000310003330332-3112312223223222-2203223313212103-3331030000010310-1121300122033021-1101322210303132-3232012201121123"></a>

## Prerequisites — xcsh_waf_attack_signatures / 012221023110 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-1233120100012021-1101210333032222-1202331231020302-3223023113333103-2132001001102122-3022223202020031-2232211312123223-3031311111002102"></a>

## Minimal configuration — xcsh_waf_attack_signatures / 012221023110 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# WAFAttackSignatures DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_waf_attack_signatures" "example" {
}

output "waf_attack_signatures_result" {
  value = data.xcsh_waf_attack_signatures.example
}
```

<a id="canonical-2002330311320311-3211201123022331-2112232110123123-0113123233202002-2101222212211301-1203333133120101-3102032011221222-0000132233122221"></a>

## Root configuration — xcsh_waf_attack_signatures / 012221023110 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-1332022232232232-0103310223222223-0132030102331313-1132101200133301-3300123123122030-0333120323111100-1330322003123222-0321322032120312"></a>

## Next pages — xcsh_waf_attack_signatures / 012221023110 / 6

- [Property reference](../guides/data-sources--waf_attack_signatures--reference--group-001.md#canonical-3120200000301112-2000001222010231-1030000101003120-0002131010003101-1121111122231111-2121223333310231-1312321220022310-2131323003202332)
- [Examples](../guides/data-sources--waf_attack_signatures--examples--group-001.md#canonical-1000123133012313-3302212201201323-1032223012210030-3000332321001301-3222211122333313-2033112213131303-0300331331022233-3010003031111302)
