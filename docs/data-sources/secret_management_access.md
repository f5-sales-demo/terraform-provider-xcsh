---
page_title: "xcsh_secret_management_access landing"
subcategory: ""
description: "Complete grouped canonical reference for xcsh_secret_management_access landing."
---

# xcsh_secret_management_access landing

<a id="canonical-18a3b230047a96b1b8af81502f9fd1e264249b4a2f0875d35905c837e772229b"></a>

<!-- Exact provider and upstream contract identifiers. -->

<!-- textlint-disable terminology -->

<a id="canonical-c1e04e1de3364e3c88ff3578dcf2d7f366cb177618c1829dbdfe05592e339fa8"></a>

## xcsh_secret_management_access — xcsh_secret_management_access / 4c558ef29998 / 2

Breadcrumbs:

- xcsh_secret_management_access

Manages secret\_management\_access creates a new object in storage backend for metadata.namespace in
F5 Distributed Cloud.

<a id="canonical-4c858084d4043625d8e7c6667761f0a49a5c70c61942e0442ba65012331eb2be"></a>

## Prerequisites — xcsh_secret_management_access / 4c558ef29998 / 3

Install Terraform and the `f5-sales-demo/xcsh` provider. Configure provider authentication and access to the target namespace.

<a id="canonical-04840ce3923232a6b5974bc7087d8a8af8d7da7fbc2d9813b8a21977f9db3ba7"></a>

## Minimal configuration — xcsh_secret_management_access / 4c558ef29998 / 4

Validated with the exact checked-out provider using `terraform validate`. This does not assert a successful live apply.

```terraform
# SecretManagementAccess Data Source Example

terraform {
  required_version = ">= 1.0"

  required_providers {
    xcsh = {
      source  = "f5-sales-demo/xcsh"
      version = ">= 0.1.0"
    }
  }
}

# Look up an existing SecretManagementAccess by name
data "xcsh_secret_management_access" "example" {
  name      = "example-secret-management-access"
  namespace = "staging"
}

output "secret_management_access_id" {
  value = data.xcsh_secret_management_access.example.id
}
```

<a id="canonical-f67af0fab2e62d6dc217be03588037964a05d6b90b4cd962689ff2486d228fe7"></a>

## Root configuration — xcsh_secret_management_access / 4c558ef29998 / 5

Required root properties: `name`, `namespace`. Full root flags and choices appear in the property reference.

<a id="canonical-248e5705ccc1292e33a9669af1ea2d94444689ac721ea6c1c9efb2809ca5fa9e"></a>

## Next pages — xcsh_secret_management_access / 4c558ef29998 / 6

- [Property reference](../guides/data-sources--secret_management_access--reference--group-001.md#canonical-b6b9bdd7091c2fb1c8ea6311c122dbad79167de8645dfc9ff9b477931c1a267e)
- [Examples](../guides/data-sources--secret_management_access--examples--group-001.md#canonical-af95116420ee9e87d5fe5c3f46181f8ef7aa75cc239ecdbe6971e0be1537c201)
