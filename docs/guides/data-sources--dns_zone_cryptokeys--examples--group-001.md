---
page_title: "xcsh_dns_zone_cryptokeys examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_cryptokeys examples."
---

# xcsh_dns_zone_cryptokeys examples

<a id="canonical-156fe799d9cc5440ea3561063fc618ca00c451de31ecda715144522dd28e6c2f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-93c49cde449d5cbe56295783297931b1b8852c5220c9f8e6d2c0070ccadaf7bc"></a>

## Examples — Examples / 632927f3241d / 2

Breadcrumbs:

- [xcsh_dns_zone_cryptokeys](../data-sources/dns_zone_cryptokeys.md#canonical-f52d736244c299ba5f4852780f6f61f432e002050062705e557a852e68c45332)
- Examples

<a id="canonical-a28ce5a3db8c6badff7edf5b74f49f311f876e49a815222f76f7e7e8385f1197"></a>

## Complete configurations — Examples / 632927f3241d / 3

- [Data source](data-sources--dns_zone_cryptokeys--examples--group-001.md#canonical-29ee00c021662e5f4c9f06db5937935e0a5bf01a99cd5a97d1d90b775b81101a): valid configuration.

<a id="canonical-e421e99c380b02177279e3c816fe2ef997b97d37bb2d2fe96ddb8029aac9b21e"></a>

## Next pages — Examples / 632927f3241d / 4

- [Data source](data-sources--dns_zone_cryptokeys--examples--group-001.md#canonical-29ee00c021662e5f4c9f06db5937935e0a5bf01a99cd5a97d1d90b775b81101a)
- [xcsh_dns_zone_cryptokeys](../data-sources/dns_zone_cryptokeys.md#canonical-f52d736244c299ba5f4852780f6f61f432e002050062705e557a852e68c45332)

<a id="canonical-29ee00c021662e5f4c9f06db5937935e0a5bf01a99cd5a97d1d90b775b81101a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-d65afc724479bec35e10fd95e7b93fd61b5c7f8609c4ff3636968c7c04224bec"></a>

## Data source — Data source / 4b7906299f4f / 2

Breadcrumbs:

- [xcsh_dns_zone_cryptokeys](../data-sources/dns_zone_cryptokeys.md#canonical-f52d736244c299ba5f4852780f6f61f432e002050062705e557a852e68c45332)
- [Examples](data-sources--dns_zone_cryptokeys--examples--group-001.md#canonical-156fe799d9cc5440ea3561063fc618ca00c451de31ecda715144522dd28e6c2f)
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

<a id="canonical-f91e0f31bfa9b48b538b97c8c284ee05da08ddb216049bf68292d5581b4594e6"></a>

## Next pages — Data source / 4b7906299f4f / 3

- [Examples](data-sources--dns_zone_cryptokeys--examples--group-001.md#canonical-156fe799d9cc5440ea3561063fc618ca00c451de31ecda715144522dd28e6c2f)
- [xcsh_dns_zone_cryptokeys](../data-sources/dns_zone_cryptokeys.md#canonical-f52d736244c299ba5f4852780f6f61f432e002050062705e557a852e68c45332)
