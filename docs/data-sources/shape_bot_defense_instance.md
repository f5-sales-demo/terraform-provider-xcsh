---
page_title: "xcsh_shape_bot_defense_instance landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_shape_bot_defense_instance landing."
---

# xcsh_shape_bot_defense_instance landing

<a id="canonical-c06d624c324df14fa465e15a9ea25b4fefdb447d6a384de57f971c0b2de209bb"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-18706186ae44521e8ba3647648ae5d08561e2ee14c3ba0e833e6235e14090a4f"></a>

## xcsh_shape_bot_defense_instance — xcsh_shape_bot_defense_instance / f466d109015b / 2

Breadcrumbs:

- xcsh_shape_bot_defense_instance

Manages a Shape Bot Defense Instance resource in F5 Distributed Cloud for get virtual host from a
given namespace. configuration. (read-only data source)

<a id="canonical-aa9976f514272e3e6a43d442cfbe07e42a3d7d865f8aa6eafbef322f5a82bb6e"></a>

## Prerequisites — xcsh_shape_bot_defense_instance / f466d109015b / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-df2db5cc7f12284ada555b229226473bedebad0daa2a59284254dd2002d27483"></a>

## Minimal configuration — xcsh_shape_bot_defense_instance / f466d109015b / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ShapeBotDefenseInstance Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ShapeBotDefenseInstance by name
data "xcsh_shape_bot_defense_instance" "example" {
  name      = "example-shape-bot-defense-instance"
  namespace = "staging"
}

output "shape_bot_defense_instance_id" {
  value = data.xcsh_shape_bot_defense_instance.example.id
}
```

<a id="canonical-7ab2585e8029759ea0918a9bdaa2184dea85bc49d69bab10cf9f4ab0d738e075"></a>

## Root configuration — xcsh_shape_bot_defense_instance / f466d109015b / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-6a50912e8e96fb796c4d7fc0d778bd5088212fb528f354dfc9a4cbfa63ce6b96"></a>

## Next pages — xcsh_shape_bot_defense_instance / f466d109015b / 6

- [Property reference](../guides/data-sources--shape_bot_defense_instance--reference--group-001.md#canonical-a2f33e8f8c2c828b37e044330ecb2d9104174d36beda291e8383c017e47bec6b)
- [Examples](../guides/data-sources--shape_bot_defense_instance--examples--group-001.md#canonical-50b5f6e1abc26f913446cc3b2915bcec61e66ce812ebcafd5537d8950015d628)
