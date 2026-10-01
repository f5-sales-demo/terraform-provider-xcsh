---
page_title: "xcsh_data_type landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_data_type landing."
---

# xcsh_data_type landing

<a id="canonical-82775399919a9b1c71c7b859aa40315ad8471412265f9074c4f55f8a02f8001f"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-32568c73a612f34822551ce34560450e0449bc1b7ded062a0d61c0b206f9fde0"></a>

## xcsh_data_type — xcsh_data_type / b2a9f6f97a41 / 2

Breadcrumbs:

- xcsh_data_type

Manages data\_type creates a new object in the storage backend for metadata.namespace in F5
Distributed Cloud.

<a id="canonical-5a829cb2884a5e94368d52ac0998a48cc80714844c756203419c292653c950ee"></a>

## Prerequisites — xcsh_data_type / b2a9f6f97a41 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-54f972c8b1b42694215b9f6e648ea60ff4e69a01c06247ca1f690032302c02ea"></a>

## Minimal configuration — xcsh_data_type / b2a9f6f97a41 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# DataType Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing DataType by name
data "xcsh_data_type" "example" {
  name      = "example-data-type"
  namespace = "staging"
}

output "data_type_id" {
  value = data.xcsh_data_type.example.id
}
```

<a id="canonical-6e302cab0df169b6a15b4b86554049cc0d9b86a6287b6daae9ee19e9698a2437"></a>

## Root configuration — xcsh_data_type / b2a9f6f97a41 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-87bfb320d4770651c13b0e1d05bebc3b285c29e8a0c26934e40caf24b597d63b"></a>

## Next pages — xcsh_data_type / b2a9f6f97a41 / 6

- [Property reference](../guides/data-sources--data_type--reference--group-001.md#canonical-11669e6830386716c752f054614619da6e1647de8f830f8c42b779b52e4ed646)
- [Examples](../guides/data-sources--data_type--examples--group-001.md#canonical-165dbfcb731f3d0db618fbe9d30ce1e3d10aad80ef7d2255d69f075394000722)
