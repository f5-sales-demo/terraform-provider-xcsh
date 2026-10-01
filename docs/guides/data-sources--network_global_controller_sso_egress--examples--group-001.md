---
page_title: "xcsh_network_global_controller_sso_egress examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_network_global_controller_sso_egress examples."
---

# xcsh_network_global_controller_sso_egress examples

<a id="canonical-1330212023033012-1303023012013222-0112221202020303-1011102330201000-3003322221102203-3310100311110131-1111331310003200-2210323333000232"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3101002302022003-3020030002333022-2220022220101330-1011331130230302-2320310101130003-3321121131320002-2020001102203021-2102102122320213"></a>

## Examples — Examples / 001132212120 / 2

Breadcrumbs:

- [xcsh_network_global_controller_sso_egress](../data-sources/network_global_controller_sso_egress.md#canonical-0312323020200211-1203202133322020-0021231103030131-3121030331030302-1102330101113213-3300113123100001-1202022320202320-2331132003123002)
- Examples

<a id="canonical-0010120023213203-2332102121320030-3222130030002231-0011202231022321-3011022003023212-0012211303130231-0022313322032302-1201020310023313"></a>

## Complete configurations — Examples / 001132212120 / 3

- [Data source](data-sources--network_global_controller_sso_egress--examples--group-001.md#canonical-1200211013301333-1222212313103000-2302100211003332-3011222110113330-2123221303110130-3332332310330221-3221212022321122-0313131012302212): valid configuration.

<a id="canonical-1003113111322112-2203233311202321-0011003020120103-3023200312020321-3031101022231102-0001120332011023-0203231220210002-0233321010312110"></a>

## Next pages — Examples / 001132212120 / 4

- [Data source](data-sources--network_global_controller_sso_egress--examples--group-001.md#canonical-1200211013301333-1222212313103000-2302100211003332-3011222110113330-2123221303110130-3332332310330221-3221212022321122-0313131012302212)
- [xcsh_network_global_controller_sso_egress](../data-sources/network_global_controller_sso_egress.md#canonical-0312323020200211-1203202133322020-0021231103030131-3121030331030302-1102330101113213-3300113123100001-1202022320202320-2331132003123002)

<a id="canonical-1200211013301333-1222212313103000-2302100211003332-3011222110113330-2123221303110130-3332332310330221-3221212022321122-0313131012302212"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-2220023320010013-3013320121300312-0123230123211101-3210103112013130-3212033320113130-0313230323200300-0222110210323320-2313102231102221"></a>

## Data source — Data source / 231123120323 / 2

Breadcrumbs:

- [xcsh_network_global_controller_sso_egress](../data-sources/network_global_controller_sso_egress.md#canonical-0312323020200211-1203202133322020-0021231103030131-3121030331030302-1102330101113213-3300113123100001-1202022320202320-2331132003123002)
- [Examples](data-sources--network_global_controller_sso_egress--examples--group-001.md#canonical-1330212023033012-1303023012013222-0112221202020303-1011102330201000-3003322221102203-3310100311110131-1111331310003200-2210323333000232)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_network_global_controller_sso_egress/data-source.tf`; digest `sha256:ec1f07213dac86d4f0226f84a49ea20a8bea75367b69b06ef9f30bb6b6ae2e00`.

```terraform
terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 11.3.0"
    }
  }
}

data "xcsh_network_global_controller_sso_egress" "https" {}

output "global_controller_sso_https" {
  value = {
    direction    = "egress"
    protocol     = "tcp"
    port         = 443
    destinations = data.xcsh_network_global_controller_sso_egress.https.cidr_blocks
  }
}
```

<a id="canonical-1003220303131011-2132233120111100-0113321323300231-2332033330003021-1310321230312103-2110312013013211-3330320002012101-0212102133032020"></a>

## Next pages — Data source / 231123120323 / 3

- [Examples](data-sources--network_global_controller_sso_egress--examples--group-001.md#canonical-1330212023033012-1303023012013222-0112221202020303-1011102330201000-3003322221102203-3310100311110131-1111331310003200-2210323333000232)
- [xcsh_network_global_controller_sso_egress](../data-sources/network_global_controller_sso_egress.md#canonical-0312323020200211-1203202133322020-0021231103030131-3121030331030302-1102330101113213-3300113123100001-1202022320202320-2331132003123002)
