---
page_title: "xcsh_app_setting landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_app_setting landing."
---

# xcsh_app_setting landing

<a id="canonical-d04647720e2ea3a725c9cb2d99855697f9ab8155486df509007c569f3bb3f6c9"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-08caaa0e3703bc1783208b6a5919ff32db0933f78d9ca1be355d19241f33e213"></a>

## xcsh_app_setting — xcsh_app_setting / f27e3b183298 / 2

Breadcrumbs:

- xcsh_app_setting

Manages App setting configuration in namespace metadata.namespace in F5 Distributed Cloud.

<a id="canonical-bbf5350ca6a33052e9d87d10f9d3a13134d8fc3ea18b4f85d2ecf62b2b58b208"></a>

## Prerequisites — xcsh_app_setting / f27e3b183298 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-7db890cd7cd11c69017ea295466b74516171ed86f77bea3b439a3394565e3174"></a>

## Minimal configuration — xcsh_app_setting / f27e3b183298 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AppSetting Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AppSetting by name
data "xcsh_app_setting" "example" {
  name      = "example-app-setting"
  namespace = "staging"
}

output "app_setting_id" {
  value = data.xcsh_app_setting.example.id
}
```

<a id="canonical-8f7d87707d12bcc72a445c861be27aba1b193b0008f07e6a7581467d1fb9f670"></a>

## Root configuration — xcsh_app_setting / f27e3b183298 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-dc6cc25a3deb72bcb6bafe0886663f3d822b453032db11ce87f2a0f820d344c2"></a>

## Next pages — xcsh_app_setting / f27e3b183298 / 6

- [Property reference](../guides/data-sources--app_setting--reference--group-001.md#canonical-cce4ca195fc8f62cf323b34bdc0892a876ca67a842353a6bb3d922c01cb45a73)
- [Examples](../guides/data-sources--app_setting--examples--group-001.md#canonical-a296127a92a7ce97e2aee94840e4cb600879e5446e23f1b20bca93486f3f3dbd)
