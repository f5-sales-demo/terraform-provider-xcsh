---
page_title: "xcsh_dns_zone_cryptokeys examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_cryptokeys examples."
---

# xcsh_dns_zone_cryptokeys examples

<a id="canonical-0111123332132121-3121303011101000-3222031112010012-0333301201203022-0000301011013132-0301323031221301-1101101011020231-3102203212300233"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Examples

Breadcrumbs:

- [xcsh_dns_zone_cryptokeys](../data-sources/dns_zone_cryptokeys.md#canonical-3311023113031202-1010300221212322-1133102011021320-0033123312013310-0302320000020011-0000120213001132-1111132220110232-1220301011030302)
- Examples

<a id="canonical-2103301021303132-1010213111302332-1112022111132003-0221132103012301-2320201102301102-0200302133203212-3102300000130030-3022312233132330"></a>

### Complete configurations for `xcsh_dns_zone_cryptokeys`

- [Data source](data-sources--dns_zone_cryptokeys--examples--group-001.md#canonical-0221323200003000-0201121202321133-1030213300123123-1121031321031132-0022112333000122-2121303111222113-3101312100231313-1123200101000122): valid configuration.

<a id="canonical-0221323200003000-0201121202321133-1030213300123123-1121031321031132-0022112333000122-2121303111222113-3101312100231313-1123200101000122"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

## Data source example

Breadcrumbs:

- [xcsh_dns_zone_cryptokeys](../data-sources/dns_zone_cryptokeys.md#canonical-3311023113031202-1010300221212322-1133102011021320-0033123312013310-0302320000020011-0000120213001132-1111132220110232-1220301011030302)
- [Examples](data-sources--dns_zone_cryptokeys--examples--group-001.md#canonical-0111123332132121-3121303011101000-3222031112010012-0333301201203022-0000301011013132-0301323031221301-1101101011020231-3102203212300233)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_dns_zone_cryptokeys/data-source.tf`; digest `sha256:b7a0c08b1e74769c7e06fd735c8df6c0ed430cc8d5516a712859e61adfe28bc8`.

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
