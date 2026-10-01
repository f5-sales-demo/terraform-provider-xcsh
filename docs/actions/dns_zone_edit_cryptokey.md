---
page_title: "xcsh_dns_zone_edit_cryptokey landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_dns_zone_edit_cryptokey landing."
---

# xcsh_dns_zone_edit_cryptokey landing

<a id="canonical-5ea34fd0806e345b9da4127afca9297afdff2223ccf404858b3a86b9b6048b4c"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-060472e300d90d263822b16164d18a6c262bdea52ac2d36ef24a8a53bfc59c8d"></a>

## xcsh_dns_zone_edit_cryptokey — xcsh_dns_zone_edit_cryptokey / 736db4e42228 / 2

Breadcrumbs:

- xcsh_dns_zone_edit_cryptokey

Resource creation operation.

<a id="canonical-77e10daf80304965f3e3387803632c1356a912be85bcf526bc8cbb41513eb4c1"></a>

## Prerequisites — xcsh_dns_zone_edit_cryptokey / 736db4e42228 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-99569bd26c14fdb77ebc6e10acb44386cee6056e0f6f6d6aad0b0ea21f808ccb"></a>

## Minimal configuration — xcsh_dns_zone_edit_cryptokey / 736db4e42228 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

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

<a id="canonical-141688703a746427828f797a655ca274e92469ca42b85dde56a43362724c5299"></a>

## Root configuration — xcsh_dns_zone_edit_cryptokey / 736db4e42228 / 5

Required root properties: none. Full root flags and choices appear in the property reference.

<a id="canonical-9152cfa6dcfd600b5242d6c50b8579579b9dc9e636bc038cb4c7f4eaf2e30da3"></a>

## Next pages — xcsh_dns_zone_edit_cryptokey / 736db4e42228 / 6

- [Property reference](../guides/actions--dns_zone_edit_cryptokey--reference--group-001.md#canonical-8c778cc8bc7098812305b349d8624d6e5e0355d7a00fe2afec983f56a25a2bb4)
- [Examples](../guides/actions--dns_zone_edit_cryptokey--examples--group-001.md#canonical-479eb8f81e1535f664c0e7ed207220585c6f8afd3f1c9c142614f65eecefb4ee)
- [Lifecycle](../guides/actions--dns_zone_edit_cryptokey--lifecycle--group-001.md#canonical-1fd6abb21cbb8c48adb6bff95cf15d52c3ec06904c1fdfb329840deeaa99dda3)
