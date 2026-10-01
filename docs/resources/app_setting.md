---
page_title: "xcsh_app_setting landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_setting landing."
---

# xcsh_app_setting landing

<a id="canonical-2f6ae5677f2d60f80270b519919167cd9a5878a0384c83ae1ca1406c83ef6a94"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c2dc6288e6ee3bafe8f0bb1b814e7c65deb39f6544c630ebe5234a4f5ef38faf"></a>

## xcsh_app_setting — xcsh_app_setting / c3a80415e6fe / 2

Breadcrumbs:

- xcsh_app_setting

Manages App setting configuration in namespace metadata.namespace in F5 Distributed Cloud.

<a id="canonical-f1cea69728597204abc1135169cfd9d1543d5f31f195f89860648108d952c1a0"></a>

## Prerequisites — xcsh_app_setting / c3a80415e6fe / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-ab8a96d78855873fbc8bb30ee403a07c0ab7affd4b8c20cb34746ecc98029cb0"></a>

## Minimal configuration — xcsh_app_setting / c3a80415e6fe / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppSetting Resource Example
# Manages App setting configuration in namespace metadata.namespace in F5 Distributed Cloud.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic AppSetting configuration
resource "xcsh_app_setting" "example" {
  name      = "example-app-setting"
  namespace = "staging"
}
```

<a id="canonical-5d5b32992357473bd91e4a671aaea4d8c9c5b96eaffab836a00754f568221ade"></a>

## Root configuration — xcsh_app_setting / c3a80415e6fe / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-a7e3fac7752c4c5722df26c172b0cbfe04f42df9f7375277e297b6027fbbbccd"></a>

## Next pages — xcsh_app_setting / c3a80415e6fe / 6

- [Property reference](../guides/resources--app_setting--reference--group-001.md#canonical-9ad21ab31c3e7b3129df74f8b6b7e6a2a6e9f291d0f10385bcc9dc540abbf804)
- [Examples](../guides/resources--app_setting--examples--group-001.md#canonical-f8697b00e9618b953de504386b5babb6260cb03fc79886704ad06f4a7bf39245)
- [Import](../guides/resources--app_setting--lifecycle--group-001.md#canonical-08403933f61643e6f70d59a2845194500eb587ef0cccd1dfc58d05ae830bc5eb)
- [Timeouts](../guides/resources--app_setting--lifecycle--group-001.md#canonical-f8ac48602978e67e98c1b28c6e0aec61f382d770703436b924de580bc37d6183)
