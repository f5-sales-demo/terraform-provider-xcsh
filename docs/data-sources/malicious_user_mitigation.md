---
page_title: "xcsh_malicious_user_mitigation landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_malicious_user_mitigation landing."
---

# xcsh_malicious_user_mitigation landing

<a id="canonical-1212221312012302-1222123021331020-2101212022233322-1130010233001121-1210032201302202-2321123033111102-3123131300131021-3230032031232211"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3201310221012232-3131130013120231-2012022132203110-0022012310101002-3322310112203230-2130022221200312-3230113103011022-2101330202222333"></a>

## xcsh_malicious_user_mitigation — xcsh_malicious_user_mitigation / 120103000003 / 2

Breadcrumbs:

- xcsh_malicious_user_mitigation

Manages malicious\_user\_mitigation creates a new object in the storage backend for
metadata.namespace in F5 Distributed Cloud.

<a id="canonical-1010310131010100-2132203020230100-1010310133211320-3121003321321231-3013021012013313-3322112310211302-2211023110013213-1012101031001200"></a>

## Prerequisites — xcsh_malicious_user_mitigation / 120103000003 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3210331022303033-3033211233300013-0321101201113310-0110001303213230-0112130112221023-1312013220123122-2311313010331220-3100221222033332"></a>

## Minimal configuration — xcsh_malicious_user_mitigation / 120103000003 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# MaliciousUserMitigation Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing MaliciousUserMitigation by name
data "xcsh_malicious_user_mitigation" "example" {
  name      = "example-malicious-user-mitigation"
  namespace = "staging"
}

output "malicious_user_mitigation_id" {
  value = data.xcsh_malicious_user_mitigation.example.id
}
```

<a id="canonical-0323113323311230-0102021233033202-0302332331132132-2301211320331211-3013101130230312-3201200212130113-2333333011222102-1233103100222130"></a>

## Root configuration — xcsh_malicious_user_mitigation / 120103000003 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-1220033000302332-3212003212323230-3030111030010300-3132013100231013-2032320233332211-3121331320132000-1123001111001203-0210111230311002"></a>

## Next pages — xcsh_malicious_user_mitigation / 120103000003 / 6

- [Property reference](../guides/data-sources--malicious_user_mitigation--reference--group-001.md#canonical-3221231120332312-2313110303121201-0120132133321332-2132312311233101-0102202111332010-1032122213331232-1113303003013220-0122121011032230)
- [Examples](../guides/data-sources--malicious_user_mitigation--examples--group-001.md#canonical-2330111100130201-3001331323202112-0233331232123201-0231220331113032-0222113203313111-3200320101001133-2111100013032023-3232130123320213)
