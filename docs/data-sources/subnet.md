---
page_title: "xcsh_subnet landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_subnet landing."
---

# xcsh_subnet landing

<a id="canonical-8cbfe9ffe7960aa3ec7fb3ac468c961c045fe194bd2557431e65ef71f6ef1624"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-e50d647089daadd83db09b365dcedf523d72270b23600f85d8683e19528a9394"></a>

## xcsh_subnet — xcsh_subnet / 385a57a68741 / 2

Breadcrumbs:

- xcsh_subnet

Manages a Subnet resource in F5 Distributed Cloud for subnet object contains configuration for an
interface of a vm/pod. it is created in user or shared namespace. configuration.

<a id="canonical-76b12b9ea1d0e6ad5329d7798342d8cf61120e29c170bdb8144d3b4cd900b9eb"></a>

## Prerequisites — xcsh_subnet / 385a57a68741 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-678020b58dc1a3358f26e6cb8ea52df5c4670b726433c6c05d3b8afa8f6dfcdf"></a>

## Minimal configuration — xcsh_subnet / 385a57a68741 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# Subnet Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing Subnet by name
data "xcsh_subnet" "example" {
  name      = "example-subnet"
  namespace = "staging"
}

output "subnet_id" {
  value = data.xcsh_subnet.example.id
}
```

<a id="canonical-893d94f9ad83682c10898ea366216f73a6500d07b98da321d6e7d2d75b53d2fe"></a>

## Root configuration — xcsh_subnet / 385a57a68741 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-22ac28d9f0c94b26a49090e077ee37154794af4f730ce2daf6ee1e23fdc75676"></a>

## Next pages — xcsh_subnet / 385a57a68741 / 6

- [Property reference](../guides/data-sources--subnet--reference--group-001.md#canonical-526725bd256a2af8a3f5474393f882f10468685417a7ce3018001dc9419ec1e0)
- [Examples](../guides/data-sources--subnet--examples--group-001.md#canonical-79d8e319da7c705c323908e69b3050b9dbfa8c2f97c84a0b8f524d43826ab865)
