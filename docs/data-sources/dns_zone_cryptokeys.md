---
page_title: "xcsh_dns_zone_cryptokeys landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_cryptokeys landing."
---

# xcsh_dns_zone_cryptokeys landing

<a id="canonical-f52d736244c299ba5f4852780f6f61f432e002050062705e557a852e68c45332"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-3e3af94d568dff3d77c45bafbdd6c64e32ebe0922cc86ee36e551b1d177cd232"></a>

## xcsh_dns_zone_cryptokeys — xcsh_dns_zone_cryptokeys / 3fa0f9ef56cd / 2

Breadcrumbs:

- xcsh_dns_zone_cryptokeys

Resource creation operation.

<a id="canonical-4ff8fa25c42343d9f56d64f8ba68b2bdc553307ecd42368fd572c98231cde4d5"></a>

## Prerequisites — xcsh_dns_zone_cryptokeys / 3fa0f9ef56cd / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-d1782286356c13aaccfd656af276ebf2ff5e5855396cd96e1e94b076bf5d597a"></a>

## Minimal configuration — xcsh_dns_zone_cryptokeys / 3fa0f9ef56cd / 4

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

<a id="canonical-d20173954ec7bd5ef8127b107b07aa3064f9599de26cb574fd3a3ce4863ec9b2"></a>

## Root configuration — xcsh_dns_zone_cryptokeys / 3fa0f9ef56cd / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-81a70acef08ac464f9d46ac8e20a8ed16fbba5aab2da6861763f6183f5fbb733"></a>

## Next pages — xcsh_dns_zone_cryptokeys / 3fa0f9ef56cd / 6

- [Property reference](../guides/data-sources--dns_zone_cryptokeys--reference--group-001.md#canonical-9b19874018ee99a46c60d650f6c9b63d534c81dd0ca88d63f5aa54f3b74462eb)
- [Examples](../guides/data-sources--dns_zone_cryptokeys--examples--group-001.md#canonical-156fe799d9cc5440ea3561063fc618ca00c451de31ecda715144522dd28e6c2f)
