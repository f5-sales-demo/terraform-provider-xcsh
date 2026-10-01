---
page_title: "xcsh_service_policy landing"
subcategory: "Security"
description: "Complete grouped canonical reference for xcsh_service_policy landing."
---

# xcsh_service_policy landing

<a id="canonical-562b2fff5c42ebb1cb9b4a4ad3343a608c21cb1e61939fa0a6d4331979d343b3"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-f09610c57c62fb82e33be575245b7cf656dc311bb9a2c37d7f15524eb4ddd7ae"></a>

## xcsh_service_policy — xcsh_service_policy / 45fc4ada863f / 2

Breadcrumbs:

- xcsh_service_policy

Manages service\_policy creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-27dcf350e8ff18b790adac4a9267ac69a70caca9cd0558ad0fdafb253ec87436"></a>

## Prerequisites — xcsh_service_policy / 45fc4ada863f / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

Required service tier: Advanced.

<a id="canonical-47913439acf3f2bea48255cb1a9421e6260629d765648ce983a0f0b0a3732d22"></a>

## Minimal configuration — xcsh_service_policy / 45fc4ada863f / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# ServicePolicy Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing ServicePolicy by name
data "xcsh_service_policy" "example" {
  name      = "example-service-policy"
  namespace = "staging"
}

output "service_policy_id" {
  value = data.xcsh_service_policy.example.id
}
```

<a id="canonical-b25bc78d1eb4816e83d10f13c0501b79ccb16710280aa4f958e7bbf0396fbba5"></a>

## Root configuration — xcsh_service_policy / 45fc4ada863f / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-8e3639a8844760c031c8f70980e2695e9f143c1d97d4acd85775f4c39161d64d"></a>

## Next pages — xcsh_service_policy / 45fc4ada863f / 6

- [Property reference](../guides/data-sources--service_policy--reference--group-001.md#canonical-f64b25532920797098edbc17516d2f6c4ee99a917eb7049e2fc3def30deda387)
- [Examples](../guides/data-sources--service_policy--examples--group-001.md#canonical-b43809c2072d0e215aefdfe6e7ca74d503f696f88f9d9f4abacdb2e701bde2c0)
