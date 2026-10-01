---
page_title: "xcsh_dns_zone_cryptokeys landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_cryptokeys landing."
---

# xcsh_dns_zone_cryptokeys landing

<a id="canonical-3311023113031202-1010300221212322-1133102011021320-0033123312013310-0302320000020011-0000120213001132-1111132220110232-1220301011030302"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-0332032233211031-1112203133330331-1313301011232233-2331311230121032-0302322332002102-0230302012323203-1232111101230131-0113133031020302"></a>

## xcsh_dns_zone_cryptokeys — xcsh_dns_zone_cryptokeys / 323311123031 / 2

Breadcrumbs:

- xcsh_dns_zone_cryptokeys

Resource creation operation.

<a id="canonical-1033332033220211-3010020310033121-3311123112103320-2322122023022331-3011110303001332-3031100203122033-3111130230212002-0301303132103111"></a>

## Prerequisites — xcsh_dns_zone_cryptokeys / 323311123031 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-3101132002022012-0311123001032222-3030333112111222-3302131232233302-3333113211201111-0321123031211232-0132211023001312-2333113111211322"></a>

## Minimal configuration — xcsh_dns_zone_cryptokeys / 323311123031 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DNSZoneCryptokeys DataSource Example

terraform {
  required_version = ">= 1.14"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

data "xcsh_dns_zone_cryptokeys" "example" {
}

output "dns_zone_cryptokeys_result" {
  value = data.xcsh_dns_zone_cryptokeys.example
}
```

<a id="canonical-3102000113032111-1032301323311132-3320010213230100-1323001322220300-1210332111212131-3202123023111310-3331032203303210-2012033230212302"></a>

## Root configuration — xcsh_dns_zone_cryptokeys / 323311123031 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-2001221300223032-3300202230101210-3321311012223020-3202002220323101-1233232322112222-2302312212201201-1312033312012003-3311332323130303"></a>

## Next pages — xcsh_dns_zone_cryptokeys / 323311123031 / 6

- [Property reference](../guides/data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-2123012120131000-0120323221212210-1230120031121100-3312302123120331-1103103020013131-0030222020311203-3311222211103303-2313101012023223)
- [Examples](../guides/data-sources--dns_zone_cryptokeys--examples--group-001.md#canonical-0111123332132121-3121303011101000-3222031112010012-0333301201203022-0000301011013132-0301323031221301-1101101011020231-3102203212300233)
