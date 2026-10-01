---
page_title: "xcsh_advertise_policy landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_advertise_policy landing."
---

# xcsh_advertise_policy landing

<a id="canonical-77ee33ff70b16dfe904993ef2a4855e549c025f272b16e2a9e8fc0423965ca2a"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-ed8472d3a190fc29ce09d4981355e77ba278aec007f3b90bb0394f64122dad58"></a>

## xcsh_advertise_policy — xcsh_advertise_policy / 17173258e110 / 2

Breadcrumbs:

- xcsh_advertise_policy

Manages a Advertise Policy resource in F5 Distributed Cloud for advertise\_policy object controls
how and where a service represented by a given virtual\_host object is advertised to consumers.
configuration.

<a id="canonical-6827282011d6d5931e39640a9e29ea952efab90af612bffe48156eb7a38107ec"></a>

## Prerequisites — xcsh_advertise_policy / 17173258e110 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-8370a8b0d609a980452cb4522262e80e05320bdaba764d5e95240989e3a4724f"></a>

## Minimal configuration — xcsh_advertise_policy / 17173258e110 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# AdvertisePolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing AdvertisePolicy by name
data "xcsh_advertise_policy" "example" {
  name      = "example-advertise-policy"
  namespace = "staging"
}

output "advertise_policy_id" {
  value = data.xcsh_advertise_policy.example.id
}
```

<a id="canonical-c21a1a6fc686397ecf3782394d4bbd307d8668f82067c78b1550a8c20c3296df"></a>

## Root configuration — xcsh_advertise_policy / 17173258e110 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-e773137e0af83dc22053ce27d7da79c86645ab2ad3d6c0004a36134caccecb84"></a>

## Next pages — xcsh_advertise_policy / 17173258e110 / 6

- [Property reference](../guides/data-sources--advertise_policy--reference--group-001.md#canonical-c4a8545fdd17da3cdb3ef89ede55dd3f490b1b6a1e1ad35e3f19bc23129b3595)
- [Examples](../guides/data-sources--advertise_policy--examples--group-001.md#canonical-c122cffd8cafcc703c56445d899d66062863b9857605e22c90abff36dee4d815)
