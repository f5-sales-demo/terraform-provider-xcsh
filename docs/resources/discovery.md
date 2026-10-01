---
page_title: "xcsh_discovery landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_discovery landing."
---

# xcsh_discovery landing

<a id="canonical-e1770e6e73850aecd566b504b1da3dd35f05c4301ee474c7da727b7cc3befce2"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ace907ce4e734d5fe2acb658f1a86b5295a89eeaf0c565525aa722074682bd9f"></a>

## xcsh_discovery — xcsh_discovery / 9dbf63f5c787 / 2

Breadcrumbs:

- xcsh_discovery

Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site
or virtual site in system namespace. configuration.

<a id="canonical-6fcce0091cef158bb078bdda3105891a0b0a82c05be8af33e1faabf92ad6ad23"></a>

## Prerequisites — xcsh_discovery / 9dbf63f5c787 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-e2080180eaae2b76e909148ea75fb755fb7e82aa900d04dd76888f980277b42c"></a>

## Minimal configuration — xcsh_discovery / 9dbf63f5c787 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Discovery Resource Example
# Manages a Discovery resource in F5 Distributed Cloud for api to create discovery object for a site or virtual site in system namespace.

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Basic Discovery configuration
resource "xcsh_discovery" "example" {
  name      = "example-discovery"
  namespace = "staging"
}
```

<a id="canonical-6a21899dbc0835614f5df4a9d866fcee66ca27ebe0a8876e7935b71e3777ec99"></a>

## Root configuration — xcsh_discovery / 9dbf63f5c787 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-f4294561101aa0ad480aa7c378af3b0ed5cbf85d2fd05161e5f7d0edd7044bbc"></a>

## Next pages — xcsh_discovery / 9dbf63f5c787 / 6

- [Property reference](../guides/resources--discovery--reference--group-001.md#canonical-4068b7597520638cadfe1040de3f671edb52e79a1536c2d36dcdc95475940a6a)
- [Examples](../guides/resources--discovery--examples--group-001.md#canonical-32c6db1798df2a0749a45f5581c13347d4a7697806d0a15a0d82c83f1477df2f)
- [Import](../guides/resources--discovery--lifecycle--group-001.md#canonical-77d48909c83e8df07a072cbc9d4375a914b7ccc899642c9183fa9130cc40ffc5)
- [Timeouts](../guides/resources--discovery--lifecycle--group-001.md#canonical-f89f7931bae011cf39ed92d6bc8d539f81cd3be3291d6f38cc64d0479a35c01c)
