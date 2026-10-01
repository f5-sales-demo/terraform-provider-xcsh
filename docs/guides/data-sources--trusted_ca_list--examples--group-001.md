---
page_title: "xcsh_trusted_ca_list examples"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_trusted_ca_list examples."
---

# xcsh_trusted_ca_list examples

<a id="canonical-c5f1bedb8562795d2af752ebb622617996bd72502544fb62b1642964f38af908"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-11de96c74f664bcd2073668bb7fea3afa16ada015d575de602fd9f8f06eaadf0"></a>

## Examples — Examples / 74ce0150c360 / 2

Breadcrumbs:

- [xcsh_trusted_ca_list](../data-sources/trusted_ca_list.md#canonical-ed5a0ca4002a9acc208559e5d135f6efdf6a9d227170b35fe4e3ac7a7857549a)
- Examples

<a id="canonical-bf679d1c83b3f04564bdea1f02d104e9961cd816e6e06a1dc4696382967cdeb1"></a>

## Complete configurations — Examples / 74ce0150c360 / 3

- [Data source](data-sources--trusted_ca_list--examples--group-001.md#canonical-e6ea2f1298d99dba5a2f95738c0b6b965bff3d92c4a3e5f240ca983b5c00366e): valid configuration.

<a id="canonical-b06983eba03b26dcbda41dd402e5446b00c29ccb3ced23002863b7ef1d3c0b57"></a>

## Next pages — Examples / 74ce0150c360 / 4

- [Data source](data-sources--trusted_ca_list--examples--group-001.md#canonical-e6ea2f1298d99dba5a2f95738c0b6b965bff3d92c4a3e5f240ca983b5c00366e)
- [xcsh_trusted_ca_list](../data-sources/trusted_ca_list.md#canonical-ed5a0ca4002a9acc208559e5d135f6efdf6a9d227170b35fe4e3ac7a7857549a)

<a id="canonical-e6ea2f1298d99dba5a2f95738c0b6b965bff3d92c4a3e5f240ca983b5c00366e"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-43e433805fb52294388c3334d60d7e8ea9d38d68e42f34dc4d065ada19197561"></a>

## Data source — Data source / 40e8412a5421 / 2

Breadcrumbs:

- [xcsh_trusted_ca_list](../data-sources/trusted_ca_list.md#canonical-ed5a0ca4002a9acc208559e5d135f6efdf6a9d227170b35fe4e3ac7a7857549a)
- [Examples](data-sources--trusted_ca_list--examples--group-001.md#canonical-c5f1bedb8562795d2af752ebb622617996bd72502544fb62b1642964f38af908)
- Data source

Schema-derived minimal configuration validated with the checked-out provider.

Expected outcome: **valid configuration**.

Source: `examples/data-sources/xcsh_trusted_ca_list/data-source.tf`; digest `sha256:444d0b10e4a27614d6854b1c286e2b90a9468950261cdb881f0b8b693ceeb1e7`.

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

<a id="canonical-dd014e57c987175cf4dd61a0ad039b8214645cf795ad3427ead30f4d82153834"></a>

## Next pages — Data source / 40e8412a5421 / 3

- [Examples](data-sources--trusted_ca_list--examples--group-001.md#canonical-c5f1bedb8562795d2af752ebb622617996bd72502544fb62b1642964f38af908)
- [xcsh_trusted_ca_list](../data-sources/trusted_ca_list.md#canonical-ed5a0ca4002a9acc208559e5d135f6efdf6a9d227170b35fe4e3ac7a7857549a)
