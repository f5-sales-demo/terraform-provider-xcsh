---
page_title: "xcsh_dns_zone_edit_cryptokey examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_edit_cryptokey examples."
---

# xcsh_dns_zone_edit_cryptokey examples

<a id="canonical-1013213223203320-0132011103113312-1210300032133231-0200130202001120-1130123320223331-0333013021300110-0212011033121132-3230323323103232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dns_zone_edit_cryptokey](../actions/dns_zone_edit_cryptokey.md#canonical-1132220310333100-2000123203101123-2131221001021322-3330222102211322-3331333302020203-3030331000102011-2023032220122321-2312001020231030)
- Examples

<a id="canonical-3330130200200213-2322003131202000-1023310100110223-2222323000013000-3210120001231312-0202301133121113-0112031313323312-0130203331020010"></a>

### Complete configurations for `xcsh_dns_zone_edit_cryptokey`

- [Action](actions--dns_zone_edit_cryptokey--examples--group-001.md#canonical-2131110033132201-0211303211102333-2013231203000120-2003230333001121-3001320202021223-1310110313112110-0130123330323322-2112303110223021): valid configuration.

<a id="canonical-2131110033132201-0211303211102333-2013231203000120-2003230333001121-3001320202021223-1310110313112110-0130123330323322-2112303110223021"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Action example

Breadcrumbs:

- [xcsh_dns_zone_edit_cryptokey](../actions/dns_zone_edit_cryptokey.md#canonical-1132220310333100-2000123203101123-2131221001021322-3330222102211322-3331333302020203-3030331000102011-2023032220122321-2312001020231030)
- [Examples](actions--dns_zone_edit_cryptokey--examples--group-001.md#canonical-1013213223203320-0132011103113312-1210300032133231-0200130202001120-1130123320223331-0333013021300110-0212011033121132-3230323323103232)
- Action

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/actions/xcsh_dns_zone_edit_cryptokey/action.tf`; digest `sha256:377745c7c273966a8712f1843e6db7821e8429814e456ec61ea024bb03bd4506`.

```terraform
# DNSZoneEditCryptokey Action Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

action "xcsh_dns_zone_edit_cryptokey" "example" {
  config {
  }
}
```
